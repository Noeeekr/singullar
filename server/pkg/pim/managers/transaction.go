package managers

import (
	"database/sql"
)

type TransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

func (m *TransactionManager) Close() error {
	return m.db.Close()
}
func (m *TransactionManager) Ping() error {
	return m.db.Ping()
}

func (m *TransactionManager) Transaction(query *Query) *Response {
	tx, err := m.db.Begin()
	if err != nil {
		return &Response{
			Description: "Unable to start transaction. " + err.Error(),
			Status:      StatusFailedTransactionStart,
		}
	}

	switch query.Method {
	case CREATE, DROP, INSERT:
		res := m.transaction(tx, query.Query)
		if res.Status != StatusSuccess {
			return res
		}
	default:
		return &Response{
			Description: "Transaction method not registered.",
			Status:      StatusUnregisteredMethod,
		}
	}

	if err := tx.Commit(); err != nil {
		return &Response{
			Status:      StatusFailedTransactionCommit,
			Description: "Transaction Failed: " + err.Error(),
		}
	}

	return &Response{
		Status:      StatusSuccess,
		Description: "Transaction commited successfully",
	}
}

func (m *TransactionManager) Transactions(queries ...*Query) *Response {
	tx, err := m.db.Begin()
	if err != nil {
		return &Response{
			Description: "Unable to start transaction. " + err.Error(),
			Status:      StatusFailedTransactionStart,
		}
	}

	for _, query := range queries {
		switch query.Method {
		case INSERT:
		case DROP:
		case CREATE:
			res := m.transaction(tx, query.Query)
			if res.Status != StatusSuccess {
				return res
			}
		default:
			return &Response{
				Description: "Transaction method not registered.",
				Status:      StatusUnregisteredMethod,
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return &Response{
			Status:      StatusFailedTransactionCommit,
			Description: "Transaction Failed: " + err.Error(),
		}
	}

	return &Response{
		Description: "Transaction finished successfully",
		Status:      StatusSuccess,
	}
}

// Transaction executes a single query in the context of a transaction. Rollbacks the transaction if the query fails. Doesn't commit at the end.
//
// Meant to be used internally to execute multi-query transactions.
func (m *TransactionManager) transaction(tx *sql.Tx, query string) *Response {
	if _, err := tx.Exec(query); err != nil {
		if err := tx.Rollback(); err != nil {
			return &Response{
				Status:      StatusFailedTransactionRollback,
				Description: "Transaction failed. Unable to rollback: " + err.Error(),
			}
		}
		return &Response{
			Status:      StatusFailedTransaction,
			Description: "Tansaction failed. Rollback executed: " + err.Error(),
		}
	}

	return &Response{
		Status:      StatusSuccess,
		Description: "Transaction executed successfully.",
	}
}
