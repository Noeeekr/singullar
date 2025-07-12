package migrations

import (
	"database/sql"
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type Migrations struct {
	// Having a single ctx for all cases is kinda bad in concurrency since the code underneath would make it go crazy:
	//
	// 	m := NewMigrationsManager(db)
	// 	go m.CreateTable(tables)
	// 	go m.CreateTable(tables)
	//

	// must be true on instance creation
	callStack int
	ctx       *Context
	*transactions.Manager
	ops *operations.Operations
}

func New(db *sql.DB) *Migrations {
	return &Migrations{
		ctx: &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		},
		Manager:   transactions.NewManager(db),
		ops:       operations.New(db),
		callStack: 0,
	}
}

func (m *Migrations) CreateDatabaseUser(user *models.CreateDatabaseUser, configuration *Configuration) (res *common.Response) {
	var dbusers []string
	res = m.Query(
		// Could be in a model, inserted in a operation
		transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_roles WHERE rolname = '%s'", user.Name)).
			WithRowsScanner(scan.DatabaseUsers(&dbusers)).ThrowErrorOnFound(),
	)
	if res != nil {
		// Return error
		if res.Status != common.StatusFound {
			return res
		}
		// Handle StatusFound cases
		if configuration.IgnoreExisting {
			logs.Info.Println("[Ignore existing flag]: Ignoring existing user: " + user.Name)
			return nil
		}
		if configuration.RecreateExisting {
			logs.Info.Println("[Recreate existing flag]: Dropping existing user: " + user.Name)
			res = m.DropDatabaseUsers(user.Name)
			if res != nil {
				return res
			}
		}
	}

	logs.Info.Println("[Creating user]: " + user.Name)
	res = m.Query(transactions.NewRequest(fmt.Sprintf(`
			CREATE USER %s WITH 
			PASSWORD '%s' 
			LOGIN;
		`, user.Name, user.Password)))
	if res != nil {
		return res
	}
	return nil
}

// CreateDatabaseUsers creates users until first error..
func (m *Migrations) CreateDatabaseUsers(users []*models.CreateDatabaseUser, configuration *Configuration) (res *common.Response) {
	for _, user := range users {
		res = m.CreateDatabaseUser(user, configuration)
		if res != nil {
			return res
		}
	}

	return res
}

// Doesn't need users.Password to be set
func (m *Migrations) GrantAllPrivilegesOnDatabase(users []*models.CreateDatabaseUser) (res *common.Response) {
	for _, user := range users {
		var dbusers []string
		res = m.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT rolname FROM pg_roles WHERE rolname = '%s'", user.Name)).
				WithRowsScanner(scan.DatabaseUsers(&dbusers)),
		)
		if res != nil {
			return res
		}

		var dbnames []string
		res = m.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT rolname FROM pg_roles WHERE rolname = '%s'", user.Name)).
				WithRowsScanner(scan.DatabaseNames(&dbnames)),
		)
		if res != nil {
			return res
		}

		logs.Info.Println("[Grantting all privileges on database]: " + user.Name + " => " + user.Database)
		res := m.Query(transactions.NewRequest(
			fmt.Sprintf(
				"GRANT ALL PRIVILEGES ON DATABASE %s TO %s;",
				user.Database, user.Name,
			),
		))
		if res != nil {
			return res
		}
	}

	return nil
}

func (m *Migrations) DropDatabaseUsers(names ...string) (res *common.Response) {
	for _, name := range names {
		// check if users already exist
		var dbusers []string
		res = m.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT rolname FROM pg_roles WHERE rolname = '%s'", name)).
				WithRowsScanner(scan.DatabaseUsers(&dbusers)),
		)
		if res != nil {
			return res
		}

		logs.Info.Println("[Dropping user owned relations]: " + name)
		res = m.Query(transactions.NewRequest(
			fmt.Sprintf("DROP OWNED BY %s CASCADE;", name)),
		)
		if res != nil {
			return res
		}

		var datnames []string
		res = m.Query(
			transactions.NewRequest(
				fmt.Sprintf("SELECT d.datname FROM pg_catalog.pg_database as d INNER JOIN pg_catalog.pg_user as u ON u.usesysid = d.datdba WHERE u.usename = '%s';", name),
			).WithRowsScanner(scan.DatabaseNames(&datnames)),
		)
		if res != nil && res.Status != common.StatusNotFound {
			return res
		}

		dnames := ""
		for _, datname := range datnames {
			dnames += " " + datname
		}
		logs.Info.Println("[Dropping user databases]: " + name + " =>" + dnames)

		res = m.DropDatabases(datnames...)
		if res != nil {
			return res
		}

		logs.Info.Println("[Dropping user]: " + name)
		res = m.Query(transactions.NewRequest(
			fmt.Sprintf("DROP USER %s;", name)),
		)
		if res != nil {
			return res
		}
	}
	return res
}

