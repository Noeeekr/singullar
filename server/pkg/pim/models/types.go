package models

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/pkg/pim/transactions"
)

type TypeQueries struct {
	Create *transactions.TransactionRequest
	Drop   *transactions.TransactionRequest
}

type TypeInfo struct {
	Name    TypeName
	Queries *TypeQueries
}

// TypeName marks the name of all types created in database
type TypeName string

// Names of types present in database
const (
	UserRolesTypeName TypeName = "roles"
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
	Name:    UserRolesTypeName,
	Queries: roleTypeQueries,
}

var roleTypeQueries = &TypeQueries{
	Create: transactions.NewRequest(fmt.Sprintf(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT * FROM pg_type WHERE typname = '%s') THEN
					CREATE TYPE %s AS ENUM ( '%s','%s','%s','%s','%s' );
				END IF;
			END $$;
		`, UserRolesTypeName, UserRolesTypeName,
		Admin, Student, Supervisor, Teacher, Unknown,
	)),
	Drop: transactions.NewRequest(fmt.Sprintf(`
		DROP TYPE IF EXISTS %s CASCADE;
	`, UserRolesTypeName)),
}
