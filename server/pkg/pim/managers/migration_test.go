package managers_test

import (
	"database/sql"
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
}
func TestMigrations(t *testing.T) {
	manager := managers.NewMigrationManager(db)
	defer manager.Close()

	t.Run("PING", func(t *testing.T) {
		if err := manager.Ping(); err != nil {
			t.Log("STATUS: " + err.Error())
			t.Fatal("DESCRIPTION: Unable to ping. " + err.Error())
		}
	})

	t.Run("DROP TABLE USERS", func(t *testing.T) {
		query := manager.DropTable(models.UsersTableName)
		if query.Status != managers.StatusSuccess {
			t.Log("DESCRIPTION: ", query.Description)
			t.Log("STATUS: ", query.Status)
			t.FailNow()
		}
	})

	t.Run("CREATE TABLE USERS", func(t *testing.T) {
		query := manager.CreateTable(models.UsersTableName, false)

		if query.Status != managers.StatusSuccess {
			t.Log("DESCRIPTION: ", query.Description)
			t.Log("STATUS: ", query.Status)
			t.FailNow()
		}
	})

	t.Run("CREATE ENUM ROLES", func(t *testing.T) {
		res := manager.CreateType(models.UserRoleName)
		if res.Status != managers.StatusSuccess {
			t.Fatal(res.Description)
		}
	})

	t.Run("DELETE ENUM ROLES", func(t *testing.T) {
		res := manager.DropType(models.UserRoleName)
		if res.Status != managers.StatusSuccess {
			t.Fatal(res.Description)
		}
	})
}
