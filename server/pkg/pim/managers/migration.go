package managers

import (
	"database/sql"
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/models"
)

type MigrationsManager struct {
	*TransactionManager
}

func NewMigrationManager(db *sql.DB) *MigrationsManager {
	return &MigrationsManager{
		&TransactionManager{
			db: db,
		},
	}
}

func (m *MigrationsManager) CreateType(typ *models.TypeInfo) *Response {
	tx, err := m.StartTransation()
	if err != nil {
		return &Response{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	res := tx.Query(&Query{
		Query:   typ.Queries.Create(),
		Returns: false,
	})

	if res.Status != StatusSuccess {
		return res.Response
	}

	return tx.Commit().Response
}

func (m *MigrationsManager) CreateTable(table *models.TableInfo) *Response {
	tx, err := m.StartTransation()
	if err != nil {
		return &Response{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	for _, typ := range table.Dependencies.Types {
		if res := tx.Query(&Query{
			Query:   typ.Queries.Create(),
			Returns: false,
		}); res.Status != StatusSuccess {
			return res.Response
		}
	}

	for _, subtable := range table.Dependencies.Tables {
		if res := tx.Query(&Query{
			Query:   subtable.Queries.Create,
			Returns: false,
		}); res.Status != StatusSuccess {
			return res.Response
		}
	}

	res := tx.Query(&Query{
		Query:   table.Queries.Create,
		Returns: false,
	})

	if res.Status != StatusSuccess {
		return res.Response
	}

	return tx.Commit().Response
}

func (m *MigrationsManager) DropTable(table *models.TableInfo) *Response {
	tx, err := m.StartTransation()
	if err != nil {
		return &Response{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	res := tx.Query(&Query{
		Query:   fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", table.Name()),
		Returns: false,
	})

	if res.Status != StatusSuccess {
		return res.Response
	}

	return tx.Commit().Response
}

func (m *MigrationsManager) DropType(Type *models.TypeInfo) *Response {
	tx, err := m.StartTransation()
	if err != nil {
		return &Response{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	res := tx.Query(&Query{
		Query:   fmt.Sprintf(`DROP TYPE IF EXISTS %s;`, Type.Name()),
		Returns: false,
		Args:    []any{},
	})

	if res.Status != StatusSuccess {
		return res.Response
	}

	return tx.Commit().Response
}
