package migrations

import (
	"database/sql"

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
	ctx *Context
	tx  *transactions.TransactionManager
}

func New(db *sql.DB) *MigrationsManager {
	return &MigrationsManager{
		tx: transactions.New(db),
	}
}

// Ignore context, it is used internally to store the dependencies that were already created, put nil instead.
func (m *MigrationsManager) CreateTables(ctx *Context, tables ...models.TableMethods) *transactions.Response {
	if ctx == nil {
		ctx = &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		}
	}

	tx, res := m.tx.Start()
	if res != nil {
		return res
	}

	// Create each table
	for _, table := range tables {
		name := table.Name()

		// Skip table if already exists
		_, exists := ctx.alreadyCreatedTables[name]
		if exists {
			continue
		}

		table_dependencies := table.CreateRequestDependencies()
		// Create table parent types
		for _, typ := range table_dependencies.Types {

			// Skip type if already exists
			_, exists := ctx.alreadyCreatedTypes[typ.Name]
			if exists {
				continue
			}

			res := tx.Query(typ.Queries.Create)
			if res != nil {
				return res
			}

			ctx.alreadyCreatedTypes[typ.Name] = true
		}

		// Create table parent tables
		for _, subtable := range table_dependencies.Tables {
			_, exists := ctx.alreadyCreatedTables[subtable.Name()]
			if exists {
				continue
			}

			res := m.CreateTables(ctx, subtable)
			if res != nil {
				return res
			}

			ctx.alreadyCreatedTables[subtable.Name()] = true
		}

		// Create table after creating its parent types and tables
		res := tx.Query(table.CreateRequest())
		if res != nil {
			return res
		}

		ctx.alreadyCreatedTables[table.Name()] = true
	}

	m.ctx = &Context{
		alreadyCreatedTables: map[models.TableName]bool{},
		alreadyCreatedTypes:  map[models.TypeName]bool{},
	}

	return tx.Commit()
}

func (m *MigrationsManager) CreateType(typ *models.TypeInfo) *transactions.Response {
	tx, res := m.tx.Start()
	if res != nil {
		return res
	}

	res = tx.Query(typ.Queries.Create)
	if res != nil {
		return res
	}

	return tx.Commit()
}

func (m *MigrationsManager) DropTables(tables ...models.TableMethods) *transactions.Response {
	tx, res := m.tx.Start()
	if res != nil {
		return res
	}

	for _, table := range tables {
		res := tx.Query(table.DropRequest())
		if res != nil {
			return res
		}
	}

	return tx.Commit()
}

func (m *MigrationsManager) DropType(typ *models.TypeInfo) *transactions.Response {
	tx, res := m.tx.Start()
	if res != nil {
		return res
	}

	res = tx.Query(typ.Queries.Drop)
	if res != nil {
		return res
	}

	return tx.Commit()
}
