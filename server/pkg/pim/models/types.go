package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

type TypeQueries struct {
	Create *transaction.QueryInfo
	Drop   *transaction.QueryInfo
}

type TypeInfo struct {
	Name    TypeName
	Queries *TypeQueries
}

// TypeName marks the name of all types created in database
type TypeName string

// Names of types present in database
const (
	userRolesTypeName TypeName = "roles"
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

var UserRolesType = &TypeInfo{
	Name:    userRolesTypeName,
	Queries: roleTypeQueries,
}

var roleTypeQueries = &TypeQueries{
	Create: transaction.NewQuery().WithQuery(fmt.Sprintf(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT * FROM pg_type WHERE typname = '%s') THEN
					CREATE TYPE %s AS ENUM ( '%s','%s','%s','%s','%s' );
				END IF;
			END $$;
		`, userRolesTypeName, userRolesTypeName,
		Admin, Student, Supervisor, Teacher, Unknown,
	)),
	Drop: transaction.NewQuery().WithQuery(fmt.Sprintf(`
		DROP TYPE IF EXISTS %s CASCADE;
	`, userRolesTypeName)),
}
