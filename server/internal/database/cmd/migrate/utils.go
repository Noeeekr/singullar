package migrate

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

type Utils struct {
	migrations *migrations.Migrations
}

func (u *Utils) MigrateTables() bool {
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

func (u *Utils) CreateEnvironmentDatabases(requests []*migrations.RequestCreateDatabase, configuration *migrations.Configuration) *common.Response {
	res := u.migrations.CreateDatabases(requests, configuration)
	if res != nil && res.Status != common.StatusFound {
		return res
	}
	return res
}
func (u *Utils) CreateEnvironmentUsers(conn connections.Connection) bool {
	user := models.NewDatabaseUser(conn.User(), conn.Password(), conn.Database())

	res := u.migrations.CreateUsers(user)
	if res != nil && res.Status != common.StatusFound {
		fmt.Println(res.ParseToString())
		return true
	}

	return false
}
func (u *Utils) GrantAllPrivilegesOnDatabase(users ...*models.CreateDatabaseUser) *common.Response {
	res := u.migrations.GrantAllPrivilegesOnDatabase(users...)
	if res != nil && res.Status != common.StatusFound {
		return res
	}
	return nil
}
