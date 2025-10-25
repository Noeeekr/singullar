package models

import "github.com/Noeeekr/borm"

type UserRole string

const UserRoleName string = "user_roles"

// Values of type UserRole
const (
	UNKNOWN    UserRole = "unknown"
	STUDENT    UserRole = "student"
	TEACHER    UserRole = "teacher"
	SUPERVISOR UserRole = "supervisor"
	ADMIN      UserRole = "admin"
)

var TypeUserRole *borm.Enum = EnvironmentDatabase.RegisterEnum(UserRoleName, string(UNKNOWN), string(STUDENT), string(TEACHER), string(SUPERVISOR), string(ADMIN))
