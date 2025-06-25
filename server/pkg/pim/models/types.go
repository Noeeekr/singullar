package models

// TypeName marks the name of all types created in database
type TypeName string

// The name of database custom type
const (
	RoleTypeName TypeName = "roles"
)

// Types marks all types created in database
type Types string

type Roles Types

// Enums of type "Roles" in database
const (
	RoleStudent    Roles = "student"
	RoleTeacher    Roles = "teacher"
	RoleSupervisor Roles = "supervisor"
	RoleAdmin      Roles = "admin"
)
