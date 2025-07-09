package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type CreateUsersNotifications struct {
	// Notification can be sent to a Class using TargetRole=Unkown and TargetId=ClassID
	// Notification can be sent to a Roles using TargetRole=Role and TargetId=InstitutionID
	TargetRole UserRole `json:"targetRole" binding:"required"`
	TargetId   int      `json:"targetId" binding:"required"`

	// The id of the notification is the id of its creator
	NotificationId int `json:"notificationId" binding:"required"`
}

type UsersNotificationsTable struct {
	TableMethods
	name         TableName
	Requests     *UsersNotificationsRequests
	dependencies *TableDependencies
}

type UsersNotificationsRequests struct {
	Create     *transactions.Request
	Drop       *transactions.Request
	InsertMany *transactions.Request
}

var usersNotificationsTable *UsersNotificationsTable = &UsersNotificationsTable{
	name:         UsersNotificationsTableName,
	dependencies: usersNotificationsTableDependencies,
	Requests:     usersNotificationsTableRequests,
}

var usersNotificationsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{usersTable, notificationsTable},
}

var usersNotificationsTableRequests = &UsersNotificationsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			target_id INT NOT NULL,
			target_role %s NOT NULL,
			notification_id INT NOT NULL,  

			FOREIGN KEY (target_id) REFERENCES %s(id),
			FOREIGN KEY (notification_id) REFERENCES %s(issuer_id)
		);
	`, UsersNotificationsTableName, UserRolesTypeName, usersTableName, NotificationsTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, UsersNotificationsTableName)),
	InsertMany: transactions.NewRequest(fmt.Sprintf(`
		INSERT INTO %s (target_id, target_role, notification_id)
		VALUES %s
		RETURNING target_id, target_role, notification_id;
	`, UsersNotificationsTableName, placeholder)).AllowValueRepeat(placeholder, 3),
}

func (t *UsersNotificationsTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *UsersNotificationsTable) GetCreateRequest() *transactions.Request {
	return t.Requests.Create
}

func (t *UsersNotificationsTable) GetDropRequest() *transactions.Request {
	return t.Requests.Drop
}

func (t *UsersNotificationsTable) Name() TableName {
	return t.name
}
