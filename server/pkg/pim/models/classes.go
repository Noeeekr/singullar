package models

import "fmt"

const classesTableName TableName = "classes" // PARTIAL

var ClassesTable = &TableInfo{
	name:    classesTableName,
	Queries: classesTableQueries,
	Dependencies: &TableDepencies{
		Types:  []*TypeInfo{},
		Tables: []*TableInfo{InstitutionsTable},
	},
}

var classesTableQueries = &TableQueries{
	Create: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name VARCHAR(256) NOT NULL,
			segment VARCHAR(256) NOT NULL,
			series VARCHAR(256) NOT NULL,

			institution_id INT NOT NULL,

			FOREIGN KEY (institution_id) REFERENCES %s(id)
		);
	`, classesTableName, DefaultFieldsQuery, institutionsTableName),
}

type Classes struct {
	DefaultFields
	Name    string `json:"name" binding:"required"`
	Segment string `json:"segment" binding:"required"`
	Series  string `json:"series" binding:"required"`

	InstitutionId int `json:"institution_id" binding:"required"`
}
