package operations

import "github.com/Noeeekr/singullar/server/internal/database/models"

type NotificationRequest struct {
	Title       string
	Description string
	IssuerId    int
	TargetId    int
	TargetRole  models.UserRole
}

type InstitutionRequest struct {
	Name string
	// for admin user
	Password string
	Email    string
}

func CreateNotificationRequest(title, description string, issuerId, targetId int, targetRole models.UserRole) *NotificationRequest {
	return &NotificationRequest{
		Title:       title,
		Description: description,
		IssuerId:    issuerId,
		TargetId:    targetId,
		TargetRole:  targetRole,
	}
}

func CreateInstitutionRequest(name, password, email string) *InstitutionRequest {
	return &InstitutionRequest{
		Name:     name,
		Password: password,
		Email:    email,
	}
}
