package models

import "fmt"

type TypeInfo struct {
	name  TypeName
	query string
}

func (t *TypeInfo) Name() TypeName {
	return t.name
}
func (t *TypeInfo) Query() string {
	return t.query
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
	name: userRoleName,
	query: fmt.Sprintf(`
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
