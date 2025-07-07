package migrations

import (
	"database/sql"
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type Context struct {
	alreadyCreatedTables map[models.TableName]bool
	alreadyCreatedTypes  map[models.TypeName]bool
}

type MigrationsManager struct {
	// Having a single ctx for all cases is kinda bad in concurrency since the code underneath would make it go crazy:
	//
	// 	m := NewMigrationsManager(db)
	// 	go m.CreateTable(tables)
	// 	go m.CreateTable(tables)
	//

	resetContext bool
	ctx          *Context
	tx           *transactions.TransactionManager
}

func New(db *sql.DB) *MigrationsManager {
	return &MigrationsManager{
		ctx: &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		},
		tx: transactions.New(db),
	}
}

func (m *MigrationsManager) StartTransaction() *transactions.Transaction {
	return m.tx.Start()
}

// Doesn't need user.Database to be set | drop users if fail happens
func (m *MigrationsManager) CreateUsers(users ...*models.CreateDatabaseUser) (res *common.Response) {
	// Create users until error
	for _, user := range users {
		// check if users already exist
		res = m.tx.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_roles WHERE rolname = '%s'", user.Name)).
				WithScanFunc(func(rows *sql.Rows) *common.Response {
					for rows.Next() {
						return common.NewResponse().
							WithDescription("User already exists").
							WithStatus(common.StatusFound)
					}

					if rows.Err() != nil {
						return common.NewResponse().
							WithDescription(rows.Err().Error()).
							WithStatus(common.StatusInternalError)
					}

					return nil
				}),
		)
		if res != nil {
			break
		}

		fmt.Println("[Creating user]: " + user.Name)
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf(`
			CREATE USER %s WITH 
				PASSWORD '%s' 
				LOGIN;
		`, user.Name, user.Password)))
		if res != nil {
			break
		}
	}

	return res
}

// Doesn't need users.Password to be set
func (m *MigrationsManager) GrantAllPrivilegesOnDatabase(users ...*models.CreateDatabaseUser) (res *common.Response) {
	for _, user := range users {
		// check if users already exist
		res = m.tx.Query(
			transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_roles WHERE rolname = '%s'", user.Name)).
				WithScanFunc(func(rows *sql.Rows) *common.Response {
					for rows.Next() {
						return nil
					}

					if rows.Err() != nil {
						return common.NewResponse().
							WithDescription(rows.Err().Error()).
							WithStatus(common.StatusInternalError)
					}

					return common.NewResponse().WithDescription("User not found " + user.Name).WithStatus(common.StatusNotFound)
				}),
		)
		if res != nil {
			return res
		}

		fmt.Println("[Grantting all privileges on database]: " + user.Name + " => " + user.Database)
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

func (m *MigrationsManager) DropUsers(names ...string) (res *common.Response) {
	for _, name := range names {
		// check if users already exist
		res = m.tx.Query(
			// returns nil if user exist
			transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_roles WHERE rolname = '%s'", name)).
				WithScanFunc(func(rows *sql.Rows) *common.Response {
					for rows.Next() {
						return nil
					}

					if rows.Err() != nil {
						return common.NewResponse().
							WithDescription(rows.Err().Error()).
							WithStatus(common.StatusInternalError)
					}

					return common.NewResponse().WithDescription("User doesn't exist").WithStatus(common.StatusNotFound)
				}),
		)
		if res != nil {
			return res
		}

		fmt.Println("[Dropping user]: " + name)
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("DROP USER %s;", name)))
		if res != nil {
			return res
		}
	}
	return res
}

func (m *MigrationsManager) CreateDatabases(names ...string) *common.Response {
	for _, name := range names {
		res := m.tx.Query(
			// returns nil if database doesn't exist
			transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s'", name)).
				WithScanFunc(func(rows *sql.Rows) *common.Response {
					defer rows.Close()
					if rows.Next() {
						return common.NewResponse().WithDescription("Database already exists.").WithStatus(common.StatusFound)
					}
					if rows.Err() != nil {
						return common.NewResponse().WithDescription(rows.Err().Error()).WithStatus(common.StatusFailedTransaction)
					}
					return nil
				}),
		)
		if res != nil {
			fmt.Println("[Database already exists]: " + name)
			return res
		}

		fmt.Println("[Creating database]: " + name)
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("CREATE DATABASE %s", name)))
		if res != nil {
			return res
		}
	}

	return nil
}

// Drops all databases found and returns an error if at least one wasn't found
func (m *MigrationsManager) DropDatabases(names ...string) (res *common.Response) {
	for _, name := range names {
		res = m.tx.Query(transactions.NewRequest(fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s'", name)).
			// Returns nil if the database exists
			WithScanFunc(func(rows *sql.Rows) *common.Response {
				defer rows.Close()
				if rows.Next() {
					return nil
				}
				if rows.Err() != nil {
					return common.NewResponse().WithDescription(rows.Err().Error()).WithStatus(common.StatusFailedTransaction)
				}
				return common.NewResponse().WithDescription("Database not found").WithStatus(common.StatusNotFound)
			}),
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

// If transaction is different than nil, executes in the context of the given transaction without commiting. Otherwise creates a new transaction and commits at the end.
func (m *MigrationsManager) CreateTables(transaction *transactions.Transaction, tables ...models.TableMethods) *transactions.Transaction {
	var tx *transactions.Transaction
	if transaction == nil {
		tx = m.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	// Create each table
	for _, table := range tables {
		name := table.Name()
		// Skip table if already exists
		_, exists := m.ctx.alreadyCreatedTables[name]
		if exists {
			continue
		}

		table_dependencies := table.CreateRequestDependencies()
		// Create table parent types
		for _, typ := range table_dependencies.Types {

			// Skip type if already exists
			_, exists := m.ctx.alreadyCreatedTypes[typ.Name]
			if exists {
				continue
			}

			tx := m.CreateType(tx, typ)
			if tx.Response != nil {
				return tx
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
			tx := m.CreateTables(tx, subtable)
			m.resetContext = true

			if tx.Response != nil {
				return tx
			}

			m.ctx.alreadyCreatedTables[subtable.Name()] = true
		}

		// Create table after creating its parent types and tables
		fmt.Println("[Creating table if not exists]: " + name)
		tx.Response = tx.Query(table.CreateRequest()).Response
		if tx.Response != nil {
			return tx
		}

		m.ctx.alreadyCreatedTables[table.Name()] = true
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

func (m *MigrationsManager) CreateType(transaction *transactions.Transaction, types ...*models.TypeInfo) *transactions.Transaction {
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
		fmt.Println("[Creating type if not exists]: " + typ.Name)
		tx.Response = tx.Query(typ.Queries.Create).Response

		if tx.Response != nil {
			return tx
		}
	}

	if transaction == nil {
		return tx.Commit()
	}

	return tx
}

func (m *MigrationsManager) DropTables(transaction *transactions.Transaction, tables ...models.TableMethods) *transactions.Transaction {
	var tx *transactions.Transaction
	if transaction == nil {
		tx = m.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	for _, table := range tables {
		fmt.Println("[Dropping table]: " + table.Name())

		tx.Response = tx.Query(table.DropRequest()).Response
		if tx.Response != nil {
			return tx
		}
	}

	return tx.Commit()
}

func (m *MigrationsManager) DropType(transaction *transactions.Transaction, types ...*models.TypeInfo) *transactions.Transaction {
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