func (m *Migrations) CreateDatabases(requests []*RequestCreateDatabase, configuration *Configuration) *common.Response {
	for _, request := range requests {
		res := m.CreateDatabase(request, configuration)
		if res != nil {
			return res
		}
	}

	return nil
}

func (m *Migrations) CreateDatabase(request *RequestCreateDatabase, configuration *Configuration) *common.Response {
	if configuration == nil {
		configuration = &Configuration{}
	}

	var dbnames []string
	res := m.Query(
		transactions.NewRequest(fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s'", request.Database)).
			WithRowsScanner(scan.DatabaseNames(&dbnames)).ThrowErrorOnFound(),
	)
	if res != nil {
		if res.Status != common.StatusFound {
			return res
		}
		if configuration.IgnoreExisting {
			logs.Info.Println("[Ignore existing flag]: Ignoring existing database: " + request.Database)
			return nil
		}
		if configuration.RecreateExisting {
			logs.Info.Println("[Recreate existing flag]: Dropping existing database: " + request.Database)
			res := m.DropDatabases(request.Database)
			if res != nil {
				return res
			}
		}
	}

	logs.Info.Println("[Creating database]: " + request.Database)
	res = m.Query(transactions.NewRequest(fmt.Sprintf("CREATE DATABASE %s WITH OWNER = %s", request.Database, request.User)))
	if res != nil {
		return res
	}
	return nil
}

// Drops all databases found and returns an error if at least one wasn't found
func (m *Migrations) DropDatabases(names ...string) (res *common.Response) {
	for _, name := range names {
		var dbnames []string
		res = m.Query(transactions.NewRequest(fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s'", name)).
			// Returns nil if the database exists
			WithRowsScanner(scan.DatabaseNames(&dbnames)),
		)
		if res != nil {
			continue
		}

		logs.Info.Println("[Dropping database]: " + name)
		res = m.Query(transactions.NewRequest(fmt.Sprintf("DROP DATABASE %s;", name)))
		if res != nil {
			return res
		}
	}

	return res
}

func (m *Migrations) CreateTable(configuration *Configuration, tableMethods models.TableMethods) *common.Response {
	if configuration == nil {
		configuration = &Configuration{}
	}

	transaction := m.Start()
	if transaction.Response != nil {
		return transaction.Response
	}

	name := tableMethods.Name()
	// Skip table if already exists
	_, exists := m.ctx.alreadyCreatedTables[name]
	if exists {
		return nil
	}

	// Check if table exists
	var dbnames []string
	transaction = transaction.Query(
		transactions.NewRequest(fmt.Sprintf("SELECT tablename FROM pg_catalog.pg_tables WHERE tablename = '%s';", tableMethods.Name())).
			WithRowsScanner(scan.DatabaseNames(&dbnames)).ThrowErrorOnFound(),
	)
	if transaction.Response != nil {
		if transaction.Response.Status != common.StatusFound {
			return transaction.Response
		}
		if configuration.IgnoreExisting {
			transaction.Response = nil
			return transaction.Response
		}
		if configuration.RecreateExisting {
			transaction.Response = nil
			logs.Info.Println("[Recreate existing flag]")
			transaction = m.DropTables(tableMethods)
			if transaction.Response != nil {
				return transaction.Response
			}
		}
	}
	table_dependencies := tableMethods.CreateRequestDependencies()
	// Create table parent types
	for _, typ := range table_dependencies.Types {
		// Skip type if already exists
		_, exists := m.ctx.alreadyCreatedTypes[typ.Name]
		if exists {
			continue
		}
		transaction := m.CreateTypes(configuration, typ)
		if transaction.Response != nil {
			return transaction.Response
		}

		m.ctx.alreadyCreatedTypes[typ.Name] = true
	}

	// Create table parent tables
	for _, subtable := range table_dependencies.Tables {
		_, exists := m.ctx.alreadyCreatedTables[subtable.Name()]
		if exists {
			continue
		}

		m.callStack++
		transaction.Response = m.CreateTable(configuration, subtable)
		m.callStack--

		if transaction.Response != nil {
			return transaction.Response
		}

		m.ctx.alreadyCreatedTables[subtable.Name()] = true
	}

	// Create table after creating its parent types and tables
	logs.Info.Println("[Creating table]: " + tableMethods.Name())
	transaction = transaction.Query(tableMethods.GetCreateRequest())
	if transaction.Response != nil {
		return transaction.Response
	}

	m.ctx.alreadyCreatedTables[tableMethods.Name()] = true

	// prevents recursive calls to reset context before the time
	if m.callStack == 0 {
		m.ctx = &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		}
	}

	return transaction.Response
}

// If transaction is different than nil, executes in the context of the given transaction without commiting. Otherwise creates a new transaction and commits at the end.
func (m *Migrations) CreateTables(configuration *Configuration, tables ...models.TableMethods) *common.Response {
	if configuration == nil {
		configuration = &Configuration{}
	}
	transaction := m.Start()
	if transaction.Response != nil {
		return transaction.Response
	}

	// Create each table until error
	for _, table := range tables {
		m.callStack++
		transaction.Response = m.CreateTable(configuration, table)
		m.callStack--
		if transaction.Response != nil {
			return transaction.Response
		}
	}
	// prevents recursive calls to reset context before the time
	if m.callStack == 0 {
		m.ctx = &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		}
	}

	return transaction.Response
}

