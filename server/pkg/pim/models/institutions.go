package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

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
	Create: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name VARCHAR(256) NOT NULL
		);
	`, institutionsTableName, DefaultFieldsQuery)),
	SelectOne: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, id FROM %s WHERE id = $1;
	`, institutionsTableName)),
	InsertOne: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, name) 
		VALUES ($1, $2, $3)
		RETURNING id;
	`, institutionsTableName)),
	Drop: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, institutionsTableName)),
}

type Institutions struct {
	DefaultFields
	Name string `json:"name" binding:"required"`
}
