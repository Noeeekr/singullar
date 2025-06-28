package models

import (
	"time"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

type TableName string

type TableQueries struct {
	Create    *transaction.QueryInfo
	InsertOne *transaction.QueryInfo
	SelectOne *transaction.QueryInfo
	Drop      *transaction.QueryInfo
}

type TableDepencies struct {
	Types  []*TypeInfo
	Tables []*TableInfo
}

type TableInfo struct {
	Dependencies *TableDepencies
	Queries      *TableQueries
	name         TableName
}

func (t *TableInfo) Name() TableName {
	return t.name
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
