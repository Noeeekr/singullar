package cmd

import (
	"os"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/server"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/gin-gonic/gin"

	"github.com/Noeeekr/singullar/server/internal/database/cmd/migrate"

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
		ignoreExisting, _ := cmd.Flags().GetBool("ignore-existing")
		recreateExisting, _ := cmd.Flags().GetBool("recreate-existing")
		if shouldMigrate {
			flags := []string{}
			if ignoreExisting {
				flags = append(flags, "--ignore-existing")
			}
			if recreateExisting {
				flags = append(flags, "--recreate-existing")
			}
			if res := Migrate(mode, flags...); res != nil {
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

	startCmd.Flags().Bool("recreate-existing", false, "Drop and recreate relations in the database migration if already exists.")
	startCmd.Flags().Bool("ignore-existing", false, "Doesn't throw errors and proceed if the database relation already exists.")
	startCmd.MarkFlagsMutuallyExclusive("ignore-existing", "recreate-existing")

	startCmd.Flags().String("port", "80", "Defines the port the server will listen to.")
	startCmd.Flags().String("mode", "development", "The environment to migrate on. Defaults to development.")

	startCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "The path to the files containing the required environment variables.")
}

func StartApi(port, mode string) *common.Response {
	db, res := connections.ConnectWithEnvironment(connections.ConnectionEnvironment(mode))
	if res != nil {
		return res
	}

	var env types.Environment
	if res := environment.Scan(&env); res != nil {
		return res
	}

	router, err := server.PrepareRouter(db, &env)
	if err != nil {
		return common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusInternalError)
	}

	server := server.New(router, ":"+port).
		WithErrLogger(logs.Error)

	logs.Info.Printf("Server is running on http://localhost:%s", port)
	err = server.ListenAndServe()

	return common.NewResponse().
		WithDescription(err.Error()).
		WithStatus(common.StatusInternalError)
}

func Migrate(mode string, flags ...string) *common.Response {
	args := []string{"./api", mode}
	args = append(args, flags...)

	os.Args = args

	err := migrate.EnvironmentCmd.Execute()
	if err != nil {
		return common.NewResponse().WithDescription(err.Error()).WithStatus(common.StatusInternalError)
	}

	os.Args = args

	err = migrate.RelationsCmd.Execute()
	if err != nil {
		return common.NewResponse().WithDescription(err.Error()).WithStatus(common.StatusInternalError)
	}

	return nil
}
