package models

import "fmt"

const institutionsTableName TableName = "institutions"

var InstitutionsTable = &TableInfo{
	name:    institutionsTableName,
	Queries: institutionsTableQueries,
	Dependencies: &TableDepencies{
		Types:  []*TypeInfo{},
		Tables: []*TableInfo{},
	},
}

var institutionsTableQueries = &TableQueries{
	Create: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name VARCHAR(256) NOT NULL
		);
	`, institutionsTableName, DefaultFieldsQuery),
	SelectOne: fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, id FROM %s WHERE id = ;
	`, institutionsTableName),
	InsertOne: fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, name) 
		VALUES ($1, $2, $3)
		RETURNING id;
	`, institutionsTableName),
}

type Institutions struct {
	DefaultFields
	Name string `json:"name" binding:"required"`
}
