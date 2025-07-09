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
	Create          *transactions.Request
	SelectOneByName *transactions.Request
	SelectOneById   *transactions.Request
	DeleteOneByName *transactions.Request
	DeleteOneById   *transactions.Request
	InsertOne       *transactions.Request
	InsertMany      *transactions.Request
	Drop            *transactions.Request
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
			%s
			name VARCHAR(256) NOT NULL
		);
	`, InstitutionsTableName, DefaultFieldsQuery, SerialId)),
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

type CreateInstitution struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Institutions struct {
	DefaultFields
	ID

	Name string `json:"name" binding:"required"`
}

func (t *InstitutionsTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *InstitutionsTable) GetCreateRequest() *transactions.Request {
	return t.Requests.Create
}

func (t *InstitutionsTable) Name() TableName {
	return t.name
}

func (t *InstitutionsTable) GetDropRequest() *transactions.Request {
	return t.Requests.Drop
}
