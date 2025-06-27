package managers

import (
	"database/sql"
)

type Transaction struct {
	tx *sql.Tx
}

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
func (m *TransactionManager) StartTransation() (*Transaction, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Transaction{tx}, nil
}

func (t *Transaction) Query(query *Query) *TransactionResponse {
	stmt, err := t.tx.Prepare(query.Query)
	if err != nil {
		return &TransactionResponse{
			Response: &Response{
				Description: "Invalid query. " + err.Error(),
				Status:      StatusInvalidSyntax,
			},
		}
	}

	var res *TransactionResponse
	if query.Returns {
		res = t.query(stmt, query.Args...)
	} else {
		res = t.exec(stmt, query.Args...)
	}
	return res
}

func (t *Transaction) query(stmt *sql.Stmt, args ...any) *TransactionResponse {
	res, err := stmt.Query(args...)
	if err != nil {
		if err := t.tx.Rollback(); err != nil {
			return &TransactionResponse{
				Response: &Response{
					Description: "Transaction failed. Unable to rollback. " + err.Error(),
					Status:      StatusFailedTransaction,
				},
			}
		}
		return &TransactionResponse{
			Response: &Response{
				Description: "Transaction failed. Rollback executed. " + err.Error(),
				Status:      StatusFailedTransaction,
			},
		}
	}
	return &TransactionResponse{
		Response: &Response{
			Description: "Transaction query executed successfully",
			Status:      StatusSuccess,
		},
		Rows: res,
	}
}

func (t *Transaction) exec(stmt *sql.Stmt, args ...any) *TransactionResponse {
	_, err := stmt.Exec(args...)
	if err != nil {
		if err := t.tx.Rollback(); err != nil {
			return &TransactionResponse{
				Response: &Response{
					Description: "Transaction failed. Unable to rollback. " + err.Error(),
					Status:      StatusFailedTransaction,
				},
			}
		}
		return &TransactionResponse{
			Response: &Response{
				Description: "Transaction failed. Rollback executed. " + err.Error(),
				Status:      StatusFailedTransaction,
			},
		}
	}
	return &TransactionResponse{
		Response: &Response{
			Description: "Transaction query executed successfully",
			Status:      StatusSuccess,
		},
	}
}

func (t *Transaction) Commit() *TransactionResponse {
	if err := t.tx.Commit(); err != nil {
		return &TransactionResponse{
			Response: &Response{
				Status:      StatusFailedTransactionCommit,
				Description: "Transaction Failed: " + err.Error(),
			},
		}
	}
	return &TransactionResponse{
		Response: &Response{
			Status:      StatusSuccess,
			Description: "Transaction commited successfully: ",
		},
	}
}

/*
	type tx struct {
		tx
	}

	// Handles choosing exec or query under the hood
	// Rollbacks if necessary
	*tx Query() {

	}

	res, tx := StartTransaction
	if res.Status != StatusSuccess {
		return res
	}

	res := tx.Query(query)
	if res.Status != StatusSuccess {
		return res
	}

	... do something with the query result

	return tx.Commit()
*/
