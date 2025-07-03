package cmd

import (
	"github.com/Noeeekr/singullar/server/common/configs"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/server"
	"github.com/Noeeekr/singullar/server/internal/api/server/types"
	"github.com/Noeeekr/singullar/server/internal/database"
	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

var startCmd *cobra.Command = &cobra.Command{
	Use:   "start",
	Short: "Starts the api.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		if debug {
			gin.SetMode(gin.ReleaseMode)
		}

		apiEnvironmentPath, _ := cmd.Flags().GetString("environment")
		postgresEnvironmentPath, _ := cmd.Flags().GetString("postgres")

		var envs []string = []string{apiEnvironmentPath, postgresEnvironmentPath}
		if err := configs.Parse(envs...); err != nil {
			logs.Info.Fatalf(err.ParseToError().Error())
		}

		var pgEnv database.PostgresEnvironment
		if err := configs.Scan(&pgEnv); err != nil {
			logs.Info.Fatalf(err.ParseToError().Error())
		}

		var apiEnv types.ApiEnvironment
		if err := configs.Scan(&apiEnv); err != nil {
			logs.Info.Fatalf(err.ParseToError().Error())
		}

		db, err := database.Connect(pgEnv.POSTGRES_CONNECTION_STRING)
		if err != nil {
			logs.Info.Fatalf(err.Error())
		}

		router, err := server.PrepareRouter(db, &apiEnv)
		if err != nil {
			logs.Info.Fatalf(err.Error())
		}

		server := server.New().
			WithAddr(":" + apiEnv.Port).
			WithErrLogger(logs.Error).
			WithRouter(router)

		logs.Info.Printf("Server is running on http://localhost:%s", apiEnv.Port)
		if err != server.ListenAndServe() {
			logs.Error.Fatal(err)
		}
	},
}

func init() {
	startCmd.Flags().StringP("postgres", "p", "", "Defines the path to the environment file containing the Postgres connection string.")
	startCmd.Flags().StringP("environment", "e", "", "Defines the path to the environment file containing the API configurations.")
	startCmd.Flags().BoolP("debug", "d", false, "Defines if the server should start in debug mode.")

	startCmd.MarkFlagRequired("environment")
	startCmd.MarkFlagRequired("postgres")
}
