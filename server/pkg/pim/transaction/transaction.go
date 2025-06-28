package transaction

import (
	"database/sql"
)

type Transaction struct {
	tx *sql.Tx
}

type Manager struct {
	db *sql.DB
}

func New(db *sql.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (m *Manager) Close() error {
	return m.db.Close()
}
func (m *Manager) Ping() error {
	return m.db.Ping()
}
func (m *Manager) Start() (*Transaction, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	return &Transaction{tx}, nil
}

func (t *Transaction) Query(query *QueryInfo) *Response {
	if query == nil {
		res := NewResponse()
		res.SetDescription("Invalid query. Empty query.")
		res.SetStatus(StatusInvalidSyntax)
		return res
	}

	stmt, err := t.tx.Prepare(query.Query)
	if err != nil {
		res := NewResponse()
		res.SetDescription(err.Error())
		res.SetStatus(StatusInvalidSyntax)
		return res
	}

	if query.ReturnFunc != nil {
		return t.query(stmt, query.ReturnFunc, query.Args...)
	}
	return t.exec(stmt, query.Args...)
}

func (t *Transaction) Commit() *Response {
	if err := t.tx.Commit(); err != nil {
		res := NewResponse()
		res.SetStatus(StatusFailedTransactionCommit)
		res.SetDescription("Transaction Failed: " + err.Error())
		return res
	}
	res := NewResponse()
	res.SetDescription("Transaction commited successfully: ")
	res.SetStatus(StatusSuccess)
	return res
}

func (t *Transaction) query(stmt *sql.Stmt, handlerFunc QueryReturnFunc, args ...any) *Response {
	rows, err := stmt.Query(args...)
	if err != nil {
		if err := t.tx.Rollback(); err != nil {
			res := NewResponse()
			res.SetStatus(StatusFailedTransaction)
			res.SetDescription("Transaction failed. Unable to rollback. " + err.Error())
			return res
		}
		res := NewResponse()
		res.SetStatus(StatusFailedTransaction)
		res.SetDescription("Transaction failed. Rollback executed. " + err.Error())
		return res
	}

	if err := handlerFunc(rows); err != nil {
		res := NewResponse().
			SetStatus(StatusFailedTransactionScan).
			SetDescription(err.Error())
		return res
	}

	res := NewResponse()
	res.SetStatus(StatusSuccess)
	res.SetDescription("Transaction query executed successfully")
	return res
}

func (t *Transaction) exec(stmt *sql.Stmt, args ...any) *Response {
	_, err := stmt.Exec(args...)
	if err != nil {
		if err := t.tx.Rollback(); err != nil {
			res := NewResponse()
			res.SetStatus(StatusFailedTransaction)
			res.SetDescription("Transaction failed. Unable to rollback. " + err.Error())
			return res
		}
		res := NewResponse()
		res.SetStatus(StatusFailedTransaction)
		res.SetDescription("Transaction failed. Rollback executed. " + err.Error())
		return res
	}
	res := NewResponse()
	res.SetDescription("Transaction query executed successfully")
	res.SetStatus(StatusSuccess)
	return res
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
