package models

import (
	"time"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/util/environment"
)

type ID struct {
	Id int `json:"id" binding:"min=0" borm:"(CONSTRAINTS, PRIMARY KEY) (TYPE, SERIAL)"`
}

type DefaultFields struct {
	CreatedAt *time.Time `json:"created_at" binding:"required" borm:"(NAME, created_at) (CONSTRAINTS, NOT NULL)"`
	UpdatedAt *time.Time `json:"updated_at" binding:"required" borm:"(NAME, updated_at) (CONSTRAINTS, NOT NULL)"`
	DeletedAt *time.Time `json:"deleted_at" borm:"(NAME, deleted_at)"`
}

var conn, _ = connections.ScanEnvironmentForConnection(environment.Settings().ApplicationMode())
var EnvironmentDatabase = borm.RegisterDatabase(conn.Database(), conn.Host(), borm.RegisterUser(conn.User(), conn.Password()))

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
