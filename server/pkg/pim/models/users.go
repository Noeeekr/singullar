package models

import "fmt"

const usersTableName TableName = "users"

var usersTableQueries = &TableQueries{
	Create: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
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
	`, usersTableName, DefaultFieldsQuery, userRoleName, institutionsTableName),
	InsertOne: fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, name, email, password, institution_id, role) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING email;
	`, usersTableName),
	SelectOne: fmt.Sprintf(`
		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture FROM %s WHERE email = $1;
	`, usersTableName),
}

var UsersTable = &TableInfo{
	name:    usersTableName,
	Queries: usersTableQueries,
	Dependencies: &TableDepencies{
		Types:  []*TypeInfo{RoleType},
		Tables: []*TableInfo{InstitutionsTable},
	},
}

type CreateUsers struct {
	Name          string   `json:"name" binding:"required,min=2,max=255"`
	Email         string   `json:"email" binding:"required,email"`
	Password      string   `json:"password" binding:"required,min=6"`
	InstitutionId int      `json:"institution_id" binding:"required"`
	Role          UserRole `json:"role" binding:"required"`
}

type Users struct {
	DefaultFields
	CreateUsers

	ProfilePicture string `json:"profile_picture"`
}
