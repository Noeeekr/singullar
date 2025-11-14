package cmd

import (
	"fmt"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/server"
	"github.com/Noeeekr/singullar/server/internal/api/server/handlers"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/common/commandutil"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
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
		if shouldMigrate {
			if res := Migrate(mode); res != nil {
				logs.Info.Fatal(res.String())
			}
		}

		if err := start(domain, port, mode); err != nil {
			logs.Error.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	commandutil.ConsumeFlagConfiguration(environment.EnvironmentFilesFlagToken, startCmd)
	commandutil.ConsumeFlagConfiguration(migrate.MigrationFlagsToken, startCmd)

	startCmd.Flags().Bool("enable-debug", false, "Defines if the server should start in debug mode. Defaults to false")
	startCmd.Flags().Bool("enable-migrations", false, "Defines if the server should start with migrations. Defaults to false")

	startCmd.Flags().String("mode", "development", "The environment to migrate on. Defaults to development.")
	startCmd.Flags().String("port", "80", "Defines the port the server will listen to.")
	startCmd.Flags().String("domain", "", "Defines the server host portion of the URI")
}

func start(domain, port, mode string) error {
	commiter, err := borm.Connect(models.EnvironmentDatabase)
	if err != nil {
		return err
	}

	var env types.Environment
	if err := environment.Scan(&env); err != nil {
		return err.ParseToError()
	}

	operator := operations.New(commiter)
	handlers := handlers.New(operator, &env)
	router, err := server.PrepareRouter(handlers, &env)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%s", domain, port)
	server := server.
		New(router, addr).
		RegisterErrorLogger(logs.Error)

	logs.Info.Printf("Mode: %s", mode)
	logs.Info.Printf("Domain: %s", domain)
	logs.Info.Printf("Database: %s", commiter.Name)
	logs.Info.Printf("Address: %s\n", addr)
	return server.ListenAndServe()
}

func Migrate(mode string, flags ...string) *common.Response {
	err := migrate.EnvironmentCmd.Execute()
	if err != nil {
		return common.NewResponse().WithDescription(err.Error()).WithStatus(common.StatusInternalError)
	}

	err = migrate.RelationsCmd.Execute()
	if err != nil {
		return common.NewResponse().WithDescription(err.Error()).WithStatus(common.StatusInternalError)
	}

	return nil
}