func (m *Migrations) CreateTypes(configuration *Configuration, types ...*models.TypeInfo) *transactions.Transaction {
	if configuration == nil {
		configuration = &Configuration{}
	}
	transaction := m.Start()
	if transaction.Response != nil {
		return transaction
	}
	for _, typ := range types {
		var typnames []string
		transaction = transaction.Query(
			transactions.NewRequest(fmt.Sprintf(`SELECT typname FROM pg_catalog.pg_type WHERE typname = '%s'`, typ.Name)).
				WithRowsScanner(scan.DatabaseTypes(&typnames)).ThrowErrorOnFound(),
		)

		if transaction.Response != nil {
			if transaction.Response.Status != common.StatusFound {
				return transaction
			}
			if configuration.IgnoreExisting {
				logs.Info.Println("[Ignore existing flag]: Ignoring existing type: " + typ.Name)
				return nil
			}
			if configuration.RecreateExisting {
				logs.Info.Println("[Recreate existing flag]: Dropping existing type: " + typ.Name)
				transaction.Response = nil
				transaction = m.DropTypes(typ)
				if transaction.Response != nil {
					return transaction
				}
			}
		}

		logs.Info.Println("[Creating type]: " + typ.Name)
		transaction.Response = transaction.Query(typ.Queries.Create).Response
		if transaction.Response != nil {
			return transaction
		}
	}

	return transaction
}

func (m *Migrations) DropTables(tablesMethods ...models.TableMethods) *transactions.Transaction {
	tx := m.Start()
	if tx.Response != nil {
		return tx
	}

	for _, tables := range tablesMethods {
		logs.Info.Println("[Dropping table]: " + tables.Name())
		tx := tx.Query(tables.GetDropRequest())
		if tx.Response != nil {
			return tx
		}
	}

	return tx
}

func (m *Migrations) DropTypes(types ...*models.TypeInfo) *transactions.Transaction {
	tx := m.Start()

	for _, typ := range types {
		logs.Info.Println("[Dropping type]: " + typ.Name)
		tx = tx.Query(typ.Queries.Drop)
		if tx.Response != nil {
			return tx
		}
	}

	return tx
}
