package transactions

import (
	"database/sql"
)

type Transaction struct {
	tx *sql.Tx
}

type TransactionManager struct {
	db *sql.DB
}

func New(db *sql.DB) *TransactionManager {
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

// No transaction.
func (m *TransactionManager) Query(query *TransactionRequest) *Response {
	stmt, err := m.db.Prepare(query.Query)
	if err != nil {
		res := NewResponse()
		res.SetDescription(err.Error())
		res.SetStatus(StatusInvalidSyntax)
		return res
	}

	rows, err := stmt.Query(query.Args...)
	if err != nil {
		res := NewResponse()
		res.SetStatus(StatusFailedTransaction)
		res.SetDescription("Failed operation. " + err.Error())
		return res
	}

	return query.ReturnHandler(rows)
}
func (m *TransactionManager) Start() (*Transaction, *Response) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, NewResponse().SetDescription(err.Error()).SetStatus(StatusFailedTransactionStart)
	}
	return &Transaction{tx}, nil
}

// On success response == nil. Doesnt commit
func (t *Transaction) Query(query *TransactionRequest) *Response {
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

	if query.ReturnHandler != nil {
		return t.query(stmt, query.ReturnHandler, query.Args...)
	}

	return t.exec(stmt, query.Args...)
}

// On success response == nil
func (t *Transaction) Commit() *Response {
	if err := t.tx.Commit(); err != nil {
		res := NewResponse()
		res.SetStatus(StatusFailedTransactionCommit)
		res.SetDescription("Transaction Failed: " + err.Error())
		return res
	}

	return nil
}

// On success response == nil
func (t *Transaction) query(stmt *sql.Stmt, handlerFunc RequestReturnHandler, args ...any) *Response {
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

	if res := handlerFunc(rows); res != nil {
		return res
	}

	return nil
}

// On success response == nil
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

	return nil
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
