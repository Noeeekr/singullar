package query_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Noeeekr/singullar/server/pkg/pim/managers"
	"github.com/Noeeekr/singullar/server/pkg/pim/models"
	"github.com/Noeeekr/singullar/server/pkg/pim/query"
	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

var args = []string{"./../postgres.env"}
var tables = []*models.TableInfo{models.InstitutionsTable, models.UsersTable}

type Table struct {
	Name      models.TableName
	QueryInfo *transaction.QueryInfo
}

/*
    insertTest := []QueryTest{
		Name: InstitutionTable.Name()
		Query: InstitutionTable.Queries.Create,
		Args: "Claretiano",
		ReturnFunc: func (rows *sql.Rows) {
			...
		}
	}
*/

type QueryTestUtil struct {
	db *sql.DB

	test *testing.T

	migrations *managers.MigrationsManager
	query      *query.QueryManager
}

func (m *QueryTestUtil) Insert(table *Table) *QueryTestUtil {
	testname := fmt.Sprintf("INSERT INTO %s", table.Name)

	m.test.Run(testname, func(test *testing.T) {
		res := m.query.Insert(table.QueryInfo)
		if res.Status != transaction.StatusSuccess {
			m.test.Fatal(res.Status.ToString(), "|", res.Description)
		}
	})

	return m
}
func (m *QueryTestUtil) Select(table *Table) *QueryTestUtil {
	testname := fmt.Sprintf("SELECT FROM %s", table.Name)

	m.test.Run(testname, func(test *testing.T) {
		res := m.query.Select(table.QueryInfo)
		if res.Status != transaction.StatusSuccess {
			m.test.Fatal(res.Status.ToString(), "|", res.Description)
		}
	})

	return m
}
func (m *QueryTestUtil) Delete(table *Table) *QueryTestUtil {
	testname := fmt.Sprintf("DELETE FROM %s", table.Name)

	m.test.Run(testname, func(test *testing.T) {
		res := m.query.Delete(table.QueryInfo)
		if res.Status != transaction.StatusSuccess {
			m.test.Fatal(res.Status.ToString(), "|", res.Description)
		}
	})

	return m
}
func (m *QueryTestUtil) CreateTables(tables []*models.TableInfo) (res *transaction.Response) {
	res = m.migrations.CreateTables(nil, tables...)
	if res.Status != transaction.StatusSuccess {
		m.test.Log(res.Status, "|", res.Description)
	}
	return res
}
func (m *QueryTestUtil) DropTables(tables []*models.TableInfo) (res *transaction.Response) {
	res = m.migrations.DropTables(tables...)
	if res.Status != transaction.StatusSuccess {
		m.test.Log(res.Status, "|", res.Description)
	}
	return res
}

// Start database connection, create necessary tables.
func (m *QueryTestUtil) Prepare(tables []*models.TableInfo) *sql.DB {
	connString := managers.GetConnectionStringFromFiles(args...)
	db, err := managers.Connect(connString)
	if err != nil {
		m.test.Fatal(err.Error())
		return nil
	}

	m.db = db
	m.migrations = managers.NewMigrationManager(m.db)
	m.query = query.NewQueryManager(m.db)

	m.test.Run("CREATE TABLES", func(t *testing.T) {
		if res := m.CreateTables(tables); res.Status != transaction.StatusSuccess {
			m.test.Fatal(res.Status.ToString(), res.Description)
		}
	})

	return m.db
}

func TestQueries(test *testing.T) {
	utils := QueryTestUtil{
		test: test,
	}
	db := utils.Prepare(tables)
	if db == nil {
		test.Fatal("Database not prepared")
	}
	defer db.Close()
	defer utils.DropTables(tables)

	var institution_name string = "claretiano"
	var institutions_ids []int = []int{}
	var institutions []*models.Institutions

	var user_email string = "noeeekr@gmail.com"
	var users_ids []int = []int{}
	var users []*models.Users

	utils.Insert(&Table{
		QueryInfo: models.InstitutionsTable.Queries.InsertOne.
			WithArgs(time.Now(), time.Now(), institution_name).
			WithScanFunc(query.ScanInstitutionsIds(&institutions_ids)),
		Name: models.InstitutionsTable.Name(),
	})
	if len(institutions_ids) == 0 {
		test.Fatal("Failed Scan")
	}
	utils.Select(&Table{
		QueryInfo: models.InstitutionsTable.Queries.SelectOne.
			WithArgs(institutions_ids[0]).
			WithScanFunc(query.ScanInstitutions(&institutions)),
		Name: models.InstitutionsTable.Name(),
	})
	if len(institutions) == 0 {
		test.Fatal("Failed Scan")
	}
	utils.Insert(&Table{
		Name: models.UsersTable.Name(),
		QueryInfo: models.UsersTable.Queries.InsertOne.
			WithArgs().
			WithScanFunc(query.ScanUsers(use)),
	})
}

