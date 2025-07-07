package migrate

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

type MigrateCommandUtils struct {
	migrations *migrations.MigrationsManager
}

func (u *MigrateCommandUtils) MigrateTables() bool {
	res := u.migrations.CreateTables(nil,
		models.TablesInfo.Users,
		models.TablesInfo.Classes,
		models.TablesInfo.Institutions,
		models.TablesInfo.Notifications,
		models.TablesInfo.UsersClasses,
		models.TablesInfo.UsersNotifications,
	).Response
	if res != nil {
		fmt.Println(res.ParseToString())
		return true
	}
	return false
}

func (u *MigrateCommandUtils) CreateEnvironmentDatabases(databases ...string) bool {
	res := u.migrations.CreateDatabases(databases...)
	if res != nil && res.Status != common.StatusFound {
		fmt.Println(res.ParseToString())
		return true
	}
	return false
}
func (u *MigrateCommandUtils) CreateEnvironmentUsers(conn connections.Connection) bool {
	user := models.NewDatabaseUser(conn.User(), conn.Password(), conn.Database())

	res := u.migrations.CreateUsers(user)
	if res != nil && res.Status != common.StatusFound {
		fmt.Println(res.ParseToString())
		return true
	}

	return false
}
func (u *MigrateCommandUtils) GrantAllPrivilegesOnDatabase(users ...*models.CreateDatabaseUser) bool {
	res := u.migrations.GrantAllPrivilegesOnDatabase(users...)
	if res != nil && res.Status != common.StatusFound {
		fmt.Println(res.ParseToString())
		return true
	}
	return false
}
