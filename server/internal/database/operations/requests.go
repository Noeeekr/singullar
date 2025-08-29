package operations

import "github.com/Noeeekr/singullar/server/internal/database/models"

type NotificationRequest struct {
	Content *models.CreateNotificationContents
	Users   []*models.CreateUsersNotifications
}

func CreateNotificationRequest(content *models.CreateNotificationContents, users []*models.CreateUsersNotifications) *NotificationRequest {
	return &NotificationRequest{
		Content: content,
		Users:   users,
	}
}

func CreateInstitutionRequest(name, password, email string) *models.CreateInstitutions {
	return &models.CreateInstitutions{
		Name:     name,
		Password: password,
		Email:    email,
	}
}
