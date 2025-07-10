package transactions

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/common"
)

// Manager creates, starts and commits transactions
type Manager struct {
	currentTransaction *Transaction
	database           *sql.DB
}

func NewManager(db *sql.DB) *Manager {
	return &Manager{
		currentTransaction: nil,
		database:           db,
	}
}

// Start starts a transaction on the manager and returns the transaction.. If another transaction is happening it returns the current transaction.
func (m *Manager) Start() *Transaction {
	if m.currentTransaction != nil {
		return m.currentTransaction
	}

	tx, err := m.database.Begin()
	if err != nil {
		t := NewTransaction(nil)
		t.Response = common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusFailedTransactionStart)
		return t
	}

	m.currentTransaction = NewTransaction(tx)

	return m.currentTransaction
}
func (m *Manager) Commit() (res *common.Response) {
	if m.currentTransaction == nil {
		return common.NewResponse().
			WithDescription("Unable to commit, no transaction in progress.").
			WithStatus(common.StatusFailedTransactionCommit)
	}

	m.currentTransaction = nil

	return m.currentTransaction.Commit().Response
}

func (m *Manager) Query(request *Request) *common.Response {
	stmt, err := m.database.Prepare(request.Query)
	if err != nil {
		return common.NewResponse().
			WithDescription(err.Error()).
			WithStatus(common.StatusInvalidSyntax)
	}

	rows, err := stmt.Query(request.Args...)
	if err != nil {
		return common.NewResponse().
			WithStatus(common.StatusFailedTransaction).
			WithDescription("Failed operation. " + err.Error())
	}

	if request.RowsScanner != nil {
		return request.RowsScanner(rows, request.throwErrorOnFound)
	}
	return nil
}
