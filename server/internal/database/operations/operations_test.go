package operations_test

import (
	"database/sql"
	"testing"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/migrations"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
)

var flagEnvironmentFile = "../../../secrets/postgres.env"
var tables = []models.TableMethods{
	models.TablesInfo.Users,
	models.TablesInfo.Institutions,
}

type InstitutionData struct {
	Name     string
	Email    string
	Password string
}

type UserData struct {
	Name          string
	Email         string
	Password      string
	Role          models.UserRole
	InstitutionId int
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

	migrations *migrations.Migrations
	operations *operations.Operations
}

func (m *QueryTestUtil) CreateTables(configuration *migrations.Configuration, tables []models.TableMethods) (res *common.Response) {
	transaction := m.migrations.CreateTables(configuration, tables...)
	if transaction.Response != nil {
		m.test.Log(transaction.Response.ParseToString())
	}
	return transaction.Response
}
func (m *QueryTestUtil) DropTables(tables []models.TableMethods) *common.Response {
	m.test.Log("Finished operations, dropping tables")
	tx := m.migrations.DropTables(tables...)
	if tx.Response != nil {
		logs.Error.Fatal(tx.Response.ParseToString())
	}
	if tx := tx.Commit(); tx.Response != nil {
		logs.Error.Fatal(tx.Response.ParseToString())
	}
	defer m.db.Close()
	return nil
}
func (m *QueryTestUtil) CreateInstitutionsData(amount int) []*InstitutionData {
	data := []*InstitutionData{}
	for range amount {
		data = append(
			data,
			&InstitutionData{
				Name:     common.GenerateRandomStrings(10),
				Email:    common.GenerateRandomStrings(10),
				Password: common.GenerateRandomStrings(10),
			},
		)
	}
	return data
}
func (m *QueryTestUtil) CreateUsersData(institutions []*models.Institutions) []*UserData {
	data := []*UserData{}
	for _, institution := range institutions {
		data = append(
			data,
			&UserData{
				Name:          common.GenerateRandomStrings(10),
				Email:         common.GenerateRandomStrings(10),
				Password:      common.GenerateRandomStrings(10),
				Role:          models.Student,
				InstitutionId: institution.Id,
			},
		)
	}
	return data
}

