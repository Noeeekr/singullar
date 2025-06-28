package query

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

type QueryManager struct {
	tx *transaction.Manager
}

func NewQueryManager(db *sql.DB) *QueryManager {
	return &QueryManager{
		tx: transaction.New(db),
	}
}

// Insert doesn't check if it exists before inserting
func (m *QueryManager) Insert(query *transaction.QueryInfo) *transaction.Response {
	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse().
			SetDescription(err.Error()).
			SetStatus(transaction.StatusFailedTransactionStart)
		return res
	}

	res := tx.Query(query)
	if res.Status != transaction.StatusSuccess {
		return res
	}

	return tx.Commit()
}

func (m *QueryManager) Select(query *transaction.QueryInfo) *transaction.Response {
	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse().
			SetDescription(err.Error()).
			SetStatus(transaction.StatusFailedTransactionStart)
		return res
	}

	res := tx.Query(query)
	if res.Status != transaction.StatusSuccess {
		return res
	}

	return tx.Commit()
}

func (m *QueryManager) Delete(query *transaction.QueryInfo) *transaction.Response {
	tx, err := m.tx.Start()
	if err != nil {
		res := transaction.NewResponse().
			SetStatus(transaction.StatusFailedTransactionStart).
			SetDescription(err.Error())
		return res
	}

	res := tx.Query(query)
	if res.Status != transaction.StatusSuccess {
		return res
	}

	return tx.Commit()
}

/*

Delete User <- Email
Delete Inst <- ID


*/
