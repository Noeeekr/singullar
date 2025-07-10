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
	resetContext bool
	ctx          *Context
	tx           *transactions.TransactionManager
	ops          *operations.Operations
}

func New(db *sql.DB) *Migrations {
	return &Migrations{
		ctx: &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		},
		tx:           transactions.New(db),
		ops:          operations.New(db),
		resetContext: true,
	}
}

func (m *Migrations) StartTransaction() *transactions.Transaction {
	return m.tx.Start()
}

func (m *Migrations) CreateDatabaseUser(user *models.CreateDatabaseUser, configuration *Configuration) (res *common.Response) {
	var dbusers []string
	res = m.tx.Query(
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
	res = m.tx.Query(transactions.NewRequest(fmt.Sprintf(`
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
		res = m.tx.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT rolname FROM pg_roles WHERE rolname = '%s'", user.Name)).
				WithRowsScanner(scan.DatabaseUsers(&dbusers)),
		)
		if res != nil {
			return res
		}

		var dbnames []string
		res = m.tx.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT rolname FROM pg_roles WHERE rolname = '%s'", user.Name)).
				WithRowsScanner(scan.DatabaseNames(&dbnames)),
		)
		if res != nil {
			return res
		}

		logs.Info.Println("[Grantting all privileges on database]: " + user.Name + " => " + user.Database)
		res := m.tx.Query(transactions.NewRequest(
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
		res = m.tx.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT rolname FROM pg_roles WHERE rolname = '%s'", name)).
				WithRowsScanner(scan.DatabaseUsers(&dbusers)),
		)
		if res != nil {
			return res
		}

		logs.Info.Println("[Dropping user]: " + name)
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("DROP USER %s;", name)))
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
	res := m.tx.Query(
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
	res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("CREATE DATABASE %s WITH OWNER = %s", request.Database, request.User)))
	if res != nil {
		return res
	}
	return nil
}

// Drops all databases found and returns an error if at least one wasn't found
func (m *Migrations) DropDatabases(names ...string) (res *common.Response) {
	for _, name := range names {
		var dbnames []string
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s'", name)).
			// Returns nil if the database exists
			WithRowsScanner(scan.DatabaseNames(&dbnames)),
		)
		if res != nil {
			continue
		}

		fmt.Println("[Dropping database]: " + name)
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("DROP DATABASE %s;", name)))
		if res != nil {
			return res
		}
	}

	return res
}

func (m *Migrations) CreateTable(configuration *Configuration, transaction *transactions.Transaction, tableMethods models.TableMethods) *transactions.Transaction {
	if configuration == nil {
		configuration = &Configuration{}
	}

	var singleOperation bool
	if transaction == nil {
		singleOperation = true
		transaction = m.tx.Start()
		if transaction.Response != nil {
			return transaction
		}
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
			return transaction
		}
		if configuration.IgnoreExisting {
			transaction.Response = nil
			return transaction
		}
		if configuration.RecreateExisting {
			transaction.Response = nil
			fmt.Println("[Recreate existing flag active]")
			transaction = m.DropTables(transaction, tableMethods)
			if transaction.Response != nil {
				return transaction
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
		transaction := m.CreateTypes(configuration, transaction, typ)
		if transaction.Response != nil {
			return transaction
		}

		m.ctx.alreadyCreatedTypes[typ.Name] = true
	}

	// Create table parent tables
	for _, subtable := range table_dependencies.Tables {
		_, exists := m.ctx.alreadyCreatedTables[subtable.Name()]
		if exists {
			continue
		}

		m.resetContext = false
		transaction := m.CreateTable(configuration, transaction, subtable)
		m.resetContext = true

		if transaction.Response != nil {
			return transaction
		}

		m.ctx.alreadyCreatedTables[subtable.Name()] = true
	}

	// Create table after creating its parent types and tables
	fmt.Println("[Creating table]: " + tableMethods.Name())
	transaction.Response = transaction.Query(tableMethods.GetCreateRequest()).Response
	if transaction.Response != nil {
		return transaction
	}

	m.ctx.alreadyCreatedTables[tableMethods.Name()] = true

	// prevents recursive calls to reset context before the time
	if m.resetContext {
		m.ctx = &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		}
	}

	if singleOperation {
		return transaction.Commit()
	}
	return transaction
}

// If transaction is different than nil, executes in the context of the given transaction without commiting. Otherwise creates a new transaction and commits at the end.
func (m *Migrations) CreateTables(configuration *Configuration, transaction *transactions.Transaction, tables ...models.TableMethods) *transactions.Transaction {
	if configuration == nil {
		configuration = &Configuration{}
	}
	var tx *transactions.Transaction
	if transaction == nil {
		tx = m.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	// Create each table until error
	for _, table := range tables {
		tx := m.CreateTable(configuration, transaction, table)
		if tx.Response != nil {
			return tx
		}
	}
	// prevents recursive calls to reset context before the time
	if m.resetContext {
		m.ctx = &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		}
	}

	if transaction == nil {
		return tx.Commit()
	}
	return tx
}

func (m *Migrations) CreateTypes(configuration *Configuration, transaction *transactions.Transaction, types ...*models.TypeInfo) *transactions.Transaction {
	if configuration == nil {
		configuration = &Configuration{}
	}

	var singleOperation bool
	if transaction == nil {
		singleOperation = true
		transaction = m.tx.Start()
		if transaction.Response != nil {
			return transaction
		}
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
				transaction = m.DropTypes(transaction, typ)
				if transaction.Response != nil {
					return transaction
				}
			}
		}

		fmt.Println("[Creating type]: " + typ.Name)
		transaction.Response = transaction.Query(typ.Queries.Create).Response
		if transaction.Response != nil {
			return transaction
		}
	}

	if singleOperation {
		return transaction.Commit()
	}

	return transaction
}

func (m *Migrations) DropTables(transaction *transactions.Transaction, tablesMethods ...models.TableMethods) *transactions.Transaction {
	var tx *transactions.Transaction
	if transaction == nil {
		tx = m.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	for _, tableMethods := range tablesMethods {
		fmt.Println("[Dropping table]: " + tableMethods.Name())
		tx.Response = tx.Query(tableMethods.GetDropRequest()).Response
		if tx.Response != nil {
			return tx
		}
	}

	if transaction == nil {
		return tx.Commit()
	}
	return tx
}

func (m *Migrations) DropTypes(transaction *transactions.Transaction, types ...*models.TypeInfo) *transactions.Transaction {
	var tx *transactions.Transaction
	if transaction == nil {
		tx = m.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	for _, typ := range types {
		fmt.Println("[Dropping type]: " + typ.Name)
		tx = tx.Query(typ.Queries.Drop)
		if tx.Response != nil {
			return tx
		}
	}

	if transaction == nil {
		return tx.Commit()
	}

	return transaction
}
