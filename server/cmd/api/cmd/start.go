package cmd

import (
	"os"

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
		if !debug {
			gin.SetMode(gin.ReleaseMode)
		}

		port, _ := cmd.Flags().GetString("port")
		if os.Getenv("PORT") == "" {
			os.Setenv("PORT", port)
		}

		var db_env database.PostgresEnvironment
		if err := configs.Scan(&db_env); err != nil {
			logs.Info.Fatal(err.ParseToError())
		}

		db, err := database.Connect(db_env.POSTGRES_CONNECTION_STRING)
		if err != nil {
			logs.Info.Fatal(err.Error())
		}

		var env types.ApiEnvironment
		if err := configs.Scan(&env); err != nil {
			logs.Info.Fatal(err.ParseToError())
		}

		router, err := server.PrepareRouter(db, &env)
		if err != nil {
			logs.Info.Fatal(err.Error())
		}

		server := server.New().
			WithAddr(":" + os.Getenv("PORT")).
			WithErrLogger(logs.Error).
			WithRouter(router)

		logs.Info.Printf("Server is running on http://localhost:%s", os.Getenv("PORT"))
		if err != server.ListenAndServe() {
			logs.Error.Fatal(err)
		}
	},
}
