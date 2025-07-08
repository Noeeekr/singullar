// Package transactions contains types that abstract the golang database/sql package transaction operations for easy chaining, gracefull errors and operations.
package transactions

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/common"
)

// Functions that are not part of Transaction will operate without starting transaction.
// Transaction Automatically switches between Query() and Exec() when necessary.
//
// Methods on Transaction created with a nil pointer will commit at the end of operation.
// Methods on Transaction created with an already started transactions won't commit at the end of operation and will execute in the transaction.
type Transaction struct {
	// Response must be nil on success
	Response *common.Response

	// Transaction context, if exists all operations will be done in this tx and won't commit at the end.
	tx *sql.Tx
}

// NewTransaction creates a transaction. If a tx is != nil all operations will be done in its context and won't commit at the end.
func NewTransaction(tx *sql.Tx) *Transaction {
	return &Transaction{
		Response: nil,
		tx:       tx,
	}
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
func (m *TransactionManager) Query(query *TransactionRequest) *common.Response {
	stmt, err := m.db.Prepare(query.Query)
	if err != nil {
		return common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusInvalidSyntax)
	}

	rows, err := stmt.Query(query.Args...)
	if err != nil {
		return common.NewResponse().
			WithStatus(common.StatusFailedTransaction).
			WithDescription("Failed operation. " + err.Error())
	}

	if query.ReturnHandler != nil {
		return query.ReturnHandler(rows)
	}
	return nil
}
func (m *TransactionManager) Start() *Transaction {
	tx, err := m.db.Begin()
	if err != nil {
		t := NewTransaction(nil)
		t.Response = common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusFailedTransactionStart)
		return t
	}

	t := NewTransaction(tx)
	return t
}

// On success Transaction == nil. Doesnt commit
func (t *Transaction) Query(query *TransactionRequest) *Transaction {
	if query == nil {
		t.Response = common.NewResponse().
			WithDescription("Invalid query. Empty query.").
			WithStatus(common.StatusInvalidSyntax)
		return t
	}

	stmt, err := t.tx.Prepare(query.Query)
	if err != nil {
		t.Response = common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusInvalidSyntax)
		return t
	}

	if query.ReturnHandler != nil {

		return t.query(stmt, query.ReturnHandler, query.Args...)
	}

	return t.exec(stmt, query.Args...)
}

func (t *Transaction) Commit() *Transaction {
	if err := t.tx.Commit(); err != nil {
		t.Response = common.NewResponse().
			WithStatus(common.StatusFailedTransactionCommit).
			WithDescription("Transaction Failed: " + err.Error())
		return t
	}

	return t
}

// On success Transaction == nil
func (t *Transaction) query(stmt *sql.Stmt, handlerFunc RequestReturnHandler, args ...any) *Transaction {
	rows, err := stmt.Query(args...)
	if err != nil {
		if err := t.tx.Rollback(); err != nil {
			t.Response = common.NewResponse().
				WithStatus(common.StatusFailedTransaction).
				WithDescription("Transaction failed. Unable to rollback. " + err.Error())
			return t
		}
		t.Response = common.NewResponse().
			WithStatus(common.StatusFailedTransaction).
			WithDescription("Transaction failed. Rollback executed. " + err.Error())
		return t
	}

	if res := handlerFunc(rows); t != nil {
		t.Response = res
		return t
	}

	return nil
}

// On success Transaction == nil
func (t *Transaction) exec(stmt *sql.Stmt, args ...any) *Transaction {
	_, err := stmt.Exec(args...)
	if err != nil {
		if err := t.tx.Rollback(); err != nil {
			t.Response = common.NewResponse().
				WithStatus(common.StatusFailedTransaction).
				WithDescription("Transaction failed. Unable to rollback. " + err.Error())
			return t
		}
		t.Response = common.NewResponse().
			WithStatus(common.StatusFailedTransaction).
			WithDescription("Transaction failed. Rollback executed. " + err.Error())
		return t
	}

	return t
}

/*
	type tx struct {
		tx
	}

	// Handles choosing exec or query under the hood
	// Rollbacks if necessary
	*tx Query() {

	}

	t, tx := StartTransaction
	if t.Status != StatusSuccess {
		return t
	}

	t := tx.Query(query)
	if t.Status != StatusSuccess {
		return t
	}

	... do something with the query tult

	return tx.Commit()
*/
