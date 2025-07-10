package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type NotificationsTable struct {
	name         TableName
	Requests     *NotificationsRequests
	dependencies *TableDependencies
}

type NotificationsRequests struct {
	Create     *transactions.Request
	Drop       *transactions.Request
	InsertMany *transactions.Request
}

var notificationsTable *NotificationsTable = &NotificationsTable{
	name:         NotificationsTableName,
	dependencies: notificationsTableDependencies,
	Requests:     notificationsTableRequests,
}

var notificationsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{classesTable, institutionsTable},
}

var notificationsTableRequests = &NotificationsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			id   		 INT PRIMARY KEY,
			title		 VARCHAR(256) NOT NULL,
			description  VARCHAR(256) NOT NULL,

			CONSTRAINT fk_notifications FOREIGN KEY (id) REFERENCES %s (id)
		);
	`, NotificationsTableName, DefaultFieldsQuery, usersTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, NotificationsTableName)),
	InsertMany: transactions.NewRequest(fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, title, description, issuer_id)
		VALUES %s
		RETURNING created_at, updated_at, deleted_at, issuer_id, title, description;
	`, NotificationsTableName, placeholder)).AllowValueRepeat(placeholder, 5),
}

/*
Notification

	id
	title
	description

UserNotification

	notId
	usrId
	usrType
*/

type CreateNotifications struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IssuerId    int    `json:"issuerId" binding:"required"`
}

type Notifications struct {
	DefaultFields
	CreateNotifications
}

func (t *NotificationsTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *NotificationsTable) GetCreateRequest() *transactions.Request {
	return t.Requests.Create
}
func (t *NotificationsTable) GetDropRequest() *transactions.Request {
	return t.Requests.Drop
}
func (t *NotificationsTable) Name() TableName {
	return t.name
}
