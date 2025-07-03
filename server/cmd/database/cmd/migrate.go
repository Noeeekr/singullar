package cmd

import (
	"github.com/Noeeekr/singullar/server/common/configs"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/database"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/spf13/cobra"
)

var migrateCmd *cobra.Command = &cobra.Command{
	Use:   "migrate --environment [FILEPATH]",
	Short: "Allows migrating tables to Postgresql database.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("environment")

		if err := configs.Parse(path); err != nil {
			logs.Error.Fatal(err.ParseToError())
		}

		var env database.PostgresEnvironment
		if err := configs.Scan(&env); err != nil {
			logs.Error.Fatal(err.ParseToError())
		}

		db, err := database.Connect(env.POSTGRES_CONNECTION_STRING)
		if err != nil {
			logs.Error.Fatal(err.Error())
		}

		mig := migrations.New(db)
		res := mig.CreateTables(nil,
			models.TablesInfo.Users,
			models.TablesInfo.Classes,
			models.TablesInfo.Institutions,
			models.TablesInfo.Notifications,
			models.TablesInfo.UsersClasses,
			models.TablesInfo.UsersNotifications,
		)
		if res != nil {
			logs.Error.Fatal(res.ParseToError().Error())
		}
	},
}
