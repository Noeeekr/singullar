package managers

import "database/sql"

type QueryManager struct {
	db *sql.DB
}

func NewQueryManager(db *sql.DB) *QueryManager {
	return &QueryManager{
		db: db,
	}
}

func (m *QueryManager) Ping() error {
	return m.db.Ping()
}
func (m *QueryManager) Close() error {
	return m.db.Close()
}
