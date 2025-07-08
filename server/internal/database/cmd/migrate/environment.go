package migrate

import (
	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
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
			logs.Info.Println("[Ignore exist flag]: Existing relations won't stop the operations neither throw errors..")
		}

		path, _ := cmd.Flags().GetStringArray("environmentFiles")
		if err := environment.Parse(path...); err != nil {
			logs.Error.Fatal("[Invalid environment file]: ", err.ParseToString())
		}

		// PARSE POSTGRES CONNECTION REQUIRED FOR MIGRATION FROM ENVIRONMENT
		pgConn, res := connections.ScanEnvironmentForConnection(connections.Postgres)
		if res != nil {
			logs.Error.Fatal(res.ParseToString())
		}

		db, res := connections.Connect(connections.ParseConnectionString(pgConn))
		if res != nil {
			logs.Error.Fatal(res.ParseToString())
		}
		defer db.Close()

		utils := Utils{migrations: migrations.New(db)}

		res = utils.MigrateEnvironment(connections.ConnectionEnvironment(args[0]), ignoreExisting)
		if res != nil {
			logs.Error.Fatal(res.ParseToString())
		}
		logs.Info.Println("[Environment migrated successfully]")
	},
}

func init() {
	EnvironmentCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "Defines the path to the environment files containing the necessary environment variables if not already supplied in the environment")
	EnvironmentCmd.Flags().BoolP("ignore-existing", "i", false, "Doesn't throw errors and proceed if the database relation already exists.")
	// Not implemented Environment.Flags().BoolP("recreate-existing", "r", false, "Drop and recreate the relation if already exists.")
}

func (u *Utils) MigrateEnvironment(environment connections.ConnectionEnvironment, ignoreExisting bool) *common.Response {
	// PARSE DESIRED ENVIRONMENT SETTINGS FROM ENVIRONMENT
	conn, res := connections.ScanEnvironmentForConnection(connections.ConnectionEnvironment(environment))
	if res != nil {
		return res
	}

	// CREATE DESIRED ENVIRONMENT ON POSTGRES
	res = u.CreateEnvironmentDatabases(
		[]*migrations.RequestCreateDatabase{
			{User: conn.User(), Database: conn.Database()},
		},
		&migrations.Configuration{
			IgnoreExisting: ignoreExisting,
		},
	)
	if res != nil {
		return res
	}

	if u.CreateEnvironmentUsers(conn) {
		res := u.migrations.DropDatabases(conn.Database())
		if res != nil {
			return res
		}
	}

	res = u.GrantAllPrivilegesOnDatabase(models.NewDatabaseUser(conn.User(), conn.Password(), conn.Database()))
	if res != nil {
		res = u.migrations.DropDatabases(conn.Database())
		if res != nil {
			return res
		}

		res = u.migrations.DropUsers(conn.User())
		if res != nil {
			return res
		}
		return res
	}
	return nil
}
