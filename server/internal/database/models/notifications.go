package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type NotificationsTable struct {
	name         TableName
	requests     *NotificationsRequests
	dependencies *TableDependencies
}

type NotificationsRequests struct {
	Create *transactions.TransactionRequest
	Drop   *transactions.TransactionRequest
}

var notificationsTable *NotificationsTable = &NotificationsTable{
	name:         NotificationsTableName,
	dependencies: notificationsTableDependencies,
	requests:     notificationsTableRequests,
}

var notificationsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{classesTable, institutionsTable},
}

var notificationsTableRequests = &NotificationsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			title		 VARCHAR(256) NOT NULL,
			description  VARCHAR(256) NOT NULL,
			target_id 	 INT 		  NOT NULL,
			target_type  %s 		  NOT NULL
		);
	`, NotificationsTableName, DefaultFieldsQuery, UserRolesTypeName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, NotificationsTableName)),
}

type Notifications struct {
	DefaultFields

	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	// Notification can be sent to a Class using UserRole=Unkown and TargetId=ClassID
	// Notification can be sent to a Roles using UserRole=Role and TargetId=InstitutionID
	TargetId   int      `json:"target_id" binding:"required"`
	TargetType UserRole `json:"target_type" binding:"required"`
}

func (t *NotificationsTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *NotificationsTable) CreateRequest() *transactions.TransactionRequest {
	return t.requests.Create
}
func (t *NotificationsTable) DropRequest() *transactions.TransactionRequest {
	return t.requests.Drop
}
func (t *NotificationsTable) Name() TableName {
	return t.name
}
