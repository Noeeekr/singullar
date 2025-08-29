package models

import "github.com/Noeeekr/borm"

type UsersClasses struct {
	UserId  int `borm:"(NAME, user_id) (FOREIGN KEY, users, id) (CONSTRAINTS, NOT NULL)"`
	ClassId int `borm:"(NAME, class_id) (FOREING KEY, classes, id) (CONSTRAINTS, NOT NULL)"`
}

var TableUsersClasses *borm.TableRegistry = EnvironmentDatabase.RegisterTable(UsersClasses{}).
	Name("users_classes").
	NeedTables(TableUsers, TableClasses, TableInstitutions)
	// var usersClassesDependencies *TableDependencies = &TableDependencies{
	// 	Types:  []*TypeInfo{UserRolesType},
	// 	Tables: []TableMethods{UsersTable, ClassesTable},
	// }

	// var usersClassesTableRequests = &UsersClassesRequests{
	// 	Create: transactions.NewRequest(fmt.Sprintf(`
	// 		CREATE TABLE IF NOT EXISTS %s (
	// 			user_id INT NOT NULL,
	// 			class_id INT NOT NULL,

	// 			FOREIGN KEY (user_id) REFERENCES %s(id),
	// 			FOREIGN KEY (class_id) REFERENCES %s(id)
	// 		);
	// 	`, UsersClassesTableName, usersTableName, ClassesTableName)),
	// 	Drop: transactions.NewRequest(fmt.Sprintf(`
	// 		DROP TABLE IF EXISTS %s CASCADE;
	// 	`, UsersClassesTableName)),
	// }
