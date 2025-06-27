package models

import "fmt"

const usersClassesTableName TableName = "users_classes" // PARTIAL

var UsersClassesTable = &TableInfo{
	name:    usersClassesTableName,
	Queries: usersClassesTableQueries,
	Dependencies: &TableDepencies{
		Types:  []*TypeInfo{},
		Tables: []*TableInfo{UsersTable, ClassesTable},
	},
}

var usersClassesTableQueries = &TableQueries{
	Create: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			class_id INT NOT NULL,

			FOREIGN KEY (user_id) REFERENCES %s(id),
			FOREIGN KEY (class_id) REFERENCES %s(id)
		);
	`, usersClassesTableName, usersTableName, classesTableName),
}

type UsersClasses struct {
	UserId        int
	InstitutionId int
}
