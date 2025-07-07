package migrate

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/spf13/cobra"
)

var RelationsCmd *cobra.Command = &cobra.Command{
	Use:   "relations [-e ENVIRONMENT_FILE_PATH] [DATABASE]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Migrates the tables and roles to the database specified in the file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("environmentFile")
		if res := environment.Parse(path); res != nil {
			fmt.Println("[Invalid environment file]:", res.ParseToString())
			return
		}

		connection, res := connections.ScanEnvironmentForConnection(connections.Postgres)
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}

		db, res := connections.Connect(connections.NewConnectionString(connection.User(), connection.Password(), connection.Host(), args[0]))
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}

		utils := MigrateCommandUtils{migrations: migrations.New(db)}

		if utils.MigrateTables() {
			return
		}

		fmt.Println("[Migration finished]")
	},
}

func init() {
	RelationsCmd.Flags().StringP("environmentFile", "e", "", "The path to the file containing the connection string.")
	RelationsCmd.MarkFlagRequired("environmentFile")
}
