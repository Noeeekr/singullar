package seeder

import (
	"math/rand/v2"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"golang.org/x/crypto/bcrypt"
)

// Package seeder
func CreateUserRequests(amount, institutionId int) []*models.CreateUsers {
	users := make([]*models.CreateUsers, amount)

	for i := range amount {
		var password string
		for range amount {
			pwd, err := bcrypt.GenerateFromPassword([]byte(common.GenerateRandomStrings(10)), 10)
			if err == nil {
				password = string(pwd)
				break
			}
		}

		role := models.Student
		luckyNumber := rand.N(50)
		if luckyNumber > 45 {
			role = models.Supervisor
		}
		if luckyNumber < 10 {
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

func CreateNotificationRequests(amount, issuerId, targetId int, targetRole models.UserRole) []*operations.NotificationRequest {
	notifications := make([]*operations.NotificationRequest, amount)
	for i := range amount {
		notifications[i] = operations.CreateNotificationRequest(
			common.GenerateRandomStrings(10),
			common.GenerateRandomStrings(10),
			issuerId,
			targetId,
			targetRole,
		)
	}
	return notifications
}
