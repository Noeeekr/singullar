package managers_test

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Noeeekr/singullar/server/pkg/common"
	"github.com/Noeeekr/singullar/server/pkg/common/configs"
	"github.com/Noeeekr/singullar/server/pkg/pim/managers"
	"github.com/Noeeekr/singullar/server/pkg/pim/models"
)

var args = []string{"./../postgres.env"}
var env = &managers.PostgrestManagerEnvironment{}
var db *sql.DB
var manager *managers.MigrationsManager

// Starts database connection
func init() {
	for i, arg := range args {
		os.Args[i+1] = arg
	}

	root_path := common.GetExecutableDir()
	if root_path == "" {
		panic("Failed to get executable path")
	}

	relative_path := common.GetFirstArgument()
	if relative_path == "" {
		panic("Please provide the path to the environment file.")
	}

	env_path := filepath.Join(root_path, relative_path)
	if err := configs.Parse(env_path); err != nil {
		panic("Status: Failed to parse environment file: " + err.Error())
	}

	if err := configs.Scan(env); err != nil {
		panic("Status: Failed to scan environment file into object: " + err.Error())
	}

	connString := managers.ParseConnectionString(
		env.POSTGRES_CONTAINER_NAME,
		env.POSTGRES_USER,
		env.POSTGRES_USER_PASSWORD,
	)

	var err error

	db, err = managers.Connect(connString)
	if err != nil {
		panic("Status: Error happened. " + err.Error())
	}

	manager = managers.NewMigrationManager(db)
}
func TestMigrations(t *testing.T) {
	defer manager.Close()

	// PUT YOUR TYPES FOR TESTS HERE
	types := []*models.TypeInfo{
		models.RoleType,
	}

	// PUT YOUR TABLE FOR TESTS HERE
	tables := []*models.TableInfo{
		models.InstitutionsTable,
		models.UsersTable,
		models.ClassesTable,
		models.NotificationsTable,
		models.UsersClassesTable,
		models.UsersNotificationsTable,
	}

	t.Run("PING", func(t *testing.T) {
		if err := manager.Ping(); err != nil {
			t.Log("STATUS: " + err.Error())
			t.Fatal("DESCRIPTION: Unable to ping. " + err.Error())
		}
	})

	var res *managers.Response
	for _, table := range tables {
		t.Run(fmt.Sprintf("CREATE TABLE %s", table.TableName()), func(t *testing.T) {
			res = manager.CreateTable(table)
			if res.Status != managers.StatusSuccess {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != managers.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, table := range tables {
		t.Run(fmt.Sprintf("DROP TABLE %s", table.TableName()), func(t *testing.T) {
			res = manager.DropTable(table)
			if res.Status != managers.StatusSuccess {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != managers.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, typ := range types {
		t.Run(fmt.Sprintf("CREATE TYPE %s", typ.Name()), func(t *testing.T) {
			res = manager.CreateType(typ)
			if res.Status != managers.StatusSuccess {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != managers.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, typ := range types {
		t.Run(fmt.Sprintf("DROP TYPE %s", typ.Name()), func(t *testing.T) {
			res = manager.DropType(typ)
			if res.Status != managers.StatusSuccess {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != managers.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}
}
