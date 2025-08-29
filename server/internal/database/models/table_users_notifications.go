package models

import "github.com/Noeeekr/borm"

// Used by API to parse a create request
type CreateUsersNotifications struct {
	// Notification can be sent to a Class using TargetRole=Unkown and TargetId=ClassID
	// Notification can be sent to a Roles using TargetRole=Role and TargetId=InstitutionID
	UserRole UserRole `json:"targetRole" binding:"required" borm:"(NAME, user_role) (TYPE, user_roles)"`
	UserId   int      `json:"targetId" binding:"required" borm:"(NAME, user_id) (FOREIGN KEY, users, id)"`
}
type UsersNotifications struct {
	*CreateUsersNotifications

	// The id of the notification is the id of its creator
	NotificationId int `json:"notificationId" binding:"required" borm:"(NAME, notification_id) (FOREIGN KEY, notifications, id)"`
}

var TableUsersNotifications *borm.TableRegistry = EnvironmentDatabase.RegisterTable(UsersNotifications{}).
	Name("users_notifications").
	NeedTables(TableUsers, TableNotificationContents).
	NeedRoles(TypeUserRole)

func NewUsersNotifications(usrRole UserRole, usrId, notificationId int) *UsersNotifications {
	return &UsersNotifications{
		CreateUsersNotifications: &CreateUsersNotifications{
			UserRole: usrRole,
			UserId:   usrId,
		},
		NotificationId: notificationId,
	}
}

// var usersNotificationsTableDependencies *TableDependencies = &TableDependencies{
// 	Types:  []*TypeInfo{UserRolesType},
// 	Tables: []TableMethods{UsersClassesTable, NotificationsTable},
// }

// var usersNotificationsTableRequests = &UsersNotificationsRequests{
// 	Create: transactions.NewRequest(fmt.Sprintf(`
// 		CREATE TABLE IF NOT EXISTS %s (
// 			target_role %s NOT NULL,
// 			user_id INT NOT NULL,
// 			notification_id INT NOT NULL,

// 			CONSTRAINT fk_users_notifications_target_id
// 			FOREIGN KEY (user_id)
// 			REFERENCES %s(id),

// 			CONSTRAINT fk_users_notifications_source_id
// 			FOREIGN KEY (notification_id)
// 			REFERENCES %s(id)
// 			ON DELETE CASCADE
// 		);
// 	`, UsersNotificationsTableName, UserRolesTypeName, usersTableName, NotificationsTableName)),
// 	Drop: transactions.NewRequest(fmt.Sprintf(`
// 		DROP TABLE IF EXISTS %s CASCADE;
// 	`, UsersNotificationsTableName)),
// 	InsertMany: transactions.NewRequest(fmt.Sprintf(`
// 		INSERT INTO %s (user_id, target_role, notification_id)
// 		VALUES %s
// 		RETURNING user_id, target_role, notification_id;
// 	`, UsersNotificationsTableName, placeholder)).AllowValueRepeat(placeholder, 3),
// }
