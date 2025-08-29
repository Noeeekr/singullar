package models

import "github.com/Noeeekr/borm"

// var notificationsTableRequests = &NotificationsRequests{
// 	Create: transactions.NewRequest(fmt.Sprintf(`
// 		CREATE TABLE IF NOT EXISTS %s (
// 			%s
// 			%s
// 			issuer_id    INT		  NOT NULL,
// 			title		 VARCHAR(256) NOT NULL,
// 			description  VARCHAR(256) NOT NULL,

// 			CONSTRAINT fk_notifications
// 			FOREIGN KEY (issuer_id)
// 			REFERENCES %s (id)
// 			ON DELETE CASCADE
// 		);
// 	`, NotificationsTableName, DefaultFieldsQuery, SerialId, usersTableName)),
// 	Drop: transactions.NewRequest(fmt.Sprintf(`
// 		DROP TABLE IF EXISTS %s CASCADE;
// 	`, NotificationsTableName)),
// 	InsertMany: transactions.NewRequest(fmt.Sprintf(`
// 		INSERT INTO %s (created_at, updated_at, title, description, issuer_id)
// 		VALUES %s
// 		RETURNING created_at, updated_at, deleted_at, id, issuer_id, title, description;
// 	`, NotificationsTableName, placeholder)).AllowValueRepeat(placeholder, 5),
// 	DeleteByIssuerId: transactions.NewRequest(fmt.Sprintf(`
// 		DELETE FROM %s WHERE issuer_id = $1
// 	`, NotificationsTableName)),
// 	SelectByTargetId: transactions.NewRequest(fmt.Sprintf(`
// 		SELECT n.created_at, n.updated_at, n.deleted_at, u.id, u.name, n.id, n.title, n.description
// 		FROM %s n
// 		INNER JOIN %s un ON un.notification_id = n.id
// 		INNER JOIN %s u ON u.id = un.user_id
// 		WHERE u.id = $1;
// 	`, NotificationsTableName, UsersNotificationsTableName, usersTableName)),
// }

type CreateNotificationContents struct {
	Title       string `binding:"required" json:"title"`
	Description string `binding:"required" json:"description"`
	IssuerId    int    `binding:"required" json:"issuerId" borm:"(NAME, issuer_id) (CONSTRAINTS, NOT NULL) (FOREIGN KEY, users, id)"`
}

type NotificationContents struct {
	ID
	DefaultFields
	CreateNotificationContents
}

type Notifications struct {
	DefaultFields
	CreateNotificationContents

	TargetId   int
	TargetName string
}

var TableNotificationContents *borm.TableRegistry = EnvironmentDatabase.
	RegisterTable(NotificationContents{}).
	NeedTables(TableUsers).
	Name("notifications")
