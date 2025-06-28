package transaction

import "database/sql"

type Response struct {
	Description string
	Status      ResponseStatus
}

func NewResponse() *Response {
	return &Response{
		Description: "Empty response.",
		Status:      StatusEmptyResponse,
	}
}
func (r *Response) SetDescription(description string) *Response {
	r.Description = description
	return r
}
func (r *Response) SetStatus(status ResponseStatus) *Response {
	r.Status = status
	return r
}

type TransactionResponse struct {
	*Response
	Rows *sql.Rows
}

func NewTransactionResponse() *TransactionResponse {
	return &TransactionResponse{
		Response: &Response{},
	}
}

func (r *TransactionResponse) SetRows(rows *sql.Rows) *TransactionResponse {
	r.Rows = rows
	return r
}
