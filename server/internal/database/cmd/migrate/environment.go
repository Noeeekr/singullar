package migrate

import (
	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/util"
	"github.com/Noeeekr/singullar/server/util/environment"
	"github.com/spf13/cobra"
)

var EnvironmentCmd *cobra.Command = &cobra.Command{
	Use:   "environment [-e ...ENVIRONMENT_FILES] { production | development }",
	Short: "Migrates the database and users of that specific environment",
	Args:  cobra.MinimumNArgs(1),
	Long:  "Migrates the database and users of that specific environment.",
	Run: func(cmd *cobra.Command, args []string) {
		ignoreExisting, _ := cmd.Flags().GetBool("ignore-existing")
		if ignoreExisting {
			borm.Settings().Migrations().IgnoreExisting()
			util.Info.Println("[Ignore existing flag]: Existing relations won't stop the operations neither throw errors..")
		}

		recreateExisting, _ := cmd.Flags().GetBool("recreate-existing")
		if recreateExisting {
			borm.Settings().Migrations().RecreateExisting()
			util.Info.Println("[Recreate existing flag]: Existing relations will be dropped and recreated...")
		}

		path, _ := cmd.Flags().GetStringArray("environmentFiles")
		if err := environment.Parse(path...); err != nil {
			util.Error.Fatal("[Invalid environment file]: ", err.String())
		}

		if args[0] == "production" {
			environment.Settings().SetApplicationMode(environment.PRODUCTION)
		} else {
			environment.Settings().SetApplicationMode(environment.DEVELOPMENT)
		}

		// Get connection information
		postgresConnection, res := connections.ScanEnvironmentForConnection(connections.POSTGRES)
		if res != nil {
			util.Error.Fatal(res.String())
		}

		environmentConnection, res := connections.ScanEnvironmentForConnection(environment.Settings().ApplicationMode())
		if res != nil {
			util.Error.Fatal(res.String())
		}

		// Register databases
		postgresUser := borm.RegisterUser(postgresConnection.User(), postgresConnection.Password())
		postgresDatabase := borm.RegisterDatabase(postgresConnection.Database(), postgresConnection.Host(), postgresUser)

		database, err := borm.Connect(postgresDatabase)
		if err != nil {
			util.Error.Fatal(err)
		}
		if err := database.DB().Ping(); err != nil {
			util.Error.Fatal("Failed to ping database:", err.Error())
		}
		defer database.DB().Close()

		environmentUser := borm.RegisterUser(environmentConnection.User(), environmentConnection.Password())
		environmentDatabase := borm.RegisterDatabase(environmentConnection.Database(), environmentConnection.Host(), environmentUser)

		borm.Settings().Migrations().Enable()

		err = database.MigrateUsers(environmentUser)
		if err != nil {
			util.Error.Fatal(err)
		}

		createdDatabase, err := database.MigrateDatabase(environmentDatabase)
		if err != nil {
			util.Error.Fatal(err)
		}
		if err := createdDatabase.DB().Ping(); err != nil {
			util.Error.Fatal("Failed to ping database:", err.Error())
		}
		defer createdDatabase.DB().Close()

		util.Info.Println("[Environment migrated successfully]")
	},
}

func init() {
	EnvironmentCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "Defines the path to the environment files containing the necessary environment variables if not already supplied in the environment")
	EnvironmentCmd.Flags().BoolP("ignore-existing", "i", false, "Doesn't throw errors and proceed if the database relation already exists.")
	// Not implemented
	EnvironmentCmd.Flags().BoolP("recreate-existing", "r", false, "Drop and recreate the relation if already exists.")
	EnvironmentCmd.MarkFlagsMutuallyExclusive("ignore-existing", "recreate-existing")
}
