package models

import "fmt"

const usersNotificationsTableName TableName = "users_notifications" // PARTIAL

var UsersNotificationsTable = &TableInfo{
	name:    usersNotificationsTableName,
	Queries: usersNotificationsTableQueries,
	Dependencies: &TableDepencies{
		Types:  []*TypeInfo{},
		Tables: []*TableInfo{UsersTable, NotificationsTable},
	},
}

var usersNotificationsTableQueries = &TableQueries{
	Create: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			notification_id INT NOT NULL,

			FOREIGN KEY (user_id) REFERENCES %s(id),
			FOREIGN KEY (notification_id) REFERENCES %s(id)
		);
	`, usersNotificationsTableName, usersTableName, notificationsTableName),
}

type UsersNotifications struct {
	UserId         int
	NotificationId int
}
