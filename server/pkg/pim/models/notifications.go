package models

import "fmt"

const notificationsTableName TableName = "notifications" // PARTIAL

var NotificationsTable = &TableInfo{
	name:    notificationsTableName,
	Queries: notificationsTableQueries,
	Dependencies: &TableDepencies{
		Types: []*TypeInfo{RoleType},
		// Classes and institutions since target id may point to one
		Tables: []*TableInfo{ClassesTable, InstitutionsTable},
	},
}

var notificationsTableQueries = &TableQueries{
	Create: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			title		 VARCHAR(256) NOT NULL,
			description  VARCHAR(256) NOT NULL,
			target_id 	 INT 		  NOT NULL,
			target_type  %s 		  NOT NULL
		);
	`, notificationsTableName, DefaultFieldsQuery, userRoleName),
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
