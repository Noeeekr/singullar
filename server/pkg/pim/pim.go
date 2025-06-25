package pim

import (
	"database/sql"
)

type PostgresInterfaceManager struct {
	db *sql.DB
}