/*
func TestQueries(t *testing.T) {
	tables := []*models.TableInfo{models.InstitutionsTable, models.UsersTable}

	utilities := QueryTestUtil{
		test: t,
	}
	db, err := utilities.Prepare(tables)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer db.Close()

	res := utilities.DropTables(tables)
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}

	res = utilities.CreateTables(tables)
	if res.Status != transaction.StatusSuccess {
		if res := utilities.DropTables(tables); res.Status != transaction.StatusSuccess {
			t.Fatal(res.Description)
		}
		t.Fatal(res.Description)
	}

	// INSERT
	institution_name := "claretiano"
	institution_id, res := utilities.InsertInstitution(institution_name)
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}

	// SELECT
	institution, res := utilities.SelectInstitution(institution_id)
	if res.Status != transaction.StatusSuccess {
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
	if res.Status != transaction.StatusSuccess {
		// If user already exists delete it and try again
		if res.Status != transaction.StatusAlreadyExists {
			t.Fatal(res.Description)
		}
		res = utilities.DeleteUser(user_email)
		if res.Status != transaction.StatusSuccess {
			t.Fatal(res.Description)
		}
		_, res = utilities.InsertUser("Andrew", user_email, "123321", institution.Id, models.Admin)
		if res.Status != transaction.StatusSuccess {
			t.Fatal(res.Description)
		}
	}
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}

	// SELECT
	user, res := utilities.SelectUser(user_email)
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}

	// COMPARE INSERT SELECT RETURNS
	equal = utilities.CompareReturns(user_email, user.Email)
	if !equal {
		t.Fatal("NAMES DOESN'T MATCH")
	}

	// DELETE USER
	res = utilities.DeleteUser(user.Email)
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}

	// DELETE INSTITUTION
	res = utilities.DeleteInstitution(institution.Id)
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}

	res = utilities.DropTables(tables)
	if res.Status != transaction.StatusSuccess {
		t.Fatal(res.Description)
	}
}


func (m *QueryTestUtil) SelectInstitution(id int) (i *models.Institutions, res *transaction.Response) {
	testname := fmt.Sprintf("SELECT FROM %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		i, res = m.query.Select(id)
		if res.Status != transaction.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) DeleteInstitution(id int) (res *transaction.Response) {
	m.test.Run("DELETE INSTITUTION", func(t *testing.T) {
		res = m.query.Delete(id)
		if res.Status != transaction.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) InsertUser(name, email, password string, institution_id int, role models.UserRole) (Email string, res *transaction.Response) {
	testname := fmt.Sprintf("INSERT INTO %s", models.UsersTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		Email, res = m.query.Insert(name, email, password, institution_id, role)
		if res.Status != transaction.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) SelectUser(email string) (i *models.Users, res *transaction.Response) {
	testname := fmt.Sprintf("SELECT FROM %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		i, res = m.query.Select(email)
		if res.Status != transaction.StatusSuccess {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) DeleteUser(email string) (res *transaction.Response) {
	m.test.Run("DELETE USER", func(t *testing.T) {
		res = m.query.Delete(email)
		if res.Status != transaction.StatusSuccess {
			t.Fail()
		}
	})
	return
}

func (m *QueryTestUtil) CompareReturns(want, have string) (equal bool) {
	m.test.Run("COMPARE RETURNS", func(t *testing.T) {
		if want != have {
			t.FailNow()
		}
		equal = true
	})
	return
}
*/
