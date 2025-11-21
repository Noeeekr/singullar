package manager_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/database/manager"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/seeder"
)

type InstitutionData struct {
	Info  *models.Institutions
	Admin *models.Users

	Users         []*models.Users
	Notifications []*models.NotificationContents
}
type InstitutionNameToData map[string]*InstitutionData

type Utils struct {
	db *sql.DB

	MainTest *testing.T

	Commiter        *borm.Commiter
	DatabaseManager *manager.DatabaseManager
}

// Test Utility Functions

func (m *Utils) ClearEnvironment() *common.Response {
	m.MainTest.Log("Finished DatabaseManager, dropping tables")
	err := m.Commiter.DropRelations()
	if err != nil {
		logs.Error.Fatal(err)
	}
	return nil
}
func (m *Utils) PrepareEnvironment() error {
	steps := []func() string{
		m.ConnectToDevelopmentDatabase,
		m.PingDatabase,
		m.PrepareTables,
	}

	for _, step := range steps {
		if err := step(); err != "" {
			return errors.New(err)
		}
	}
	return nil
}
func (m *Utils) ConnectToDevelopmentDatabase() string {
	var res *common.Response
	m.MainTest.Run("PARSE ENVIRONMENT", func(t *testing.T) {
		res = environment.Parse(databaseSecretsFilepath)
		if res != nil {
			t.Fatal(res.String())
		}
	})
	if res != nil {
		return res.String()
	}

	var err error
	m.MainTest.Run("CONNECT TO DATABASE IN DEVELOPMENT ENVIRONMENT", func(t *testing.T) {
		var commiter *borm.Commiter
		commiter, err = borm.Connect(models.EnvironmentDatabase)
		if err != nil {
			t.Fatal(err)
		}

		m.db = commiter.DB()
		m.Commiter = commiter
		m.DatabaseManager = manager.New(commiter)
	})
	if err != nil {
		return err.Error()
	}
	return ""
}
func (m *Utils) PingDatabase() string {
	var err error
	m.MainTest.Run("PING DATABASE", func(t *testing.T) {
		err = m.db.Ping()
		if err != nil {
			t.Fatal("Unable to ping database. " + err.Error())
		}
	})
	if err != nil {
		return err.Error()
	}
	return ""
}
func (m *Utils) PrepareTables() string {
	var err error
	m.MainTest.Run("CREATE TABLES", func(t *testing.T) {
		borm.Settings().Migrations().Enable().RecreateExisting().UndoOnError()
		err = m.Commiter.MigrateRelations()
		if err != nil {
			t.Fatal(err)
		}
	})
	if err != nil {
		return err.Error()
	}
	return ""
}
func (m *Utils) MustPass(title string, fun func(t *testing.T)) {
	ok := m.MainTest.Run(title, fun)
	if !ok {
		m.MainTest.FailNow()
	}
}
func NewUtil(test *testing.T) *Utils {
	return &Utils{
		MainTest: test,
	}
}

// Test Configuration

var databaseSecretsFilepath = "../../../secrets/postgres.env"

// Test Data
var institutions = InstitutionNameToData{}

