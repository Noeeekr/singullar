package migrate

import (
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/spf13/cobra"
)

var RelationsCmd *cobra.Command = &cobra.Command{
	Use:   "relations [-f ...ENVIRONMENT_FILES] [--ignore-existing] [ production | environment ]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Migrates the tables and roles to the database specified in the file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetStringArray("environmentFiles")
		if res := environment.Parse(path...); res != nil {
			logs.Error.Fatal("[Invalid environment file]:", res.ParseToString())
			return
		}

		connection, res := connections.ScanEnvironmentForConnection(connections.ConnectionEnvironment(args[0]))
		if res != nil {
			logs.Error.Fatal(res.ParseToString())
			return
		}

		db, res := connections.Connect(connections.NewConnectionString(connection.User(), connection.Password(), connection.Host(), connection.Database()))
		if res != nil {
			logs.Error.Fatal(res.ParseToString())
			return
		}
		defer db.Close()

		utils := Utils{migrations: migrations.New(db)}

		if utils.MigrateTables() {
			return
		}

		logs.Info.Println("[Migration finished]")
	},
}

func init() {
	RelationsCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "Defines the path to the environment files containing the necessary environment variables if not already supplied in the environment")
	// not implemented
	RelationsCmd.Flags().BoolP("ignore-existing", "i", false, "Doesn't throw errors if the database relation already exists.")

	// Not implemented RelationsCmd.Flags().BoolP("recreate-existing", "r", false, "Drop and recreate the relation if already exists.")
}