// Start database connection, create necessary tables.
func (m *QueryTestUtil) PrepareDatabaseWithTables(configuration *migrations.Configuration, tables []models.TableMethods) *sql.DB {
	m.test.Run("CONNECT INTO ENVIRONMENT DATABASE AS TEST USER", func(t *testing.T) {
		if err := environment.Parse(flagEnvironmentFile); err != nil {
			t.Fatal(err.Status, err.Description)
		}

		db, res := connections.ConnectWithEnvironment(connections.Development)
		if res != nil {
			t.Fatal(res.ParseToError().Error())
		}

		m.db = db
		m.migrations = migrations.New(db)
		m.operations = operations.New(db)
	})

	m.Run("PING", func(t *testing.T) {
		err := m.db.Ping()
		if err != nil {
			t.Fatal("Unable to ping database. " + err.Error())
		}
	})

	m.Run("CREATE TABLES", func(t *testing.T) {
		if res := m.CreateTables(configuration, tables); res != nil {
			t.Fatal(res.Status, res.Description)
		}
		if res := m.migrations.Commit(); res != nil {
			t.Fatal(res.Status, res.Description)
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

func TestOperations(test *testing.T) {
	utils := QueryTestUtil{
		test: test,
	}

	configuration := migrations.Configuration{
		RecreateExisting: true,
	}

	db := utils.PrepareDatabaseWithTables(&configuration, tables)
	if db == nil {
		test.Fatal("Database not prepared")
	}
	defer utils.DropTables(tables)

	institution_source_data := utils.CreateInstitutionsData(10)

	created_users := []*models.Users{}
	created_institutions := []*models.Institutions{}

	var res *common.Response

	utils.
		Run("INSERT INSTITUTIONS", func(t *testing.T) {
			for _, data := range institution_source_data {
				institution := operations.CreateInstitutionRequest(data.Name, data.Email, data.Password)
				users, tx := utils.operations.InsertInstitutions(institution)
				if tx.Response != nil {
					t.Fatal(tx.Response.ParseToString())
				}
				if res := utils.operations.Commit(); res != nil {
					t.Fatal(res.ParseToString())
				}
				created_users = append(created_users, users...)
			}
		}).
		Run("SELECT INSTITUTIONS BY ID", func(t *testing.T) {
			for _, user := range created_users {
				institution, res := utils.operations.SelectInstitutionById(user.InstitutionId)
				if res != nil {
					t.Fatal(res.ParseToString())
				}
				if institution == nil {
					t.Fatal("Nil institution")
				}
				created_institutions = append(created_institutions, institution)
			}
		}).
		Run("SELECT INSTITUTIONS BY NAME", func(t *testing.T) {
			for _, data := range institution_source_data {
				institution, res := utils.operations.SelectInstitutionByName(data.Name)
				if res != nil {
					t.Fatal(res.Description)
				}
				if institution == nil {
					t.Fatal("Nil institution")
				}
			}
		})

	utils.Run("INSERT USERS", func(t *testing.T) {
		for _, institution := range created_institutions {
			email := common.GenerateRandomStrings(10)
			users, tx := utils.operations.InsertManyUsers(&models.CreateUsers{
				Name:          common.GenerateRandomStrings(10),
				Email:         email,
				Password:      common.GenerateRandomStrings(10),
				InstitutionId: institution.Id,
				Role:          models.Student,
			})
			if tx.Response != nil {
				t.Fatal(tx.Response.ParseToString())
			}
			if len(users) != 1 {
				t.Fatal("User not returned correctly")
			}
			if users[0].Email != email {
				t.Fatal("Emails don't match")
			}
			if res := utils.operations.Commit(); res != nil {
				t.Fatal(res.ParseToString())
			}
			created_users = append(created_users, users[0])
		}
	}).Run("SELECT USERS BY EMAIL", func(t *testing.T) {
		for _, user := range created_users {
			user, res := utils.operations.SelectUserByEmail(user.Email)
			if res != nil {
				t.Fatal(res.Description)
			}
			if user == nil {
				t.Fatal("Nil user")
			}
		}
	}).Run("SELECT USERS BY ID", func(t *testing.T) {
		for _, user := range created_users {
			user, res = utils.operations.SelectUserById(user.Id)
			if res != nil {
				t.Fatal(res.Description)
			}
			if user == nil {
				t.Fatal("Nil user")
			}
		}
	})

	utils.Run("DELETE USERS BY EMAIL AND ID", func(t *testing.T) {
		if len(created_users) < 2 {
			t.Fatal("Insufficient test cases for users")
		}
		for i, user := range created_users {
			if i > len(created_users)/2 {
				res = utils.operations.DeleteUserByEmail(user.Email).Response
			} else {
				res = utils.operations.DeleteUserById(user.Id).Response
			}
			if res != nil {
				test.Fatal(res.Status, res.Description)
			}
		}
	})
	utils.Run("DELETE INSTITUTIONS BY EMAIL AND ID", func(t *testing.T) {
		if len(created_institutions) < 2 {
			t.Fatal("Insufficient test cases for institutions")
		}
		for i, institution := range created_institutions {
			if i > len(created_institutions)/2 {
				res = utils.operations.DeleteInstitutionById(nil, institution.Id).Response
			} else {
				res = utils.operations.DeleteInstitutionByName(nil, institution.Name).Response
			}
			if res != nil {
				test.Fatal(res.Status, res.Description)
			}
		}
		if res := utils.operations.Commit(); res != nil {
			test.Fatal(res.ParseToString())
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
