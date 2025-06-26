package models

import "fmt"

type TableInfo struct {
	Dependencies *TableDepencies
	Query        string
	Name         TableName
}

type TableDepencies struct {
	Types  []TypeName
	Tables []TableName
}

type TableName string

const (
	UsersTableName TableName = "users"
)

// TABLE DEFAUT FIELDS
// Should be put appended before other fields
var defaultFields = `
	ID        INT         PRIMARY KEY,
	CreatedAt TIMESTAMPTZ NOT NULL,
	UpdatedAt TIMESTAMPTZ NOT NULL,
	DeletedAt TIMESTAMPTZ NOT NULL,
`

// TABLE USERS
var UsersTable = TableInfo{
	Name: UsersTableName,
	Query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS users (
			%s
			Name     VARCHAR(256)   NOT NULL,
			Email    VARCHAR(256)   NOT NULL UNIQUE,
			Password VARCHAR(256)   NOT NULL,
			Role     %s             NOT NULL,
		)
	`, defaultFields, UserRoleName),
	Dependencies: &TableDepencies{
		Types:  []TypeName{UserRoleName},
		Tables: []TableName{},
	},
}

type Users struct {
	Name     string   `json:"name" binding:"required,min=2,max=255"`
	Email    string   `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required,min=2"`
	Role     UserRole `json:"role" binding:"required"`
}
