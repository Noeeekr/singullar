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
		models.UsersTable,
		models.ClassesTable,
		models.InstitutionsTable,
		models.NotificationsTable,
		models.UsersClassesTable,
		models.UsersNotificationsTable,
	).Response
	if res != nil {
		fmt.Println(res.ParseToString())
		return true
	}
	return false
}
