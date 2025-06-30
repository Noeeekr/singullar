package transactions

import (
	"database/sql"
)

type TransactionRequest struct {
	// Wether or not that query returns rows
	ReturnHandler RequestReturnHandler
	Query         string
	Args          []any
}

type RequestReturnHandler func(rows *sql.Rows) *Response

func NewRequest(query string) *TransactionRequest {
	return &TransactionRequest{
		ReturnHandler: nil,
		Query:         query,
		Args:          []any{},
	}
}

// WithArgs inserts the args to be passed to UPDATE | INSERT | DELETE | SELECT queries
func (q *TransactionRequest) WithArgs(args ...any) *TransactionRequest {
	return &TransactionRequest{
		ReturnHandler: q.ReturnHandler,
		Query:         q.Query,
		Args:          args,
	}
}

// Defines a function to handle returned rows. If no function is passed at all then it doesn't query the returned rows.
func (q *TransactionRequest) WithScanFunc(fun RequestReturnHandler) *TransactionRequest {
	return &TransactionRequest{
		ReturnHandler: fun,
		Query:         q.Query,
		Args:          q.Args,
	}
}
