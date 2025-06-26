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

func (m *MigrationsManager) CreateType(role models.TypeName) *Response {
	switch role {
	case models.UserRoleName:
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
	case models.UsersTable.TableName():
		return m.createTableUsers()
	default:
		return &Response{
			Status:      StatusUnregisteredMigration,
			Description: "Migration Not Registered",
		}
	}
}
func (m *MigrationsManager) DropTable(table models.TableName) *Response {
	switch table {
	case models.NotificationsTable.TableName():
	case models.InstitutionsTable.TableName():
	case models.ClassesTable.TableName():
	case models.UsersTable.TableName():
		return m.Transaction(&Query{
			Query:  fmt.Sprintf("DROP TABLE IF EXISTS %s;", table),
			Method: DROP,
		})
	}

	return &Response{
		Status:      StatusUnregisteredMigration,
		Description: "Unable to drop. " + string(table) + " is not a registered table.",
	}
}

func (m *MigrationsManager) DropType(role models.TypeName) *Response {
	return m.Transaction(&Query{
		Method: DROP,
		Query: fmt.Sprintf(`
			DROP TYPE IF EXISTS %s;
		`, models.UserRoleName),
	})
}

func (m *MigrationsManager) createTableUsers() *Response {
	return m.Transaction(&Query{
		Query:  models.UsersTable.TableQuery(),
		Method: CREATE,
	})
}
func (m *MigrationsManager) createTypeRoles() *Response {
	var roles string = fmt.Sprintf(
		"'%s','%s','%s','%s','%s'",
		models.Admin,
		models.Student,
		models.Supervisor,
		models.Teacher,
		models.Unknown,
	)

	query := fmt.Sprintf(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT * FROM pg_type WHERE typname = '%s') THEN
					CREATE TYPE %s AS ENUM (%s);
				END IF;
			END $$;
		`, models.UserRoleName, models.UserRoleName, roles)

	parsedQuery := &Query{
		Method: CREATE,
		Query:  query,
	}

	return m.Transaction(parsedQuery)
}
