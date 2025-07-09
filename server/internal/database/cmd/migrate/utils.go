package migrate

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

type Utils struct {
	migrations *migrations.Migrations
}

func (u *Utils) MigrateTables(configuration *migrations.Configuration) bool {
	res := u.migrations.CreateTables(
		configuration,
		nil,
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
