package operations_test

import (
	"database/sql"
	"testing"

	"github.com/Noeeekr/singullar/server/pkg/pim"
	"github.com/Noeeekr/singullar/server/pkg/pim/migrations"
	"github.com/Noeeekr/singullar/server/pkg/pim/models"
	"github.com/Noeeekr/singullar/server/pkg/pim/operations"

	"github.com/Noeeekr/singullar/server/pkg/pim/transactions"
)

var args = []string{"./postgres.env"}
var tables = []models.TableMethods{models.TablesInfo.Users, models.TablesInfo.Institutions}

type Table struct {
	Name      models.TableName
	QueryInfo *transactions.TransactionRequest
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

	migrations *migrations.MigrationsManager
	operations *operations.Operations
}

func (m *QueryTestUtil) CreateTables(tables []models.TableMethods) (res *transactions.Response) {
	res = m.migrations.CreateTables(nil, tables...)
	if res != nil {
		m.test.Log(res.Status, "|", res.Description)
	}
	return res
}
func (m *QueryTestUtil) DropTables(tables []models.TableMethods) (res *transactions.Response) {
	res = m.migrations.DropTables(tables...)
	if res != nil {
		m.test.Log(res.Status, "|", res.Description)
	}
	return res
}

// Start database connection, create necessary tables.
func (m *QueryTestUtil) Prepare(tables []models.TableMethods) *sql.DB {
	connString := pim.GetConnectionStringFromFiles(args...)
	db, err := pim.Connect(connString)
	if err != nil {
		m.test.Fatal(err.Error())
		return nil
	}

	m.db = db
	m.migrations = migrations.New(m.db)
	m.operations = operations.New(m.db)

	m.test.Run("CREATE TABLES", func(t *testing.T) {
		if res := m.CreateTables(tables); res != nil {
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
	var institution_id int = -1
	var institution *models.Institutions

	var user_email string = "noeeekr@gmail.com"

	var res *transactions.Response
	if !test.Run("INSERT INSTITUTION", func(t *testing.T) {
		institution_id, res = utils.operations.InsertInstitution(institution_name)
		if res != nil {
			t.FailNow()
		}
	}) {
		test.Fatal(res.Description)
	}
	if !test.Run("SELECT INSTITUTION BY ID", func(t *testing.T) {
		institution, res = utils.operations.SelectInstitutionById(institution_id)
		if res != nil {
			t.Fatal(res.Description)
		}
		if institution == nil {
			t.Fatal("Nil institution")
		}
	}) {
		test.FailNow()
	}

	if !test.Run("SELECT INSTITUTION BY NAME", func(t *testing.T) {
		institution = nil
		institution, res = utils.operations.SelectInstitutionByName(institution_name)
		if res != nil {
			t.Fatal(res.Description)
		}
		if institution == nil {
			t.Fatal("Nil institution")
		}
	}) {
		test.FailNow()
	}

	if !test.Run("INSERT USER", func(t *testing.T) {
		user_email, res = utils.operations.InsertUser("noeeekr", "noeeekr@gmail.com", "123123123", institution_id, models.Admin)
		if res != nil {
			t.Fatal(res.Description)
		}
		if user_email != user_email {
			t.Fatal("Nil institution")
		}
	}) {
		test.FailNow()
	}
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
	if res != nil {
		t.Fatal(res.Description)
	}

	res = utilities.CreateTables(tables)
	if res != nil {
		if res := utilities.DropTables(tables); res != nil {
			t.Fatal(res.Description)
		}
		t.Fatal(res.Description)
	}

	// INSERT
	institution_name := "claretiano"
	institution_id, res := utilities.InsertInstitution(institution_name)
	if res != nil {
		t.Fatal(res.Description)
	}

	// SELECT
	institution, res := utilities.SelectInstitution(institution_id)
	if res != nil {
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
	if res != nil {
		// If user already exists delete it and try again
		if res.Status != transactions.StatusAlreadyExists {
			t.Fatal(res.Description)
		}
		res = utilities.DeleteUser(user_email)
		if res != nil {
			t.Fatal(res.Description)
		}
		_, res = utilities.InsertUser("Andrew", user_email, "123321", institution.Id, models.Admin)
		if res != nil {
			t.Fatal(res.Description)
		}
	}
	if res != nil {
		t.Fatal(res.Description)
	}

	// SELECT
	user, res := utilities.SelectUser(user_email)
	if res != nil {
		t.Fatal(res.Description)
	}

	// COMPARE INSERT SELECT RETURNS
	equal = utilities.CompareReturns(user_email, user.Email)
	if !equal {
		t.Fatal("NAMES DOESN'T MATCH")
	}

	// DELETE USER
	res = utilities.DeleteUser(user.Email)
	if res != nil {
		t.Fatal(res.Description)
	}

	// DELETE INSTITUTION
	res = utilities.DeleteInstitution(institution.Id)
	if res != nil {
		t.Fatal(res.Description)
	}

	res = utilities.DropTables(tables)
	if res != nil {
		t.Fatal(res.Description)
	}
}


func (m *QueryTestUtil) SelectInstitution(id int) (i *models.Institutions, res *transactions.Response) {
	testname := fmt.Sprintf("SELECT FROM %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		i, res = m.operations.Select(id)
		if res != nil {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) DeleteInstitution(id int) (res *transactions.Response) {
	m.test.Run("DELETE INSTITUTION", func(t *testing.T) {
		res = m.operations.Delete(id)
		if res != nil {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) InsertUser(name, email, password string, institution_id int, role models.UserRole) (Email string, res *transactions.Response) {
	testname := fmt.Sprintf("INSERT INTO %s", models.UsersTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		Email, res = m.operations.Insert(name, email, password, institution_id, role)
		if res != nil {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) SelectUser(email string) (i *models.Users, res *transactions.Response) {
	testname := fmt.Sprintf("SELECT FROM %s", models.InstitutionsTable.Name())
	m.test.Run(testname, func(t *testing.T) {
		i, res = m.operations.Select(email)
		if res != nil {
			t.Fail()
		}
	})
	return
}
func (m *QueryTestUtil) DeleteUser(email string) (res *transactions.Response) {
	m.test.Run("DELETE USER", func(t *testing.T) {
		res = m.operations.Delete(email)
		if res != nil {
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
