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

// var roleTypeQueries = &TypeQueries{
// 	Create: transactions.NewRequest(fmt.Sprintf(`
// 			DO $$
// 			BEGIN
// 				IF NOT EXISTS (SELECT * FROM pg_type WHERE typname = '%s') THEN
// 					CREATE TYPE %s AS ENUM ( '%s','%s','%s','%s','%s' );
// 				END IF;
// 			END $$;
// 		`, UserRolesTypeName, UserRolesTypeName,
// 		Admin, Student, Supervisor, Teacher, Unknown,
// 	)),
// 	Drop: transactions.NewRequest(fmt.Sprintf(`
// 		DROP TYPE IF EXISTS %s CASCADE;
// 	`, UserRolesTypeName)),
// }
