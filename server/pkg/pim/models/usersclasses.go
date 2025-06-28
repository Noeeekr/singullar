package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

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
	Create: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			class_id INT NOT NULL,

			FOREIGN KEY (user_id) REFERENCES %s(id),
			FOREIGN KEY (class_id) REFERENCES %s(id)
		);
	`, usersClassesTableName, usersTableName, classesTableName)),
	Drop: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, usersClassesTableName)),
}

type UsersClasses struct {
	UserId        int
	InstitutionId int
}
