package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type InstitutionsTable struct {
	name         TableName
	Requests     *InstitutionsRequests
	dependencies *TableDependencies
}

type InstitutionsRequests struct {
	Create          *transactions.TransactionRequest
	SelectOneByName *transactions.TransactionRequest
	SelectOneById   *transactions.TransactionRequest
	DeleteOneByName *transactions.TransactionRequest
	DeleteOneById   *transactions.TransactionRequest
	InsertOne       *transactions.TransactionRequest
	Drop            *transactions.TransactionRequest
}

var institutionsTable = &InstitutionsTable{
	name:         InstitutionsTableName,
	dependencies: institutionsTableDependencies,
	Requests:     institutionsTableRequests,
}

var institutionsTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{},
	Tables: []TableMethods{},
}

var institutionsTableRequests = &InstitutionsRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name VARCHAR(256) NOT NULL
		);
	`, InstitutionsTableName, DefaultFieldsQuery)),
	SelectOneById: transactions.NewRequest(fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, id FROM %s WHERE id = $1;
	`, InstitutionsTableName)),
	SelectOneByName: transactions.NewRequest(fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, id FROM %s WHERE name = $1;
	`, InstitutionsTableName)),
	InsertOne: transactions.NewRequest(fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, name) 
		VALUES ($1, $2, $3)
		RETURNING id;
	`, InstitutionsTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, InstitutionsTableName)),
	DeleteOneById: transactions.NewRequest(fmt.Sprintf(`
		DELETE FROM %s WHERE id = $1;
	`, InstitutionsTableName)),
	DeleteOneByName: transactions.NewRequest(fmt.Sprintf(`
		DELETE FROM %s WHERE name = $1;
	`, InstitutionsTableName)),
}

type Institutions struct {
	DefaultFields
	Name string `json:"name" binding:"required"`
}

func (t *InstitutionsTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *InstitutionsTable) CreateRequest() *transactions.TransactionRequest {
	return t.Requests.Create
}

func (t *InstitutionsTable) Name() TableName {
	return t.name
}

func (t *InstitutionsTable) DropRequest() *transactions.TransactionRequest {
	return t.Requests.Drop
}
