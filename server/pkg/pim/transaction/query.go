package transaction

import "database/sql"

type QueryInfo struct {
	// Wether or not that query returns rows
	ReturnFunc QueryReturnFunc
	Query      string
	Args       []any
}

type QueryReturnFunc func(rows *sql.Rows) error

func NewQuery() *QueryInfo {
	return &QueryInfo{
		ReturnFunc: nil,
		Query:      "",
		Args:       []any{},
	}
}

// WithArgs inserts the args to be passed to UPDATE | INSERT | DELETE | SELECT queries
func (q *QueryInfo) WithArgs(args ...any) *QueryInfo {
	return &QueryInfo{
		ReturnFunc: q.ReturnFunc,
		Query:      q.Query,
		Args:       args,
	}
}

// WithQuery defines the query to be executed in database. It may be chained with WithArgs() if the query needs dynamic arguments. It also may be chained with WithReturn() to handle the return rows if exist.
func (q *QueryInfo) WithQuery(query string) *QueryInfo {
	return &QueryInfo{
		ReturnFunc: q.ReturnFunc,
		Query:      query,
		Args:       q.Args,
	}
}

// Defines a function to handle returned rows. If no function is passed at all then it doesn't query the returned rows.
func (q *QueryInfo) WithScanFunc(fun QueryReturnFunc) *QueryInfo {
	return &QueryInfo{
		ReturnFunc: fun,
		Query:      q.Query,
		Args:       q.Args,
	}
}
