package migrate

import (
	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/common/commandutil"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/spf13/cobra"
)

const (
	RecreateExistingFlagName string = "recreate-existing"
	IgnoreExistingFlagName   string = "ignore-existing"
)

// Swapped in init
var MigrationFlagsToken = commandutil.RegisterConfigurationToken()

var EnvironmentCmd *cobra.Command = &cobra.Command{
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	Use:   "environment [-e ...ENVIRONMENT_FILES] { production | development }",
	Short: "Migrates the database and users of that specific environment",
	Args:  cobra.MinimumNArgs(1),
	Long:  "Migrates the database and users of that specific environment.",
	Run: func(cmd *cobra.Command, args []string) {
		ConfigureMigration(cmd, args)

		// Connect to postgres
		postgresConnection, res := connections.ScanEnvironmentForConnection(connections.POSTGRES)
		if res != nil {
			logs.Error.Fatal(res.String())
		}

		environmentConnection, res := connections.ScanEnvironmentForConnection(environment.Settings().ApplicationMode())
		if res != nil {
			logs.Error.Fatal(res.String())
		}

		// Register databases
		postgresUser := borm.RegisterUser(postgresConnection.User(), postgresConnection.Password())
		postgresDatabase := borm.RegisterDatabase(postgresConnection.Database(), postgresConnection.Host(), postgresUser)

		database, err := borm.Connect(postgresDatabase)
		if err != nil {
			logs.Error.Fatal(err)
		}
		if err := database.DB().Ping(); err != nil {
			logs.Error.Fatal("Failed to ping database:", err.Error())
		}
		defer database.DB().Close()

		environmentUser := borm.RegisterUser(environmentConnection.User(), environmentConnection.Password())
		environmentDatabase := borm.RegisterDatabase(environmentConnection.Database(), environmentConnection.Host(), environmentUser)

		borm.Settings().Migrations().Enable()

		err = database.MigrateUsers(environmentUser)
		if err != nil {
			logs.Error.Fatal(err)
		}

		createdDatabase, err := database.MigrateDatabase(environmentDatabase)
		if err != nil {
			logs.Error.Fatal(err)
		}
		if err := createdDatabase.DB().Ping(); err != nil {
			logs.Error.Fatal("Failed to ping database:", err.Error())
		}
		defer createdDatabase.DB().Close()

		logs.Info.Println("[Environment migrated successfully]")
	},
}

func ConfigureMigration(cmd *cobra.Command, args []string) {
	if args[0] == "production" {
		environment.Settings().SetApplicationMode(environment.PRODUCTION)
	} else {
		environment.Settings().SetApplicationMode(environment.DEVELOPMENT)
	}
	ignoreExisting, _ := cmd.Flags().GetBool(IgnoreExistingFlagName)
	if ignoreExisting {
		borm.Settings().Migrations().IgnoreExisting()
		logs.Info.Println("[Ignore existing flag]: Existing relations won't stop the operations neither throw errors..")
	}

	recreateExisting, _ := cmd.Flags().GetBool(RecreateExistingFlagName)
	if recreateExisting {
		borm.Settings().Migrations().RecreateExisting()
		logs.Info.Println("[Recreate existing flag]: Existing relations will be dropped and recreated...")
	}

	path, _ := cmd.Flags().GetStringArray(environment.EnvironmentFilesFlagName)
	if err := environment.Parse(path...); err != nil {
		logs.Error.Fatal("[Invalid environment file]: ", err.String())
	}
}
func init() {
	commandutil.RegisterFlagConfiguration(MigrationFlagsToken, func(c *cobra.Command) {
		c.Flags().BoolP(IgnoreExistingFlagName, string(IgnoreExistingFlagName[0]), false, "Doesn't throw errors and proceed if the database relation already exists.")
		c.Flags().BoolP(RecreateExistingFlagName, string(RecreateExistingFlagName[0]), false, "Drop and recreate the relation if already exists.")
		c.MarkFlagsMutuallyExclusive(IgnoreExistingFlagName, RecreateExistingFlagName)
	})
	commandutil.ConsumeFlagConfiguration(MigrationFlagsToken, EnvironmentCmd)
	commandutil.ConsumeFlagConfiguration(environment.EnvironmentFilesFlagToken, EnvironmentCmd)
}
