package migrations_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Noeeekr/singullar/server/common/configs"
	"github.com/Noeeekr/singullar/server/internal/database"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

var migrations_args = []string{"./postgres.env"}

func TestMigrations(t *testing.T) {
	var db *sql.DB
	t.Run("SETUP", func(t *testing.T) {
		if err := configs.Parse(migrations_args...); err != nil {
			t.Fatal(err.Status, err.Description)
		}

		var env database.PostgresEnvironment
		if err := configs.Scan(&env); err != nil {
			t.Fatal(err.ParseToError())
		}

		var err error
		db, err = database.Connect(env.POSTGRES_CONNECTION_STRING)
		if err != nil {
			t.Fatal(err.Error())
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

	var res *transactions.Response
	for _, table := range tables {
		t.Run(fmt.Sprintf("CREATE TABLE %s", table.Name()), func(t *testing.T) {
			res = migrations.CreateTables(nil, table)

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
			res = migrations.DropTables(table)
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
			res = migrations.CreateType(typ)
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
			res = migrations.DropType(typ)
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
