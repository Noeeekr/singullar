package transactions

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Noeeekr/singullar/server/common"
)

type TransactionRequest struct {
	// Wether or not that query returns rows
	ReturnHandler RequestReturnHandler
	Query         string
	Args          []any

	placeholder string
	valueLength int
}

type RequestReturnHandler func(rows *sql.Rows) *common.Response

func NewRequest(query string) *TransactionRequest {
	return &TransactionRequest{
		ReturnHandler: nil,
		Query:         query,
		Args:          []any{},
	}
}

// Allows putting a placeholder instead of values inside the sql query to use a later to define many values. WithArgs automatically handles the parsing and throws error if amount is not sufficient
func (q *TransactionRequest) AllowValueRepeat(placeholder string, valueLength int) *TransactionRequest {
	q.placeholder = placeholder
	q.valueLength = valueLength
	return q
}

// WithArgs inserts the args to be passed to UPDATE | INSERT | DELETE | SELECT queries
func (t *TransactionRequest) WithArgs(args ...any) *TransactionRequest {
	q := *t

	q.Args = args

	// Value Repeat not enabled
	if q.valueLength == 0 {
		return &q
	}

	var values string

	var index int = 1
	for range len(args) / q.valueLength {
		stack := make([]string, q.valueLength)
		for i := range q.valueLength {
			stack[i] = fmt.Sprintf("$%d", index)
			index++
		}
		values += "(" + strings.Join(stack, ", ") + ")"

	}

	q.Query = strings.Replace(q.Query, q.placeholder, values, 1)

	return &q
}

// Defines a function to handle returned rows. If no function is passed at all then it doesn't query the returned rows.
func (q *TransactionRequest) WithScanFunc(fun RequestReturnHandler) *TransactionRequest {
	q.ReturnHandler = fun
	return q
}
