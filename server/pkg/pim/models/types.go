package models

import "fmt"

type TypeQueries struct {
	create string
}

func (t *TypeQueries) Create() string {
	return t.create
}

type TypeInfo struct {
	name    TypeName
	Queries *TypeQueries
}

func (t *TypeInfo) Name() TypeName {
	return t.name
}

// TypeName marks the name of all types created in database
type TypeName string

// Names of types present in database
const (
	userRoleName TypeName = "roles"
)

type UserRole string

// Values of type UserRole
const (
	Unknown    UserRole = "unknown"
	Student    UserRole = "student"
	Teacher    UserRole = "teacher"
	Supervisor UserRole = "supervisor"
	Admin      UserRole = "admin"
)

var RoleType = &TypeInfo{
	name:    userRoleName,
	Queries: roleTypeQueries,
}

var roleTypeQueries = &TypeQueries{
	create: fmt.Sprintf(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT * FROM pg_type WHERE typname = '%s') THEN
					CREATE TYPE %s AS ENUM ( '%s','%s','%s','%s','%s' );
				END IF;
			END $$;
		`, userRoleName, userRoleName,
		Admin, Student, Supervisor, Teacher, Unknown,
	),
}
