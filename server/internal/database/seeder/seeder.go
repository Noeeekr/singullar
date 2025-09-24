package seeder

import (
	"math/rand/v2"

	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/util"
)

func CreateInstitutionRequest(amount int) []*models.CreateInstitutions {
	data := []*models.CreateInstitutions{}
	for range amount {
		data = append(
			data,
			&models.CreateInstitutions{
				Name:     util.GenerateRandomStrings(10),
				Email:    util.GenerateRandomStrings(10),
				Password: util.GenerateRandomStrings(10),
			},
		)
	}
	return data
}
func CreateUserRequests(amount, institutionId int) []*models.CreateUsers {
	users := make([]*models.CreateUsers, amount)

	for i := range amount {
		password := util.GenerateRandomStrings(10)
		role := models.STUDENT
		luckyNumber := rand.N(50)
		if luckyNumber > 45 {
			role = models.SUPERVISOR
		}
		if luckyNumber <= 10 {
			role = models.TEACHER
		}
		segment := models.EF1
		users[i] = models.CreateUser(
			util.GenerateRandomStrings(10),
			util.GenerateRandomStrings(10),
			password,
			institutionId,
			role,
			&segment,
		)
	}

	return users
}

func CreateNotificationContentRequests(amount, issuerId int) []*models.CreateNotificationContents {
	notifications := make([]*models.CreateNotificationContents, amount)
	for i := range amount {
		notifications[i] = &models.CreateNotificationContents{
			Title:       util.GenerateRandomStrings(10),
			Description: util.GenerateRandomStrings(10),
			IssuerId:    issuerId,
		}
	}
	return notifications
}

func CreateUsersNotificationsRequests(userRole models.UserRole, usersIds ...int) []*models.CreateUsersNotifications {
	usersNotifications := make([]*models.CreateUsersNotifications, len(usersIds))
	for i, id := range usersIds {
		usersNotifications[i] = &models.CreateUsersNotifications{
			UserRole: userRole,
			UserId:   id,
		}
	}
	return usersNotifications
}
