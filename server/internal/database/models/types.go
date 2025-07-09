package models

import (
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type TypeQueries struct {
	Create *transactions.Request
	Drop   *transactions.Request
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
