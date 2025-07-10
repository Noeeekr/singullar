// Package transactions contains types that abstract the golang database/sql package transaction operations for easy chaining, gracefull errors and operations.
//
//	Request
//
// Allows creating queries and making requests from those queries.
//
//	Transaction
//
// Allows making chained transactions with gracefull error handling.
package transactions

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/common"
)

// Transaction Automatically switches between Query() and Exec() when necessary.
//
// Transaction contains a common.Response that is different than nil if an error happened at any moment.
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

// On success Transaction == nil. Doesnt commit
func (t *Transaction) Query(request *Request) *Transaction {
	if request == nil {
		t.Response = common.NewResponse().
			WithDescription("Invalid query. Empty query.").
			WithStatus(common.StatusInvalidSyntax)
		return t
	}

	stmt, err := t.tx.Prepare(request.Query)
	if err != nil {
		t.Response = common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusInvalidSyntax)
		return t
	}

	if request.RowsScanner != nil {
		return t.query(stmt, request)
	}
	return t.exec(stmt, request.Args...)
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
func (t *Transaction) query(stmt *sql.Stmt, request *Request) *Transaction {
	rows, err := stmt.Query(request.Args...)
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

	if res := request.RowsScanner(rows, request.throwErrorOnFound); t != nil {
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
