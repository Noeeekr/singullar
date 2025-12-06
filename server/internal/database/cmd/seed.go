package cmd

import (
	"fmt"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/common/commandutil"
	"github.com/Noeeekr/singullar/server/internal/database/manager"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/seeder"
	"github.com/spf13/cobra"
)

var seedCmd *cobra.Command = &cobra.Command{
	Use:   "seed [ -f ENVIRONMENT_FILE_PATH... ]",
	Short: "Creates and seeds a single institution with the given name into the development environment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		environment.
			Settings().
			SetApplicationMode(environment.DEVELOPMENT)
		borm.
			Settings().
			Migrations().
			Enable().RecreateExisting().UndoOnError()

		files, _ := cmd.Flags().GetStringArray("environmentFiles")
		if err := environment.Parse(files...); err != nil {
			logs.Error.Fatal(err.String())
		}

		if err := seedInstitution(); err != nil {
			logs.Error.Fatal(err)
		}
		logs.Info.Println("|\n[Finished seeding]")
	},
}

func seedInstitution() error {
	institutionRequest := seeder.CreateInstitutionRequest(1)[0]
	institutionRequest.Name = "Test Institution"
	institutionRequest.Email = "admin@test.com"
	institutionRequest.Password = "admin"
	fmt.Println("Institution Name: ", institutionRequest.Name)
	fmt.Println("[Test-Admin] Email:", institutionRequest.Email)
	fmt.Println("[Test-Admin] Password:", institutionRequest.Password)

	commiter, err := borm.Connect(models.EnvironmentDatabase)
	if err != nil {
		return err
	}

	manager := manager.New(commiter)
	operator, err := manager.NewTransactionOperator()
	if err != nil {
		return err
	}

	admin, err := InsertInstitution(operator, institutionRequest)
	if err != nil {
		logs.Error.Fatal(err)
	}

	teachers, studentIds, err := InsertUsers(operator, admin.InstitutionId)
	if err != nil {
		logs.Error.Fatal(err)
	}

	err = InsertNotifications(operator, teachers, studentIds)
	if err != nil {
		logs.Error.Fatal(err)
	}

	return operator.Commit()
}
func InsertInstitution(operator *manager.Operator, r *models.CreateInstitutions) (*models.Users, error) {
	// Create institution
	admins, err := operator.InsertInstitutions(r)
	if err != nil {
		return nil, err
	}
	return admins[0], nil
}
func InsertUsers(operator *manager.Operator, institutionId int) (teachers []*models.Users, studentIds []int, err error) {
	// Create students, teachers and supervisors
	segment := models.EF2
	testuser := models.CreateUser("student", "student@test.com", "student", institutionId, models.STUDENT, &segment)
	fmt.Println("|")
	fmt.Println("[Test-Student] Email:", testuser.Email)
	fmt.Println("[Test-Student] Password:", testuser.Password)

	createdUsers := seeder.CreateUserRequests(10, institutionId)
	createdUsers = append(createdUsers, testuser)

	users, err := operator.InsertManyUsers(createdUsers...)
	if err != nil {
		return teachers, studentIds, err
	}

	for _, user := range users {
		switch user.Role {
		case models.TEACHER:
			teachers = append(teachers, user)
		case models.STUDENT:
			studentIds = append(studentIds, user.Id)
		}
	}

	return teachers, studentIds, err
}
func InsertNotifications(operator *manager.Operator, teachers []*models.Users, studentIds []int) error {
	notificationRequests := []*manager.NotificationRequest{}
	// Iterate over the teachers creations 10 notifications for each
	for _, teacher := range teachers {
		// Iterate over all notifications appending them to all users
		notificationContentRequests := seeder.CreateNotificationContentRequests(10, teacher.Id)
		for _, notificationContentRequest := range notificationContentRequests {
			usersNotificationsRequests := seeder.CreateUsersNotificationsRequests(models.STUDENT, studentIds...)
			notificationRequest := manager.CreateNotificationRequest(notificationContentRequest, usersNotificationsRequests)
			notificationRequests = append(notificationRequests, notificationRequest)
		}
	}
	_, err := operator.InsertNotifications(notificationRequests...)
	if err != nil {
		return err
	}
	return nil
}
func init() {
	commandutil.ConsumeFlagConfiguration(environment.EnvironmentFilesFlagToken, seedCmd)

	rootCmd.AddCommand(seedCmd)
}
