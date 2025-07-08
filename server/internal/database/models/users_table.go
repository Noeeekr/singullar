package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type CreateUsers struct {
	Name          string   `json:"name" binding:"required,min=2,max=255"`
	Email         string   `json:"email" binding:"required,email"`
	Password      string   `json:"password" binding:"required,min=6"`
	InstitutionId int      `json:"institution_id" binding:"required"`
	Role          UserRole `json:"role" binding:"required"`
}

func CreateUser(name, email, password string, institutionId int, role UserRole) *CreateUsers {
	return &CreateUsers{
		Name:          name,
		Email:         email,
		Password:      password,
		InstitutionId: institutionId,
		Role:          role,
	}
}

type Users struct {
	DefaultFields
	ID
	CreateUsers

	ProfilePicture string `json:"profile_picture"`
}

type UsersRequests struct {
	Create           *transactions.TransactionRequest
	InsertMany       *transactions.TransactionRequest
	DeleteOneByEmail *transactions.TransactionRequest
	DeleteOneById    *transactions.TransactionRequest
	SelectOneByEmail *transactions.TransactionRequest
	SelectOneById    *transactions.TransactionRequest
	Drop             *transactions.TransactionRequest

	SelectManyByInstitutionId *transactions.TransactionRequest
}

type UsersTable struct {
	TableMethods
	name         TableName
	Requests     *UsersRequests
	dependencies *TableDependencies
}

var usersTable = &UsersTable{
	name:         usersTableName,
	dependencies: usersTableDependencies,
	Requests:     usersTableRequests,
}

var usersTableRequests = &UsersRequests{
	Create: transactions.NewRequest(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			%s
			name             VARCHAR(256)   NOT NULL,
			email            VARCHAR(256)   NOT NULL UNIQUE,
			password 	     VARCHAR(256)   NOT NULL,
			institution_id   INT      	    NOT NULL,
			profile_picture  VARCHAR(256)   DEFAULT 'userprofilepicture.jpg',
			role             %s             NOT NULL,

			CONSTRAINT fk_institutions 
				FOREIGN KEY (institution_id) 
				REFERENCES %s(id)
				ON DELETE CASCADE
		);
	`, usersTableName, DefaultFieldsQuery, SerialId, UserRolesTypeName, InstitutionsTableName)),
	InsertMany: transactions.NewRequest(fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, name, email, password, institution_id, role)
		VALUES %s
		RETURNING created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture;
	`, usersTableName, placeholder)).AllowValueRepeat(placeholder, 7),
	SelectOneByEmail: transactions.NewRequest(fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture 
		FROM %s 
		WHERE email = $1;
	`, usersTableName)),
	SelectOneById: transactions.NewRequest(fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture 
		FROM %s 
		WHERE id = $1;
	`, usersTableName)),
	SelectManyByInstitutionId: transactions.NewRequest(fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture
		FROM %s
		WHERE institution_id = $1;
	`, usersTableName)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TABLE IF EXISTS %s CASCADE;
	`, usersTableName)),
	DeleteOneByEmail: transactions.NewRequest(fmt.Sprintf(`
		DELETE FROM %s WHERE email = $1;
	`, usersTableName)),
	DeleteOneById: transactions.NewRequest(fmt.Sprintf(`
		DELETE FROM %s WHERE id = $1;
	`, usersTableName)),
}

var usersTableDependencies *TableDependencies = &TableDependencies{
	Types:  []*TypeInfo{UserRolesType},
	Tables: []TableMethods{institutionsTable},
}

func (t *UsersTable) CreateRequestDependencies() *TableDependencies {
	return t.dependencies
}

func (t *UsersTable) CreateRequest() *transactions.TransactionRequest {
	return t.Requests.Create
}

func (t *UsersTable) Name() TableName {
	return t.name
}

func (t *UsersTable) DropRequest() *transactions.TransactionRequest {
	return t.Requests.Drop
}
