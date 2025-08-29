package common

import "fmt"

type ResponseStatus string

const (
	StatusInternalError             ResponseStatus = "Status internal error"
	StatusInvalidRequest            ResponseStatus = "Status invalid request"
	StatusEmpty                     ResponseStatus = "Status empty"
	StatusFailedTransactionRollback ResponseStatus = "Status failed transaction rollback"
	StatusFailedTransactionStart    ResponseStatus = "Status failed transaction start"
	StatusFailedTransactionCommit   ResponseStatus = "Status failed transaction commit"
	StatusFailedTransactionScan     ResponseStatus = "Status failed transaction scan"
	StatusFailedTransaction         ResponseStatus = "Status failed transaction"
	StatusUnregisteredMigration     ResponseStatus = "Status unregistered migration"
	StatusUnregisteredMethod        ResponseStatus = "Status unregistered method"
	StatusFound                     ResponseStatus = "Status found"
	StatusNotEqual                  ResponseStatus = "Status not equal"
	StatusNotFound                  ResponseStatus = "Status not found"
	StatusInvalidResponse           ResponseStatus = "Status invalid response"
	StatusInvalidSyntax             ResponseStatus = "Status invalid syntax"
)

// Must be nil on success
type Response struct {
	Status      ResponseStatus
	Description string
}

func NewResponse() *Response {
	return &Response{
		Status:      StatusEmpty,
		Description: "Empty response.",
	}
}
func (r *Response) String() string {
	return fmt.Sprintf("[%s]: %s", r.Status, r.Description)
}
func (r *Response) ParseToError() error {
	return fmt.Errorf("[%s]: %s", r.Status, r.Description)
}
func (r *Response) WithDescription(description string) *Response {
	r.Description = description
	return r
}
func (r *Response) WithStatus(status ResponseStatus) *Response {
	r.Status = status
	return r
}
