package cmd

import (
	"fmt"
	"os"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/api/server"
	"github.com/Noeeekr/singullar/server/internal/api/server/handlers"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/util"
	"github.com/Noeeekr/singullar/server/util/environment"
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
		mode = environment.OverrideEmpty("API_ENVIRONMENT", mode)

		port, _ := cmd.Flags().GetString("port")
		port = environment.OverrideEmpty("API_PORT", port)

		domain, _ := cmd.Flags().GetString("domain")
		domain = environment.OverrideEmpty("API_DOMAIN", domain)

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
				util.Info.Fatal(res.String())
			}
		}

		if err := StartApi(domain, port, mode); err != nil {
			util.Error.Fatal(err.Error())
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

	startCmd.Flags().String("mode", "development", "The environment to migrate on. Defaults to development.")
	startCmd.Flags().String("port", "80", "Defines the port the server will listen to.")
	startCmd.Flags().String("domain", "", "Defines the server host portion of the URI")

	startCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "The path to the files containing the required environment variables.")
}

func StartApi(domain, port, mode string) error {
	commiter, err := borm.Connect(models.EnvironmentDatabase)
	if err != nil {
		return err
	}

	var env types.Environment
	if err := environment.Scan(&env); err != nil {
		return err.ParseToError()
	}

	router, err := server.PrepareRouter(handlers.New(operations.New(commiter), &env), &env)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%s", domain, port)
	server := server.New(router, addr).
		WithErrLogger(util.Error)

	util.Info.Printf("Mode: %s", mode)
	util.Info.Printf("Domain: %s", domain)
	util.Info.Printf("Database: %s", commiter.Name)
	util.Info.Printf("Address: %s\n", addr)
	return server.ListenAndServe()
}

func Migrate(mode string, flags ...string) *util.Response {
	args := []string{"./api", mode}
	args = append(args, flags...)

	os.Args = args

	err := migrate.EnvironmentCmd.Execute()
	if err != nil {
		return util.NewResponse().WithDescription(err.Error()).WithStatus(util.StatusInternalError)
	}

	os.Args = args

	err = migrate.RelationsCmd.Execute()
	if err != nil {
		return util.NewResponse().WithDescription(err.Error()).WithStatus(util.StatusInternalError)
	}

	return nil
}