func TestDatabaseOperations(test *testing.T) {
	testutil := NewUtil(test)
	if err := testutil.PrepareEnvironment(); err != nil {
		test.Fatal(err)
	}
	defer testutil.ClearEnvironment()

	testutil.MustPass("INSERT INSTITUTIONS", func(t *testing.T) {
		operator, err := testutil.DatabaseManager.NewTransactionOperator()
		if err != nil {
			t.Fatal(err)
		}
		for _, data := range seeder.CreateInstitutionRequest(10) {
			institution := manager.CreateInstitutionRequest(data.Name, data.Email, data.Password)
			createdUsers, err := operator.InsertInstitutions(institution)
			if err != nil {
				t.Fatal(err)
			}

			admin := createdUsers[0]

			institutions[institution.Name] = &InstitutionData{
				Info:          nil,
				Admin:         admin,
				Notifications: []*models.NotificationContents{},
				Users:         []*models.Users{},
			}
		}
		if err := operator.Commit(); err != nil {
			t.Fatal(err)
		}
	})
	testutil.MustPass("SELECT INSTITUTIONS BY ID", func(t *testing.T) {
		for _, institution := range institutions {
			institution, err := testutil.DatabaseManager.SelectInstitutionById(institution.Admin.InstitutionId)
			if err != nil {
				t.Fatal(err)
			}
			if institutions[institution.Name] == nil {
				t.Fatal("[Target institution name] and [Recovered institution name] doesn't match.")
			}
			institutions[institution.Name].Info = institution
		}
	})
	testutil.MustPass("SELECT INSTITUTIONS BY NAME", func(t *testing.T) {
		for _, institution := range institutions {
			_, err := testutil.DatabaseManager.SelectInstitutionByName(institution.Info.Name)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	testutil.MustPass("INSERT USERS", func(t *testing.T) {
		operator, err := testutil.DatabaseManager.NewTransactionOperator()
		if err != nil {
			t.Fatal(err)
		}
		for _, institution := range institutions {
			userRequests := seeder.CreateUserRequests(20, institution.Info.Id)
			users, err := operator.InsertManyUsers(userRequests...)
			if err != nil {
				t.Fatal(err)
			}
			institution.Users = append(institution.Users, users...)
		}
		if err := operator.Commit(); err != nil {
			t.Fatal(err)
		}
	})
	testutil.MustPass("SELECT USERS BY EMAIL", func(t *testing.T) {
		for _, institution := range institutions {
			for _, user := range institution.Users {
				_, err := testutil.DatabaseManager.SelectUserByEmail(user.Email)
				if err != nil {
					t.Fatal(err)
				}
			}
		}
	})

	testutil.MustPass("SELECT USERS BY ID", func(t *testing.T) {
		for _, institution := range institutions {
			for _, user := range institution.Users {
				_, err := testutil.DatabaseManager.SelectUsersById(user.Id)
				if err != nil {
					t.Fatal(err)
				}
			}
		}
	})

	testutil.MustPass("INSERT NOTIFICATIONS", func(t *testing.T) {
		operator, err := testutil.DatabaseManager.NewTransactionOperator()
		if err != nil {
			t.Fatal(err)
		}
		for _, institution := range institutions {
			// Go through every institution and separe the teachers
			var teachers []*models.Users
			var student_ids []int
			for _, user := range institution.Users {
				if user.Role == models.TEACHER {
					teachers = append(teachers, user)
				}
				if user.Role == models.STUDENT {
					student_ids = append(student_ids, user.Id)
				}
			}

			// Create 10 notification content from every teacher
			notificationRequests := []*manager.NotificationRequest{}
			for _, teacher := range teachers {
				notificationContentRequests := seeder.CreateNotificationContentRequests(10, teacher.Id)

				// Append the 10 notications from every teacher to all users
				for _, notificationContentRequest := range notificationContentRequests {
					usersNotificationsRequest := seeder.CreateUsersNotificationsRequests(models.STUDENT, student_ids...)
					notificationRequests = append(notificationRequests, manager.CreateNotificationRequest(notificationContentRequest, usersNotificationsRequest))
				}
			}
			notifications, err := operator.InsertNotifications(notificationRequests...)
			if err != nil {
				t.Fatal(err)
			}
			institution.Notifications = append(institution.Notifications, notifications...)
		}
		if err := operator.Commit(); err != nil {
			t.Fatal(err)
		}
	})

	testutil.MustPass("SELECT NOTIFICATIONS BY TARGET ID", func(t *testing.T) {
		var student_ids []int
		for _, institution := range institutions {
			for _, user := range institution.Users {
				if user.Role == models.STUDENT {
					student_ids = append(student_ids, user.Id)
				}
			}
		}
		for _, id := range student_ids {
			notifications, err := testutil.DatabaseManager.SelectNotificationsByTargetId(id)
			if err != nil {
				t.Fatal(err)
			}
			if len(*notifications) == 0 {
				t.Fatal("Notifications not found")
			}
		}
	})
	testutil.MustPass("DELETE NOTIFICATIONS BY ISSUER ID", func(t *testing.T) {
		operator, err := testutil.DatabaseManager.NewTransactionOperator()
		if err != nil {
			t.Fatal(err)
		}
		for _, institution := range institutions {
			for _, notification := range institution.Notifications {
				err := operator.DeleteNotificationsByIssuerId(notification.IssuerId)
				if err != nil {
					t.Fatal(err)
				}
			}
		}
		if err := operator.Commit(); err != nil {
			t.Fatal(err)
		}
	})

	testutil.MustPass("DELETE USERS BY EMAIL AND ID", func(t *testing.T) {
		operator, err := testutil.DatabaseManager.NewTransactionOperator()
		if err != nil {
			t.Fatal(err)
		}
		for _, institution := range institutions {
			userAmount := len(institution.Users)
			if userAmount < 2 {
				t.Fatal("Insufficient test cases for users")
			}
			for i, user := range institution.Users {
				var err error
				if i > userAmount/2 {
					err = operator.DeleteUserByEmail(user.Email)
				} else {
					err = operator.DeleteUserById(user.Id)
				}
				if err != nil {
					t.Fatal(err)
				}
			}
		}
		if err := operator.Commit(); err != nil {
			t.Fatal(err)
		}
	})
	testutil.MustPass("DELETE INSTITUTIONS BY EMAIL AND ID", func(t *testing.T) {
		operator, err := testutil.DatabaseManager.NewTransactionOperator()
		if err != nil {
			t.Fatal(err)
		}
		institutionAmount := len(institutions)
		if institutionAmount < 2 {
			t.Fatal("Insufficient test cases for institutions")
		}
		var index int = 0
		for _, institution := range institutions {
			if index > institutionAmount/2 {
				err = operator.DeleteInstitutionById(institution.Info.Id)
			} else {
				err = operator.DeleteInstitutionByName(institution.Info.Name)
			}
			if err != nil {
				t.Fatal(err)
			}
			index++
		}
		if err := operator.Commit(); err != nil {
			t.Fatal(err)
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
