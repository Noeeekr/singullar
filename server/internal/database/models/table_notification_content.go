package models

import "github.com/Noeeekr/borm"

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
