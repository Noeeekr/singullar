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
	return m.Transaction(&Query{
		Query:  typ.Query(),
		Method: CREATE,
	})
}
func (m *MigrationsManager) CreateTable(table *models.TableInfo) *Response {
	tx, err := m.db.Begin()
	if err != nil {
		return &Response{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	for _, typ := range table.Dependencies.Types {
		res := m.transaction(tx, typ.Query())
		if res.Status != StatusSuccess {
			return res
		}
	}

	for _, subtable := range table.Dependencies.Tables {
		res := m.transaction(tx, subtable.TableQuery())
		if res.Status != StatusSuccess {
			return res
		}
	}

	res := m.transaction(tx, table.TableQuery())
	if res.Status != StatusSuccess {
		return res
	}

	if err := tx.Commit(); err != nil {
		return &Response{
			Status:      StatusFailedTransaction,
			Description: "Unable to commit transaction. " + err.Error(),
		}
	}

	return &Response{
		Status:      StatusSuccess,
		Description: "Transaction commited successfully.",
	}
}

func (m *MigrationsManager) DropTable(table *models.TableInfo) *Response {
	return m.Transaction(&Query{
		Query:  fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", table.TableName()),
		Method: DROP,
	})
}

func (m *MigrationsManager) DropType(Type *models.TypeInfo) *Response {
	return m.Transaction(&Query{
		Method: DROP,
		Query: fmt.Sprintf(`
			DROP TYPE IF EXISTS %s;
		`, Type.Name()),
	})
}
