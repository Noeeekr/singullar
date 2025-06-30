package transactions

import "database/sql"

type Response struct {
	Description string
	Status      ResponseStatus
}

func NewResponse() *Response {
	return &Response{
		Description: "Empty response.",
		Status:      StatusInvalidResponse,
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
