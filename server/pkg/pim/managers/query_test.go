package managers_test

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/Noeeekr/singullar/server/pkg/pim/managers"
	"github.com/Noeeekr/singullar/server/pkg/pim/models"
)

var queries_args = []string{"./../postgres.env"}

type TestQueriesUtilities struct {
	db         *sql.DB
	test       *testing.T
	migrations *managers.MigrationsManager
	queries    *managers.QueryManager
}

func (m *TestQueriesUtilities) CreateTables(tables []*models.TableInfo) (res *managers.Response) {
	for _, table := range tables {
		testname := fmt.Sprintf("CREATE TABLE %s", table.Name())
		m.test.Run(testname, func(t *testing.T) {
			res = m.migrations.CreateTable(table)
			if res.Status != managers.StatusSuccess {
				t.Fail()
			}
		})
		if res.Status != managers.StatusSuccess {
			return res
		}
	}
	return &managers.Response{
		Status:      managers.StatusSuccess,
		Description: "Tables created.",
	}
}

func (m *TestQueriesUtilities) DropTables(tables []*models.TableInfo) (res *managers.Response) {
	for _, table := range tables {
		testname := fmt.Sprintf("DROP TABLE %s", table.Name())
		m.test.Run(testname, func(t *testing.T) {
			res = m.migrations.DropTable(table)
			if res.Status != managers.StatusSuccess {
				t.Fail()
			}
		})
		if res.Status != managers.StatusSuccess {
			return res
		}
	}
	return &managers.Response{
		Status:      managers.StatusSuccess,
		Description: "Tables created.",
	}
}

// Start database connection, create necessary tables.
func (m *TestQueriesUtilities) Prepare(tables []*models.TableInfo) (*sql.DB, error) {
	connString := managers.GetConnectionStringFromFiles(queries_args...)
	db, err := managers.Connect(connString)
	if err != nil {
		return nil, err
	}
	m.db = db
	m.migrations = managers.NewMigrationManager(m.db)
	m.queries = managers.NewQueryManager(m.db)

	if res := m.CreateTables(tables); res.Status != managers.StatusSuccess {
		if res := m.DropTables(tables); res.Status != managers.StatusSuccess {
			return nil, errors.New(res.Description)
		}
		return nil, errors.New(res.Description)
	}

	return m.db, nil
}

func (m *TestQueriesUtilities) InsertInstitution(name string) (id int, res *managers.Response) {
	testname := fmt.Sprintf("INSERT INTO %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		id, res = m.queries.InsertInstitution(name)
	})
	return
}

func (m *TestQueriesUtilities) SelectInstitution(id int) (i *models.Institutions, res *managers.Response) {
	testname := fmt.Sprintf("SELECT FROM %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		i, res = m.queries.SelectInstitution(id)
		if res.Status != managers.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *TestQueriesUtilities) DeleteInstitution(id int) (res *managers.Response) {
	m.test.Run("DELETE INSTITUTION", func(t *testing.T) {
		res = m.queries.DeleteInstitution(id)
		if res.Status != managers.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *TestQueriesUtilities) InsertUser(name, email, password string, institution_id int, role models.UserRole) (Email string, res *managers.Response) {
	testname := fmt.Sprintf("INSERT INTO %s", models.UsersTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		Email, res = m.queries.InsertUser(name, email, password, institution_id, role)
		if res.Status != managers.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *TestQueriesUtilities) SelectUser(email string) (i *models.Users, res *managers.Response) {
	testname := fmt.Sprintf("SELECT FROM %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		i, res = m.queries.SelectUser(email)
		if res.Status != managers.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *TestQueriesUtilities) DeleteUser(email string) (res *managers.Response) {
	m.test.Run("DELETE USER", func(t *testing.T) {
		res = m.queries.DeleteUser(email)
		if res.Status != managers.StatusSuccess {
			t.Fail()
		}
	})
	return
}

func (m *TestQueriesUtilities) CompareReturns(want, have string) (equal bool) {
	m.test.Run("COMPARE RETURNS", func(t *testing.T) {
		if want != have {
			t.FailNow()
		}
		equal = true
	})
	return
}

func TestQueries(t *testing.T) {
	tables := []*models.TableInfo{models.InstitutionsTable, models.UsersTable}

	utilities := TestQueriesUtilities{
		test: t,
	}
	db, err := utilities.Prepare(tables)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer db.Close()

	res := utilities.DropTables(tables)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	res = utilities.CreateTables(tables)
	if res.Status != managers.StatusSuccess {
		if res := utilities.DropTables(tables); res.Status != managers.StatusSuccess {
			t.Fatal(res.Description)
		}
		t.Fatal(res.Description)
	}

	// INSERT
	institution_name := "claretiano"
	institution_id, res := utilities.InsertInstitution(institution_name)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	// SELECT
	institution, res := utilities.SelectInstitution(institution_id)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	// COMPARE INSERT SELECT RETURNS
	equal := utilities.CompareReturns(institution_name, institution.Name)
	if !equal {
		t.Fatal("NAMES DOESN'T MATCH")
	}

	// INSERT
	user_email := "noeeekr@gmail.com"
	_, res = utilities.InsertUser("Andrew", user_email, "123321", institution.Id, models.Admin)
	if res.Status != managers.StatusSuccess {
		// If user already exists delete it and try again
		if res.Status != managers.StatusAlreadyExists {
			t.Fatal(res.Description)
		}
		res = utilities.DeleteUser(user_email)
		if res.Status != managers.StatusSuccess {
			t.Fatal(res.Description)
		}
		_, res = utilities.InsertUser("Andrew", user_email, "123321", institution.Id, models.Admin)
		if res.Status != managers.StatusSuccess {
			t.Fatal(res.Description)
		}
	}
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	// SELECT
	user, res := utilities.SelectUser(user_email)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	// COMPARE INSERT SELECT RETURNS
	equal = utilities.CompareReturns(user_email, user.Email)
	if !equal {
		t.Fatal("NAMES DOESN'T MATCH")
	}

	// DELETE USER
	res = utilities.DeleteUser(user.Email)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	// DELETE INSTITUTION
	res = utilities.DeleteInstitution(institution.Id)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}

	res = utilities.DropTables(tables)
	if res.Status != managers.StatusSuccess {
		t.Fatal(res.Description)
	}
}
