package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type ClassesTable struct {
	name         TableName
	dependencies *TableDependencies
	requests     *ClassesRequests
}

type ClassesRequests struct {
	Create *transactions.Request
	Drop   *transactions.Request
}

var classesTable *ClassesTable = &ClassesTable{
	name:         ClassesTableName,
	dependencies: classesTableDependencies,
	requests:     classesTableRequests,
}

var classesTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{},
	Tables: []TableMethods{institutionsTable},
}

var classesTableRequests = &ClassesRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			%s
			name VARCHAR(256) NOT NULL,
			segment VARCHAR(256) NOT NULL,
			series VARCHAR(256) NOT NULL,

			institution_id INT NOT NULL,

			FOREIGN KEY (institution_id) REFERENCES %s(id)
		);
	`, ClassesTableName, DefaultFieldsQuery, SerialId, InstitutionsTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, ClassesTableName)),
}

type Classes struct {
	DefaultFields
	ID
	Name    string `json:"name" binding:"required"`
	Segment string `json:"segment" binding:"required"`
	Series  string `json:"series" binding:"required"`

	InstitutionId int `json:"institution_id" binding:"required"`
}

func (t *ClassesTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *ClassesTable) GetCreateRequest() *transactions.Request {
	return t.requests.Create
}

func (t *ClassesTable) GetDropRequest() *transactions.Request {
	return t.requests.Drop
}

func (t *ClassesTable) Name() TableName {
	return t.name
}
