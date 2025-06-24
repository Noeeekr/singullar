package pim

import (
	"database/sql"
)

type PostgresInterfaceManager struct {
	db *sql.DB
}

func NewQueryManager(db *sql.DB) *PostgresInterfaceManager {
	return &PostgresInterfaceManager{
		db: db,
	}
}

func (m *PostgresInterfaceManager) Ping() error {
	return m.db.Ping()
}
func (m *PostgresInterfaceManager) Close() error {
	return m.db.Close()
}
