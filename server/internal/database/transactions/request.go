package transactions

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Noeeekr/singullar/server/common"
)

// Allows creating queries for insert, create, update, select and more
// Allows creating requests from created queries
type Request struct {
	// Query writting related fields
	Query             string
	valuesPlaceholder string
	valuesFieldAmount int

	// Query return related fields
	RowsScanner RequestRowsScanner

	// Query request values related fields
	// If present the Transaction will make the operation in a query and use RowsScanner to scan the rows
	Args              []any
	throwErrorOnFound bool
}

type RequestRowsScanner func(rows *sql.Rows, throwErrorOnFound bool) *common.Response

func NewRequest(query string) *Request {
	return &Request{
		RowsScanner: nil,
		Query:       query,
	}
}

// Allows putting a placeholder instead of values inside the sql query to use a later to define many values. WithArgs automatically handles the parsing and throws error if amount is not sufficient
func (r *Request) AllowValueRepeat(placeholder string, fieldAmount int) *Request {
	r.valuesPlaceholder = placeholder
	r.valuesFieldAmount = fieldAmount
	return r
}

// WithArgs created a copy of the query and inserts the args to be passed to UPDATE | INSERT | DELETE | SELECT queries
func (r Request) WithArgs(args ...any) *Request {
	r.Args = args

	// Value Repeat feature
	if r.valuesFieldAmount == 0 {
		return &r
	}

	return r.setValueFieldSizeToArgsLength(len(args))
}

// Defines a function to handle returned rows. If no function is passed at all then it doesn't query the returned rows.
func (r *Request) WithRowsScanner(fun RequestRowsScanner) *Request {
	r.RowsScanner = fun
	return r
}

func (r *Request) GetValueFieldSize() int {
	return r.valuesFieldAmount
}

// Switch to throw response error on found instead of not found..
func (r *Request) ThrowErrorOnFound() *Request {
	r.throwErrorOnFound = true
	return r
}

func (r *Request) setValueFieldSizeToArgsLength(argAmount int) *Request {
	var placeholders []string

	var index int = 1
	for range argAmount / r.valuesFieldAmount {
		stack := make([]string, r.valuesFieldAmount)
		for i := range r.valuesFieldAmount {
			stack[i] = fmt.Sprintf("$%d", index)
			index++
		}
		placeholders = append(placeholders, "("+strings.Join(stack, ", ")+") ")
	}

	r.Query = strings.Replace(
		r.Query,
		r.valuesPlaceholder,
		strings.Join(placeholders, ", "),
		1,
	)

	return r
}
