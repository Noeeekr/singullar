package models

import (
	"time"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type TableName string

const (
	usersTableName         TableName = "users"
	InstitutionsTableName  TableName = "institutions"
	ClassesTableName       TableName = "classes"       // PARTIAL
	NotificationsTableName TableName = "notifications" // PARTIAL

	UsersClassesTableName       TableName = "users_classes"       // PARTIAL
	UsersNotificationsTableName TableName = "users_notifications" // PARTIAL
)

type Tables struct {
	Users         *UsersTable
	Classes       *ClassesTable
	Institutions  *InstitutionsTable
	Notifications *NotificationsTable

	UsersClasses       *UsersClassesTable
	UsersNotifications *UsersNotificationsTable
}

type TableMethods interface {
	CreateRequest() *transactions.TransactionRequest
	DropRequest() *transactions.TransactionRequest

	CreateRequestDependencies() *TableDependencies
	Name() TableName
}

type TableDependencies struct {
	Types  []*TypeInfo
	Tables []TableMethods
}

type DefaultFields struct {
	Id        int        `json:"id" binding:"required"`
	CreatedAt time.Time  `json:"created_at" binding:"required"`
	UpdatedAt time.Time  `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

const DefaultFieldsQuery = `
	id         SERIAL      PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ,
`

var TablesInfo = Tables{
	Users:              usersTable,
	Institutions:       institutionsTable,
	Notifications:      notificationsTable,
	Classes:            classesTable,
	UsersClasses:       usersClassesTable,
	UsersNotifications: usersNotificationsTable,
}

//
//
// Tables
//	- Users
// 		Implemented:
// 			Migration: DONE
// 			API JSON: DONE
// 			API CREATE: DONE
//  - Institutions
//		Implemented:
//  		Migration:
//  		API JSON:
//			API CREATE:
//	- Notifications
//		Implemented:
//  		Migration:
//  		API JSON:
//			API CREATE:
//  - Classes
//		Implemented:
//  		Migration:
//  		API JSON:
//			API CREATE:
//
// SubTables
//  - UsersNotifications
//		Depencies
//			- Users - Notifications
//
// 	- UsersInstitutions
// 		Depencies
// 			- Institutions - Users
//
//  - UsersClasses
// 		Depencies
// 			- Classes - Users
//
//  Institutions Config
//		(CHECK ACCESS COOKIE) Then find users in UsersInstitutions WHERE userID == user, instID == institution, role == ADMIN access INSTITUION CONFIG PANEL
//	Classes Config
//		(CHECK ACCESS COOKIE) Then find users in classesUsers WHERE userId == user, role == TEACHER
//  Notifications
//  	Same from above
