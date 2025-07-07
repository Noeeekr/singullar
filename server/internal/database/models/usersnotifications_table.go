package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type UsersNotifications struct {
	UserId         int
	NotificationId int
}

type UsersNotificationsTable struct {
	TableMethods
	name         TableName
	requests     *UsersNotificationsRequests
	dependencies *TableDependencies
}

type UsersNotificationsRequests struct {
	Create *transactions.TransactionRequest
	Drop   *transactions.TransactionRequest
}

var usersNotificationsTable *UsersNotificationsTable = &UsersNotificationsTable{
	name:         UsersNotificationsTableName,
	dependencies: usersNotificationsTableDependencies,
	requests:     usersNotificationsTableRequests,
}

var usersNotificationsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{usersTable, notificationsTable},
}

var usersNotificationsTableRequests = &UsersNotificationsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			notification_id INT NOT NULL,

			FOREIGN KEY (user_id) REFERENCES %s(id),
			FOREIGN KEY (notification_id) REFERENCES %s(id)
		);
	`, UsersNotificationsTableName, usersTableName, NotificationsTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, UsersNotificationsTableName)),
}

func (t *UsersNotificationsTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *UsersNotificationsTable) CreateRequest() *transactions.TransactionRequest {
	return t.requests.Create
}

func (t *UsersNotificationsTable) DropRequest() *transactions.TransactionRequest {
	return t.requests.Drop
}

func (t *UsersNotificationsTable) Name() TableName {
	return t.name
}
