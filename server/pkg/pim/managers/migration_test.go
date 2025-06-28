package managers_test

import (
	"fmt"
	"testing"

	"github.com/Noeeekr/singullar/server/pkg/pim/managers"
	"github.com/Noeeekr/singullar/server/pkg/pim/models"
	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

var migrations_args = []string{"./../postgres.env"}

func TestMigrations(t *testing.T) {
	connString := managers.GetConnectionStringFromFiles(migrations_args...)
	db, err := managers.Connect(connString)
	if err != nil {
		panic("Status: Error happened. " + err.Error())
	}
	migrations := managers.NewMigrationManager(db)
	defer db.Close()

	// PUT YOUR TYPES FOR TESTS HERE
	types := []*models.TypeInfo{
		models.UserRolesType,
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
		if err := db.Ping(); err != nil {
			t.Log("STATUS: " + err.Error())
			t.Fatal("DESCRIPTION: Unable to ping. " + err.Error())
		}
	})

	var res *transaction.Response
	for _, table := range tables {
		t.Run(fmt.Sprintf("CREATE TABLE %s", table.Name()), func(t *testing.T) {
			res = migrations.CreateTables(nil, table)
			if res.Status != transaction.StatusSuccess {
				t.Log("STATUS: ", res.Status.ToString())
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != transaction.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, table := range tables {
		t.Run(fmt.Sprintf("DROP TABLE %s", table.Name()), func(t *testing.T) {
			res = migrations.DropTables(table)
			if res.Status != transaction.StatusSuccess {
				t.Log("STATUS: ", res.Status.ToString())
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != transaction.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, typ := range types {
		t.Run(fmt.Sprintf("CREATE TYPE %s", typ.Name), func(t *testing.T) {
			res = migrations.CreateType(typ)
			if res.Status != transaction.StatusSuccess {
				t.Log("STATUS: ", res.Status.ToString())
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != transaction.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}

	for _, typ := range types {
		t.Run(fmt.Sprintf("DROP TYPE %s", typ.Name), func(t *testing.T) {
			res = migrations.DropType(typ)
			if res.Status != transaction.StatusSuccess {
				t.Log("STATUS: ", res.Status.ToString())
				t.Fatal(res.Description)
				return
			}
		})
		if res.Status != transaction.StatusSuccess {
			t.Fatal("Tests failed.")
			return
		}
	}
}
