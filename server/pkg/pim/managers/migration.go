package managers

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/pkg/pim/models"
	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
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
	tx  *transaction.Manager
}

func NewMigrationManager(db *sql.DB) *MigrationsManager {
	return &MigrationsManager{
		tx: transaction.New(db),
	}
}

func (m *MigrationsManager) CreateType(typ *models.TypeInfo) *transaction.Response {
	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse()
		res.SetStatus(transaction.StatusFailedTransactionStart)
		res.SetDescription("Unable to start Transaction. " + err.Error())
		return res
	}

	res := tx.Query(typ.Queries.Create)

	if res.Status != transaction.StatusSuccess {
		return res
	}

	return tx.Commit()
}

// Ignore context, it is used internally to store the dependencies that were already created, put nil instead.
func (m *MigrationsManager) CreateTables(ctx *Context, tables ...*models.TableInfo) *transaction.Response {
	if ctx == nil {
		ctx = &Context{
			alreadyCreatedTables: map[models.TableName]bool{},
			alreadyCreatedTypes:  map[models.TypeName]bool{},
		}
	}

	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse()
		res.SetStatus(transaction.StatusFailedTransactionStart)
		res.SetDescription("Unable to start Transaction. " + err.Error())
		return res
	}

	// Create each table
	for _, table := range tables {
		// Skip table if already exists
		_, exists := ctx.alreadyCreatedTables[table.Name()]
		if exists {
			continue
		}

		// Create table parent types
		for _, typ := range table.Dependencies.Types {
			// Skip type if already exists
			_, exists := ctx.alreadyCreatedTypes[typ.Name]
			if exists {
				continue
			}

			res := m.CreateType(typ)
			if res.Status != transaction.StatusSuccess {
				return res
			}

			ctx.alreadyCreatedTypes[typ.Name] = true
		}

		// Create table parent tables
		for _, subtable := range table.Dependencies.Tables {
			_, exists := ctx.alreadyCreatedTables[subtable.Name()]
			if exists {
				continue
			}

			res := m.CreateTables(ctx, subtable)
			if res.Status != transaction.StatusSuccess {
				return res
			}

			ctx.alreadyCreatedTables[subtable.Name()] = true
		}

		// Create table after creating its parent types and tables
		res := tx.Query(table.Queries.Create)
		if res.Status != transaction.StatusSuccess {
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

func (m *MigrationsManager) DropTables(tables ...*models.TableInfo) *transaction.Response {
	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse()
		res.SetStatus(transaction.StatusFailedTransactionStart)
		res.SetDescription("Unable to start Transaction. " + err.Error())
		return res
	}

	for _, table := range tables {
		res := tx.Query(table.Queries.Drop)
		if res.Status != transaction.StatusSuccess {
			return res
		}
	}

	return tx.Commit()
}

func (m *MigrationsManager) DropType(typ *models.TypeInfo) *transaction.Response {
	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse()
		res.SetStatus(transaction.StatusFailedTransactionStart)
		res.SetDescription("Unable to start Transaction. " + err.Error())
		return res
	}

	res := tx.Query(typ.Queries.Drop)
	if res.Status != transaction.StatusSuccess {
		return res
	}

	return tx.Commit()
}
