package operations_test

import (
	"crypto/rand"
	"database/sql"
	"math/big"
	"testing"

	"github.com/Noeeekr/singullar/server/internal/database"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

var args = []string{"./postgres.env"}
var tables = []models.TableMethods{models.TablesInfo.Users, models.TablesInfo.Institutions}

type InstitutionTestCase struct {
	Name string
	Id   int

	TargetInstance *models.Institutions
}

type UserTestCase struct {
	Id            int
	Name          string
	Email         string
	Password      string
	InstitutionId int
	Role          models.UserRole

	TargetInstance *models.Users
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
func (m *QueryTestUtil) CreateInstitutionTestCases(amount int) []*InstitutionTestCase {
	cases := []*InstitutionTestCase{}
	for range amount {
		cases = append(
			cases,
			&InstitutionTestCase{
				Name:           m.GenerateRandomStrings(10),
				Id:             -1,
				TargetInstance: nil,
			},
		)
	}
	return cases
}
func (m *QueryTestUtil) CreateUserTestCases(institutions []*InstitutionTestCase) []*UserTestCase {
	cases := []*UserTestCase{}
	for _, institution := range institutions {
		cases = append(
			cases,
			&UserTestCase{
				Id:             -1,
				Name:           m.GenerateRandomStrings(10),
				Email:          m.GenerateRandomStrings(6) + "@" + m.GenerateRandomStrings(4) + ".com",
				Password:       m.GenerateRandomStrings(15),
				Role:           models.Admin,
				InstitutionId:  institution.Id,
				TargetInstance: nil,
			},
		)
	}

	return cases
}
func (m *QueryTestUtil) GenerateRandomStrings(size int) string {
	var s string = ""
	for i := len(s); i < size; {
		letter, err := rand.Int(rand.Reader, big.NewInt(91))
		if err != nil {
			panic(err)
		}
		if letter.Int64() > 64 {
			s += string(rune(letter.Int64()))
			i++
		}
	}
	return s
}

// Start database connection, create necessary tables.
func (m *QueryTestUtil) Prepare(tables []models.TableMethods) *sql.DB {
	connString := database.GetConnectionStringFromFiles(args...)
	db, err := database.Connect(connString)
	if err != nil {
		m.test.Fatal(err.Error())
		return nil
	}

	m.db = db
	m.migrations = migrations.New(m.db)
	m.operations = operations.New(m.db)

	m.Run("PING", func(t *testing.T) {
		err := m.db.Ping()
		if err != nil {
			t.Fatal("Unable to ping database. " + err.Error())
		}
	})

	m.Run("CREATE TABLES", func(t *testing.T) {
		if res := m.CreateTables(tables); res != nil {
			t.Fatal(res.Status.ToString(), res.Description)
		}
	})

	return m.db
}

func (m *QueryTestUtil) Run(title string, fun func(t *testing.T)) *QueryTestUtil {
	if !m.test.Run(title, fun) {
		m.test.FailNow()
	}
	return m
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

	institutions := utils.CreateInstitutionTestCases(10)

	var res *transactions.Response

	utils.
		Run("INSERT INSTITUTIONS", func(t *testing.T) {
			for index, institution := range institutions {
				institutions[index].Id, res = utils.operations.InsertInstitution(institution.Name)
				if res != nil {
					t.FailNow()
				}
			}
		}).
		Run("SELECT INSTITUTIONS BY ID", func(t *testing.T) {
			for index, inst := range institutions {
				institutions[index].TargetInstance, res = utils.operations.SelectInstitutionById(inst.Id)
				if res != nil {
					t.Fatal(res.Description)
				}
				if institutions[index].TargetInstance == nil {
					t.Fatal("Nil institution")
				}
			}
		}).
		Run("SELECT INSTITUTIONS BY NAME", func(t *testing.T) {
			for index, institution := range institutions {
				institutions[index].TargetInstance = nil
				institutions[index].TargetInstance, res = utils.operations.SelectInstitutionByName(institution.Name)
				if res != nil {
					t.Fatal(res.Description)
				}
				if institutions[index].TargetInstance == nil {
					t.Fatal("Nil institution")
				}
			}
		})

	users := utils.CreateUserTestCases(institutions)

	utils.Run("INSERT USERS", func(t *testing.T) {
		for _, user := range users {
			email, res := utils.operations.InsertUser(user.Name, user.Email, user.Password, user.InstitutionId, user.Role)
			if res != nil {
				t.Fatal(res.Description)
			}
			if user.Email != email {
				t.Fatal("Emails don't match")
			}
		}
	}).Run("SELECT USERS BY EMAIL", func(t *testing.T) {
		for _, userTestCase := range users {
			(*userTestCase).TargetInstance = nil
			(*userTestCase).TargetInstance, res = utils.operations.SelectUserByEmail(userTestCase.Email)
			if res != nil {
				t.Fatal(res.Description)
			}
			if userTestCase.Email != (*userTestCase).TargetInstance.Email {
				t.Fatal("Nil user")
			}
			(*userTestCase).Id = (*userTestCase).TargetInstance.Id
		}
	}).Run("SELECT USERS BY ID", func(t *testing.T) {
		for _, userTestCase := range users {
			(*userTestCase).TargetInstance = nil
			(*userTestCase).TargetInstance, res = utils.operations.SelectUserById(userTestCase.Id)
			if res != nil {
				t.Fatal(res.Description)
			}
			if userTestCase.Email != (*userTestCase).TargetInstance.Email {
				t.Fatal("Nil user")
			}
		}
	})

	utils.Run("DELETE USERS BY EMAIL AND ID", func(t *testing.T) {
		if len(users) < 2 {
			t.Fatal("Insufficient test cases for users")
		}
		for i, userTestCase := range users {
			if i > len(users)/2 {
				res = utils.operations.DeleteUserByEmail((*userTestCase).TargetInstance.Email)
			} else {
				res = utils.operations.DeleteUserById((*userTestCase).TargetInstance.Id)
			}
			if res != nil {
				test.Fatal(res.Status.ToString(), res.Description)
			}
		}
	})
	utils.Run("DELETE INSTITUTIONS BY EMAIL AND ID", func(t *testing.T) {
		if len(institutions) < 2 {
			t.Fatal("Insufficient test cases for institutions")
		}
		for i, institutionsTestCase := range institutions {
			if i > len(users)/2 {
				res = utils.operations.DeleteInstitutionById((*institutionsTestCase).TargetInstance.Id)
			} else {
				res = utils.operations.DeleteInstitutionByName((*institutionsTestCase).TargetInstance.Name)
			}
			if res != nil {
				test.Fatal(res.Status.ToString(), res.Description)
			}
		}
	})
}

/*
	INSERT USER                         done
	INSERT INSTITUTION                  done
	INSERT CLASS
		INSERT USERCLASS
	INSERT NOTIFICATION
		INSERT USERNOTIFICATION

	SELECT USER                         done
	SELECT INSTITUTION                  done
	SELECT CLASS
		SELECT USERCLASS
	SELECT NOTIFICATION
		SELECT USERNOTIFICATION

	UPDATE? USER
	UPDATE? INSTITUTION
	UPDATE? CLASS
	UPDATE? NOTIFICATION

	DELETE USER
	DELETE INSTITUTION
	DELETE CLASS
		DELETE USERCLASS
	DELETE NOTIFICATION
		DELETE USERNOTIFICATION

*/
