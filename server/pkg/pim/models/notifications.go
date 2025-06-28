package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

const notificationsTableName TableName = "notifications" // PARTIAL

var NotificationsTable = &TableInfo{
	name:    notificationsTableName,
	Queries: notificationsTableQueries,
	Dependencies: &TableDepencies{
		Types: []*TypeInfo{UserRolesType},
		// Classes and institutions since target id may point to one
		Tables: []*TableInfo{ClassesTable, InstitutionsTable},
	},
}

var notificationsTableQueries = &TableQueries{
	Create: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			title		 VARCHAR(256) NOT NULL,
			description  VARCHAR(256) NOT NULL,
			target_id 	 INT 		  NOT NULL,
			target_type  %s 		  NOT NULL
		);
	`, notificationsTableName, DefaultFieldsQuery, userRolesTypeName)),
	Drop: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, notificationsTableName)),
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
