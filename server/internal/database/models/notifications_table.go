package models

import (
	"fmt"
	"time"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type NotificationsTableInformation struct {
	name         TableName
	Requests     *NotificationsRequests
	dependencies *TableDependencies
}

type NotificationsRequests struct {
	Create           *transactions.Request
	Drop             *transactions.Request
	SelectByTargetId *transactions.Request
	InsertMany       *transactions.Request
	DeleteByIssuerId *transactions.Request
}

var NotificationsTable *NotificationsTableInformation = &NotificationsTableInformation{
	name:         NotificationsTableName,
	dependencies: notificationsTableDependencies,
	Requests:     notificationsTableRequests,
}

var notificationsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{ClassesTable, InstitutionsTable},
}

var notificationsTableRequests = &NotificationsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			%s
			issuer_id    INT		  NOT NULL,
			title		 VARCHAR(256) NOT NULL,
			description  VARCHAR(256) NOT NULL,

			CONSTRAINT fk_notifications 
			FOREIGN KEY (issuer_id) 
			REFERENCES %s (id)
			ON DELETE CASCADE
		);
	`, NotificationsTableName, DefaultFieldsQuery, SerialId, usersTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, NotificationsTableName)),
	InsertMany: transactions.NewRequest(fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, title, description, issuer_id)
		VALUES %s
		RETURNING created_at, updated_at, deleted_at, id, issuer_id, title, description;
	`, NotificationsTableName, placeholder)).AllowValueRepeat(placeholder, 5),
	DeleteByIssuerId: transactions.NewRequest(fmt.Sprintf(`
		DELETE FROM %s WHERE issuer_id = $1
	`, NotificationsTableName)),
	SelectByTargetId: transactions.NewRequest(fmt.Sprintf(`
		SELECT n.created_at, n.updated_at, n.deleted_at, u.id, u.name, n.id, n.title, n.description
		FROM %s n
		INNER JOIN %s un ON un.notification_id = n.id
		INNER JOIN %s u ON u.id = un.user_id
		WHERE u.id = $1;
	`, NotificationsTableName, UsersNotificationsTableName, usersTableName)),
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
	ID
	DefaultFields
	CreateNotifications
}

type DetailedNotifications struct {
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	TargetId    int
	TargetName  string
	IssuerId    int
	Title       string
	Description string
}

func (t *NotificationsTableInformation) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *NotificationsTableInformation) GetCreateRequest() *transactions.Request {
	return t.Requests.Create
}
func (t *NotificationsTableInformation) GetDropRequest() *transactions.Request {
	return t.Requests.Drop
}
func (t *NotificationsTableInformation) Name() TableName {
	return t.name
}
