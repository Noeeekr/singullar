package managers

import (
	"database/sql"
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/models"
)

// Should be put appended before other fields
var defaultFields = `
ID        INT         PRIMARY KEY,
CreatedAt TIMESTAMPTZ NOT NULL,
UpdatedAt TIMESTAMPTZ NOT NULL,
DeletedAt TIMESTAMPTZ NOT NULL,
`

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

func (m *MigrationsManager) CreateType(role models.TypeName) *Response {
	switch role {
	case models.RoleTypeName:
		return m.createTypeRoles()
	default:
		return &Response{
			Description: "Migration not registered.",
			Status:      StatusUnregisteredMigration,
		}
	}
}
func (m *MigrationsManager) CreateTable(table models.TableName, disableDefaults bool) *Response {
	switch table {
	case models.UsersTableName:
		return m.createTableUsers(disableDefaults)
	default:
		return &Response{
			Status:      StatusUnregisteredMigration,
			Description: "Migration Not Registered",
		}
	}
}
func (m *MigrationsManager) DropTable(table models.TableName) *Response {
	switch table {
	case models.UsersTableName:
		return m.Transaction(&Query{
			Query:  fmt.Sprintf("DROP TABLE IF EXISTS %s;", table),
			Method: DROP,
		})
	default:
		return &Response{
			Status:      StatusUnregisteredMigration,
			Description: "Unable to drop. " + string(table) + " is not a registered table.",
		}
	}
}
func (m *MigrationsManager) DropType(role models.TypeName) *Response {
	return m.Transaction(&Query{
		Method: DROP,
		Query: fmt.Sprintf(`
			DROP TYPE IF EXISTS %s;
		`, models.RoleTypeName),
	})
}

func (m *MigrationsManager) createTableUsers(disableDefaults bool) *Response {
	if disableDefaults {
		defaultFields = ""
	}

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS users (
			%s
			Name     VARCHAR(256)   NOT NULL,
			Email    VARCHAR(256)   NOT NULL UNIQUE,
			Password VARCHAR(256)   NOT NULL
		)
	`, defaultFields)

	tx, err := m.db.Begin()
	if err != nil {
		return &Response{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	if _, err := tx.Exec(query); err != nil {
		if err := tx.Rollback(); err != nil {
			return &Response{
				Status:      StatusFailedTransactionRollback,
				Description: "Migration transaction failed. Unable to rollback: " + err.Error(),
			}
		}
		return &Response{
			Status:      StatusFailedTransaction,
			Description: "Migration transaction failed. Rollback executed: " + err.Error(),
		}
	}

	if err := tx.Commit(); err != nil {
		return &Response{
			Status:      StatusFailedTransaction,
			Description: "Transaction Failed: " + err.Error(),
		}
	}

	return &Response{
		Status:      StatusSuccess,
		Description: "Migration transaction sucessfull.",
	}
}
func (m *MigrationsManager) createTypeRoles() *Response {
	var roles string = fmt.Sprintf(
		"'%s','%s','%s','%s'",
		models.RoleAdmin,
		models.RoleStudent,
		models.RoleSupervisor,
		models.RoleTeacher,
	)

	query := fmt.Sprintf(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT * FROM pg_type WHERE typname = '%s') THEN
					CREATE TYPE %s AS ENUM (%s);
				END IF;
			END $$;
		`, models.RoleTypeName, models.RoleTypeName, roles)

	parsedQuery := &Query{
		Method: CREATE,
		Query:  query,
	}

	return m.Transaction(parsedQuery)
}
