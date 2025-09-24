package migrate

import (
	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/util"
	"github.com/Noeeekr/singullar/server/util/environment"
	"github.com/spf13/cobra"
)

var RelationsCmd *cobra.Command = &cobra.Command{
	Use:   "relations [-f ...ENVIRONMENT_FILES] [--ignore-existing] [--recreate-existing] [ production | environment ]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Migrates the tables and roles to the ENVIRONMENT_DATABASE specified in the file",
	Run: func(cmd *cobra.Command, args []string) {
		// Configure migration settings
		borm.Settings().Migrations().Enable()
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

		// Configure environment
		if args[0] == "production" {
			environment.Settings().SetApplicationMode(environment.PRODUCTION)
		} else {
			environment.Settings().SetApplicationMode(environment.DEVELOPMENT)
		}

		path, _ := cmd.Flags().GetStringArray("environmentFiles")
		if res := environment.Parse(path...); res != nil {
			util.Error.Fatal("[Invalid environment file]:", res.String())
			return
		}

		// Connect to environment database
		database, err := borm.Connect(models.EnvironmentDatabase)
		if err != nil {
			util.Error.Fatal(err)
		}
		defer database.DB().Close()

		// Migrate environment database relations
		if err := database.MigrateRelations(); err != nil {
			util.Error.Fatal(err)
		}
		defer database.DB().Close()

		util.Info.Println("[Migration finished]")
	},
}

func init() {
	RelationsCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "Defines the path to the environment files containing the necessary environment variables if not already supplied in the environment")
	RelationsCmd.Flags().BoolP("ignore-existing", "i", false, "Doesn't throw errors if the ENVIRONMENT_DATABASE relation already exists.")
	RelationsCmd.Flags().BoolP("recreate-existing", "r", false, "Drop and recreate the relation if already exists.")
	RelationsCmd.MarkFlagsMutuallyExclusive("ignore-existing", "recreate-existing")
}
