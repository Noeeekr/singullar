package operations

import "github.com/Noeeekr/singullar/server/internal/database/models"

type NotificationRequest struct {
	// For notification
	Title       string
	Description string
	IssuerId    int
}

type UsersNotificationRequest struct {
	// For usersNotifications
	NotificationId int
	TargetRole     models.UserRole
	TargetsIds     []int
}

type InstitutionRequest struct {
	Name string
	// for admin user
	Password string
	Email    string
}

func CreateNotificationRequest(title, description string, issuerId int, targetsIds []int, targetRole models.UserRole) *NotificationRequest {
	return &NotificationRequest{
		Title:       title,
		Description: description,
		IssuerId:    issuerId,
	}
}
func CreateUsersNotificationsRequest(notificationId int, targetRole models.UserRole, targetsIds ...int) *UsersNotificationRequest {
	return &UsersNotificationRequest{
		NotificationId: notificationId,
		TargetRole:     targetRole,
		TargetsIds:     targetsIds,
	}
}

func CreateInstitutionRequest(name, password, email string) *InstitutionRequest {
	return &InstitutionRequest{
		Name:     name,
		Password: password,
		Email:    email,
	}
}
