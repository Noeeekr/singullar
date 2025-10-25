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
	Use:   "relations [-f ...ENVIRONMENT_FILES] [--ignore-existing] [--recreate-existing] [ production | environment ]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Migrates the tables and roles to the ENVIRONMENT_DATABASE specified in the file",
	Run: func(cmd *cobra.Command, args []string) {
		// Configure migration settings
		borm.Settings().Migrations().Enable()
		ignoreExisting, _ := cmd.Flags().GetBool("ignore-existing")
		if ignoreExisting {
			borm.Settings().Migrations().IgnoreExisting()
			logs.Info.Println("[Ignore existing flag]: Existing relations won't stop the operations neither throw errors..")
		}
		recreateExisting, _ := cmd.Flags().GetBool("recreate-existing")
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

		path, _ := cmd.Flags().GetStringArray("environmentFiles")
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

		database.RegisterMigrationQueries(
			models.TableQuestionDifficulty.
				Insert("difficulty_level", "difficulty_name").
				Values(
					models.Fundamental, models.QuestionListDifficulties[models.Fundamental],
					models.Beginner, models.QuestionListDifficulties[models.Beginner],
					models.Intermediare, models.QuestionListDifficulties[models.Intermediare],
					models.Advanced, models.QuestionListDifficulties[models.Advanced],
					models.Expert, models.QuestionListDifficulties[models.Expert],
				),
		)

		// Migrate environment database relations
		if err := database.MigrateRelations(); err != nil {
			logs.Error.Fatal(err)
		}
		defer database.DB().Close()

		logs.Info.Println("[Migration finished]")
	},
}

func init() {
	commandutil.ConsumeFlagConfiguration(MigrationFlagsToken, RelationsCmd)
	commandutil.ConsumeFlagConfiguration(environment.EnvironmentFilesFlagToken, EnvironmentCmd)
}
