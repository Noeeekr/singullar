package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

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
	Create: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			notification_id INT NOT NULL,

			FOREIGN KEY (user_id) REFERENCES %s(id),
			FOREIGN KEY (notification_id) REFERENCES %s(id)
		);
	`, usersNotificationsTableName, usersTableName, notificationsTableName)),
	Drop: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, usersNotificationsTableName)),
}

type UsersNotifications struct {
	UserId         int
	NotificationId int
}
