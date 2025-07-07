package cmd

import (
	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/server"
	"github.com/Noeeekr/singullar/server/internal/api/server/types"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

var startCmd *cobra.Command = &cobra.Command{
	Use:   "start [FLAGS]",
	Short: "Start command starts the main API. The flags are used if the environment variables are not set",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("enable-debug")
		if !debug {
			gin.SetMode(gin.ReleaseMode)
		}

		mode, _ := cmd.Flags().GetString("mode")
		mode, err := environment.SetIfNotEmpty("API_ENVIRONMENT", mode)
		if err != nil {
			logs.Info.Fatal(err.Error())
			return
		}

		port, _ := cmd.Flags().GetString("port")
		port, err = environment.SetIfNotEmpty("API_PORT", port)
		if err != nil {
			logs.Info.Fatal(err.Error())
			return
		}

		// Execute migrations if enable-migrations is present
		shouldMigrate, _ := cmd.Flags().GetBool("enable-migrations")
		if shouldMigrate {
			res := Migrate()
			if res != nil {
				logs.Info.Fatal(res.ParseToString())
			}
		}

		if res := StartApi(port, mode); res != nil {
			logs.Info.Fatal(res.ParseToString())
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Bool("enable-debug", false, "Defines if the server should start in debug mode. Defaults to false")
	startCmd.Flags().Bool("enable-migrations", false, "Defines if the server should start with migrations. Defaults to false")

	startCmd.Flags().String("port", "80", "Defines the port the server will listen to.")
	startCmd.Flags().String("mode", "", "The environment to migrate on. Defaults to development.")

	startCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "The path to the files containing the required environment variables.")
}

func StartApi(port, mode string) *common.Response {
	db, res := connections.ConnectWithEnvironment(connections.ConnectionEnvironment(mode))
	if res != nil {
		return res
	}

	var env types.ApiEnvironment
	if err := environment.Scan(&env); err != nil {
		return res
	}

	router, err := server.PrepareRouter(db, &env)
	if err != nil {
		return res
	}

	server := server.New().
		WithAddr(":" + port).
		WithErrLogger(logs.Error).
		WithRouter(router)

	logs.Info.Printf("Server is running on http://localhost:%s", port)
	return server.ListenAndServe()
}

func Migrate() *common.Response {
	db, err := connections.ConnectWithEnvironment(connections.Postgres)
	if err != nil {
		return err
	}

	mig := migrations.New(db)
	tx := mig.StartTransaction()
	if tx.Response != nil {
		return tx.Response
	}

	tx = mig.CreateTables(
		tx,
		models.TablesInfo.Institutions,
		models.TablesInfo.Users,
		models.TablesInfo.Classes,
		models.TablesInfo.Notifications,
		models.TablesInfo.UsersClasses,
		models.TablesInfo.UsersNotifications,
	)
	if tx.Response != nil {
		return tx.Response
	}
	return nil
}
