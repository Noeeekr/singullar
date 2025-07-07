package migrate

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/spf13/cobra"
)

var EnvironmentCmd *cobra.Command = &cobra.Command{
	Use:   "environment [-e ENVIRONMENT_FILE_PATH] { production | development }",
	Short: "Migrates the database and users of that specific environment",
	Args:  cobra.MinimumNArgs(1),
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("environmentFile")

		if err := environment.Parse(path); err != nil {
			fmt.Println("[Invalid environment file]: ", err.ParseToString())
		}

		// OPEN POSTGRES DB TO CREATED REQUIRED USERS AND DATABASES BEFORE MIGRATING
		pgConn, res := connections.ScanEnvironmentForConnection(connections.Postgres)
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}

		db, res := connections.Connect(connections.ParseConnectionString(pgConn))
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}
		defer db.Close()

		conn, res := connections.ScanEnvironmentForConnection(connections.ConnectionEnvironment(args[0]))
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}
		utils := MigrateCommandUtils{migrations: migrations.New(db)}

		if utils.CreateEnvironmentDatabases(conn.Database()) {
			return
		}

		if utils.CreateEnvironmentUsers(conn) {
			res := utils.migrations.DropDatabases(conn.Database())
			if res != nil {
				fmt.Println(res.ParseToString())
				return
			}
		}

		if utils.GrantAllPrivilegesOnDatabase(models.NewDatabaseUser(conn.User(), conn.Password(), conn.Database())) {
			res = utils.migrations.DropDatabases(conn.Database())
			if res != nil {
				fmt.Println(res.ParseToString())
			}

			res = utils.migrations.DropUsers(conn.User())
			if res != nil {
				fmt.Println(res.ParseToString())
			}
			return
		}
		fmt.Println("[Environment migrated successfully]")
	},
}

func init() {
	EnvironmentCmd.Flags().StringP("environmentFile", "e", "", "The path to the file containing the connection string.")
	EnvironmentCmd.MarkFlagRequired("environmentFile")
}
