package seeder

import (
	"math/rand/v2"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

// Package seeder
func CreateUserRequests(amount, institutionId int) []*models.CreateUsers {
	users := make([]*models.CreateUsers, amount)

	for i := range amount {
		password := common.GenerateRandomStrings(10)
		role := models.Student
		luckyNumber := rand.N(50)
		if luckyNumber > 45 {
			role = models.Supervisor
		}
		if luckyNumber <= 10 {
			role = models.Teacher
		}
		users[i] = models.CreateUser(
			common.GenerateRandomStrings(10),
			common.GenerateRandomStrings(10),
			password,
			institutionId,
			role,
		)
	}

	return users
}

func CreateNotificationRequests(amount, issuerId int) []*models.CreateNotifications {
	notifications := make([]*models.CreateNotifications, amount)
	for i := range amount {
		notifications[i] = &models.CreateNotifications{
			Title:       common.GenerateRandomStrings(10),
			Description: common.GenerateRandomStrings(10),
			IssuerId:    issuerId,
		}
	}
	return notifications
}

func CreatedNotificationUserRequest(notificationId int, userRole models.UserRole, usersIds ...int) []*models.CreateUsersNotifications {
	usersNotifications := make([]*models.CreateUsersNotifications, len(usersIds))
	for i, id := range usersIds {
		usersNotifications[i] = &models.CreateUsersNotifications{
			UserRole:       userRole,
			NotificationId: notificationId,
			UserId:         id,
		}
	}
	return usersNotifications
}
