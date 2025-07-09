package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type UsersClasses struct {
	UserId        int
	InstitutionId int
}

type UsersClassesTable struct {
	TableMethods
	name         TableName
	dependencies *TableDependencies
	Requests     *UsersClassesRequests
}

type UsersClassesRequests struct {
	Create *transactions.Request
	Drop   *transactions.Request
}

var usersClassesTable *UsersClassesTable = &UsersClassesTable{
	name:         UsersClassesTableName,
	dependencies: usersClassesDependencies,
	Requests:     usersClassesTableRequests,
}

var usersClassesDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{usersTable, classesTable},
}

var usersClassesTableRequests = &UsersClassesRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			class_id INT NOT NULL,

			FOREIGN KEY (user_id) REFERENCES %s(id),
			FOREIGN KEY (class_id) REFERENCES %s(id)
		);
	`, UsersClassesTableName, usersTableName, ClassesTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, UsersClassesTableName)),
}

func (t *UsersClassesTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *UsersClassesTable) GetCreateRequest() *transactions.Request {
	return t.Requests.Create
}

func (t *UsersClassesTable) Name() TableName {
	return t.name
}

func (t *UsersClassesTable) GetDropRequest() *transactions.Request {
	return t.Requests.Drop
}
