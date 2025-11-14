package migrate

import (
	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/common/commandutil"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/spf13/cobra"
)

var RelationsCmd *cobra.Command = &cobra.Command{
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	Use:   "relations [-f ...ENVIRONMENT_FILES] [--ignore-existing] [--recreate-existing] [ production | environment ]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Migrates the tables and roles to the ENVIRONMENT_DATABASE specified in the file",
	Run: func(cmd *cobra.Command, args []string) {
		// Configure migration settings
		borm.Settings().Migrations().Enable()
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

		// Configure environment
		if args[0] == "production" {
			environment.Settings().SetApplicationMode(environment.PRODUCTION)
		} else {
			environment.Settings().SetApplicationMode(environment.DEVELOPMENT)
		}

		path, _ := cmd.Flags().GetStringArray(environment.EnvironmentFilesFlagName)
		if res := environment.Parse(path...); res != nil {
			logs.Error.Fatal("[Invalid environment file]:", res.String())
			return
		}

		// Connect to environment database
		database, err := borm.Connect(models.EnvironmentDatabase)
		if err != nil {
			logs.Error.Fatal(err)
		}
		defer database.DB().Close()

		// Set population queries
		database.RegisterMigrationQueries(models.TableQuestionDifficulty.
			Insert("difficulty_name", "difficulty_level").
			Values(
				"Fundamental", 1,
				"Iniciante", 2,
				"Intermediario", 3,
				"Intermediario+", 4,
				"Intermediario++", 5,
				"Avançado", 6,
				"Avançado+", 7,
				"Avançado++", 8,
				"Profissional", 9,
				"Expert", 10,
			),
		)

		// Migrate environment database relations
		if err := database.MigrateRelations(); err != nil {
			logs.Error.Fatal(err.Error())
		}
		defer database.DB().Close()

		logs.Info.Println("[Migration finished]")
	},
}

func init() {
	commandutil.ConsumeFlagConfiguration(MigrationFlagsToken, RelationsCmd)
	commandutil.ConsumeFlagConfiguration(environment.EnvironmentFilesFlagToken, RelationsCmd)
}
