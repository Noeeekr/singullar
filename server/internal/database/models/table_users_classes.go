package models

import "github.com/Noeeekr/borm"

type UsersClasses struct {
	UserId  int `borm:"(NAME, user_id) (FOREIGN KEY, users, id) (CONSTRAINTS, NOT NULL)"`
	ClassId int `borm:"(NAME, class_id) (FOREING KEY, classes, id) (CONSTRAINTS, NOT NULL)"`
}

var TableUsersClasses *borm.TableRegistry = EnvironmentDatabase.RegisterTable(UsersClasses{}).
	Name("users_classes").
	NeedTables(TableUsers, TableClasses, TableInstitutions)
