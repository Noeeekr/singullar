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
	"github.com/Noeeekr/singullar/server/internal/database/seeder"
)

var databaseInformationFile = "../../../secrets/postgres.env"
var requiredTables = []models.TableMethods{
	models.UsersTable,
	models.InstitutionsTable,
	models.NotificationsTable,
	models.UsersNotificationsTable,
}
var migrationConfiguration = &migrations.Configuration{
	RecreateExisting: true,
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

type Utils struct {
	db *sql.DB

	MainTest *testing.T

	migrations *migrations.Migrations
	operations *operations.Operations
}

func (m *Utils) DropTables(tables []models.TableMethods) *common.Response {
	m.MainTest.Log("Finished operations, dropping tables")
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
func (m *Utils) CreateInstitutionsData(amount int) []*InstitutionData {
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
func (m *Utils) CreateUsersData(institutions []*models.Institutions) []*UserData {
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
func (m *Utils) CurrentDatabase() *sql.DB {
	return m.db
}
func (m *Utils) MustPrepareEnvironmentForTests() {
	m.MustConnectToDevelopmentDatabase()
	m.MustPingDatabase()
	m.MustPrepareTables()
}
func (m *Utils) MustConnectToDevelopmentDatabase() {
	var res *common.Response
	m.MainTest.Run("CONNECT TO DATABASE IN DEVELOPMENT ENVIRONMENT", func(t *testing.T) {
		res = environment.Parse(databaseInformationFile)
		if res != nil {
			t.Fatal(res.ParseToString())
		}

		m.db, res = connections.ConnectWithEnvironment(connections.Development)
		if res != nil {
			t.Fatal(res.ParseToString())
		}

		if m.db == nil {
			t.Fatal("Database pointer empty")
		}
		m.migrations = migrations.New(m.db)
		m.operations = operations.New(m.db)
	})
	if res != nil {
		m.MainTest.Fatal(res.ParseToString())
	}
}
func (m *Utils) MustPingDatabase() {
	var err error
	m.MainTest.Run("PING", func(t *testing.T) {
		err = m.db.Ping()
		if err != nil {
			t.Fatal("Unable to ping database. " + err.Error())
		}
	})
	if err != nil {
		m.MainTest.Fatal(err.Error())
	}
}
func (m *Utils) MustPrepareTables() {
	var res *common.Response
	m.MainTest.Run("CREATE TABLES", func(t *testing.T) {
		res = m.migrations.CreateTables(migrationConfiguration, requiredTables...)
		if res != nil {
			t.Fatal(res.ParseToString())
		}
		res = m.migrations.Commit()
		if res != nil {
			t.Fatal(res.ParseToString())
		}
	})
	if res != nil {
		m.MainTest.Fatal(res.ParseToString())
	}
}
func (m *Utils) MustPass(title string, fun func(t *testing.T)) {
	ok := m.MainTest.Run(title, fun)
	if !ok {
		m.MainTest.FailNow()
	}
}
func TestOperations(test *testing.T) {
	utils := Utils{
		MainTest: test,
	}

	utils.MustPrepareEnvironmentForTests()
	defer utils.DropTables(requiredTables)

	institutions_notifications := map[string][]*models.Notifications{}
	institutions_users := map[string][]*models.Users{}
	institutions := []*models.Institutions{}

	utils.MustPass("INSERT INSTITUTIONS", func(t *testing.T) {
		for _, data := range utils.CreateInstitutionsData(10) {
			institution := operations.CreateInstitutionRequest(data.Name, data.Email, data.Password)
			users, res := utils.operations.InsertInstitutions(institution)
			if res != nil {
				t.Fatal(res.ParseToString())
			}
			if res := utils.operations.Commit(); res != nil {
				t.Fatal(res.ParseToString())
			}
			institutions_users[institution.Name] = append(institutions_users[institution.Name], users...)
		}
	})
	utils.MustPass("SELECT INSTITUTIONS BY ID", func(t *testing.T) {
		for _, institution_users := range institutions_users {
			institution, res := utils.operations.SelectInstitutionById(institution_users[0].InstitutionId)
			if res != nil {
				t.Fatal(res.ParseToString())
			}
			if institution == nil {
				t.Fatal("Nil institution")
			}
			institutions = append(institutions, institution)
		}
	})
	utils.MustPass("SELECT INSTITUTIONS BY NAME", func(t *testing.T) {
		for _, created_institution := range institutions {
			institution, res := utils.operations.SelectInstitutionByName(created_institution.Name)
			if res != nil {
				t.Fatal(res.Description)
			}
			if institution == nil {
				t.Fatal("Nil institution")
			}
		}
	})

	utils.MustPass("INSERT USERS", func(t *testing.T) {
		for _, institution := range institutions {
			userRequests := seeder.CreateUserRequests(20, institution.Id)
			users, res := utils.operations.InsertManyUsers(userRequests...)
			if res != nil {
				t.Fatal(res.ParseToString())
			}
			if res := utils.operations.Commit(); res != nil {
				t.Fatal(res.ParseToString())
			}
			institutions_users[institution.Name] = append(institutions_users[institution.Name], users...)
		}
	})
	utils.MustPass("SELECT USERS BY EMAIL", func(t *testing.T) {
		for _, users := range institutions_users {
			for _, user := range users {
				_, res := utils.operations.SelectUserByEmail(user.Email)
				if res != nil {
					t.Fatal(res.Description)
				}
			}
		}
	})

	utils.MustPass("SELECT USERS BY ID", func(t *testing.T) {
		for _, users := range institutions_users {
			for _, user := range users {
				_, res := utils.operations.SelectUserById(user.Id)
				if res != nil {
					t.Fatal(res.Description)
				}
			}
		}
	})

	utils.MustPass("INSERT NOTIFICATIONS", func(t *testing.T) {
		for institution, users := range institutions_users {
			var teachers []*models.Users
			var student_ids []int
			for _, user := range users {
				if user.Role == models.Teacher {
					teachers = append(teachers, user)
				}
				if user.Role == models.Student {
					student_ids = append(student_ids, user.Id)
				}
			}
			for _, teacher := range teachers {
				notificationRequests := seeder.CreateNotificationRequests(10, teacher.Id)
				notifications, res := utils.operations.InsertNotifications(notificationRequests...)
				if res != nil {
					t.Fatal(res.ParseToString())
				}
				for _, notification := range notifications {
					requests := seeder.CreatedNotificationUserRequest(notification.Id, models.Student, student_ids...)
					res := utils.operations.InsertUsersNotifications(requests...)
					if res != nil {
						t.Fatal(res.ParseToString())
					}
				}
				institutions_notifications[institution] = append(institutions_notifications[institution], notifications...)
			}
			res := utils.operations.Commit()
			if res != nil {
				t.Fatal(res.ParseToString())
			}
		}
	})

	utils.MustPass("SELECT NOTIFICATIONS BY TARGET ID", func(t *testing.T) {
		var student_ids []int
		for _, users := range institutions_users {
			for _, user := range users {
				if user.Role == models.Student {
					student_ids = append(student_ids, user.Id)
				}
			}
		}
		for _, id := range student_ids {
			notifications, res := utils.operations.SelectNotificationsByTargetId(id)
			if res != nil {
				t.Fatal(res.ParseToString())
			}
			if len(*notifications) == 0 {
				t.Fatal("Notifications not found")
			}
		}
	})
	utils.MustPass("DELETE NOTIFICATIONS BY ISSUER ID", func(t *testing.T) {
		for _, notifications := range institutions_notifications {
			for _, notification := range notifications {
				res := utils.operations.DeleteNotificationsByIssuerId(notification.IssuerId)
				if res != nil {
					t.Fatal(res.ParseToString())
				}
			}
		}
		res := utils.operations.Commit()
		if res != nil {
			t.Fatal(res.ParseToString())
		}
	})

	utils.MustPass("DELETE USERS BY EMAIL AND ID", func(t *testing.T) {
		for _, users := range institutions_users {
			if len(users) < 2 {
				t.Fatal("Insufficient test cases for users")
			}
			for i, user := range users {
				var res *common.Response
				if i > len(users)/2 {
					res = utils.operations.DeleteUserByEmail(user.Email).Response
				} else {
					res = utils.operations.DeleteUserById(user.Id).Response
				}
				if res != nil {
					t.Fatal(res.Status, res.Description)
				}
			}
			if res := utils.operations.Commit(); res != nil {
				t.Fatal(res.ParseToString())
			}
		}
	})
	utils.MustPass("DELETE INSTITUTIONS BY EMAIL AND ID", func(t *testing.T) {
		if len(institutions) < 2 {
			t.Fatal("Insufficient test cases for institutions")
		}
		var res *common.Response
		for i, institution := range institutions {
			if i > len(institutions)/2 {
				res = utils.operations.DeleteInstitutionById(nil, institution.Id).Response
			} else {
				res = utils.operations.DeleteInstitutionByName(nil, institution.Name).Response
			}
			if res != nil {
				t.Fatal(res.Status, res.Description)
			}
		}
		if res := utils.operations.Commit(); res != nil {
			t.Fatal(res.ParseToString())
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
