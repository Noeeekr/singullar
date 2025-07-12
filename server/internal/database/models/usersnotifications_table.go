package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type CreateUsersNotifications struct {
	// Notification can be sent to a Class using TargetRole=Unkown and TargetId=ClassID
	// Notification can be sent to a Roles using TargetRole=Role and TargetId=InstitutionID
	UserRole UserRole `json:"targetRole" binding:"required"`
	UserId   int      `json:"targetId" binding:"required"`

	// The id of the notification is the id of its creator
	NotificationId int `json:"notificationId" binding:"required"`
}
type UsersNotifications struct {
	CreateUsersNotifications
}
type UsersNotificationsTableInformation struct {
	TableMethods
	name         TableName
	Requests     *UsersNotificationsRequests
	dependencies *TableDependencies
}

type UsersNotificationsRequests struct {
	Create     *transactions.Request
	Drop       *transactions.Request
	InsertMany *transactions.Request
	// DELETE should be done via deleting the notification
}

var UsersNotificationsTable *UsersNotificationsTableInformation = &UsersNotificationsTableInformation{
	name:         UsersNotificationsTableName,
	dependencies: usersNotificationsTableDependencies,
	Requests:     usersNotificationsTableRequests,
}

var usersNotificationsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{UsersClassesTable, NotificationsTable},
}

var usersNotificationsTableRequests = &UsersNotificationsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			target_role %s NOT NULL,
			user_id INT NOT NULL,
			notification_id INT NOT NULL,  

			CONSTRAINT fk_users_notifications_target_id 
			FOREIGN KEY (user_id) 
			REFERENCES %s(id),

			CONSTRAINT fk_users_notifications_source_id 
			FOREIGN KEY (notification_id) 
			REFERENCES %s(id)
			ON DELETE CASCADE
		);
	`, UsersNotificationsTableName, UserRolesTypeName, usersTableName, NotificationsTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, UsersNotificationsTableName)),
	InsertMany: transactions.NewRequest(fmt.Sprintf(`
		INSERT INTO %s (user_id, target_role, notification_id)
		VALUES %s
		RETURNING user_id, target_role, notification_id;
	`, UsersNotificationsTableName, placeholder)).AllowValueRepeat(placeholder, 3),
}

func (t *UsersNotificationsTableInformation) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *UsersNotificationsTableInformation) GetCreateRequest() *transactions.Request {
	return t.Requests.Create
}

func (t *UsersNotificationsTableInformation) GetDropRequest() *transactions.Request {
	return t.Requests.Drop
}

func (t *UsersNotificationsTableInformation) Name() TableName {
	return t.name
}
