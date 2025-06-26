package models

// TypeName marks the name of all types created in database
type TypeName string

// Names of types present in database
const (
	UserRoleName TypeName = "roles"
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
