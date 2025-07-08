package migrations_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

var flagEnvironmentFile = "../../../secrets/postgres.env"

func TestMigrations(t *testing.T) {
	var db *sql.DB
	t.Run("SETUP", func(t *testing.T) {
		if err := environment.Parse(flagEnvironmentFile); err != nil {
			t.Fatal(err.Status, err.Description)
		}

		environment := connections.ConnectionEnvironment(connections.Postgres)

		var res *common.Response
		db, res = connections.ConnectWithEnvironment(environment)
		if res != nil {
			t.Fatal(res.ParseToString())
		}
	})

	migrations := migrations.New(db)
	defer db.Close()

	// PUT YOUR TYPES FOR TESTS HERE
	types := []*models.TypeInfo{
		models.UserRolesType,
	}

	// PUT YOUR TABLE FOR TESTS HERE
	tables := []models.TableMethods{
		models.TablesInfo.Institutions,
		models.TablesInfo.Users,
		models.TablesInfo.Classes,
		models.TablesInfo.Notifications,
		models.TablesInfo.UsersClasses,
		models.TablesInfo.UsersNotifications,
	}

	t.Run("PING", func(t *testing.T) {
		if err := db.Ping(); err != nil {
			t.Log("STATUS: " + err.Error())
			t.Fatal("DESCRIPTION: Unable to ping. " + err.Error())
		}
	})

	var res *common.Response
	for _, table := range tables {
		t.Run(fmt.Sprintf("CREATE TABLE %s", table.Name()), func(t *testing.T) {
			res = migrations.CreateTables(nil, table).Response

			if res != nil {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res != nil {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, table := range tables {
		t.Run(fmt.Sprintf("DROP TABLE %s", table.Name()), func(t *testing.T) {
			res = migrations.DropTables(nil, table).Response
			if res != nil {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res != nil {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, typ := range types {
		t.Run(fmt.Sprintf("CREATE TYPE %s", typ.Name), func(t *testing.T) {
			res = migrations.CreateType(nil, typ).Response
			if res != nil {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res != nil {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, typ := range types {
		t.Run(fmt.Sprintf("DROP TYPE %s", typ.Name), func(t *testing.T) {
			res = migrations.DropType(nil, typ).Response
			if res != nil {
				t.Log("STATUS: ", res.Status)
				t.Fatal(res.Description)
				return
			}
		})
		if res != nil {
			t.Fatal("Tests failed.")
			return
		}
	}
}
