package transactions

import (
	"fmt"
)

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

func (r *Response) ParseToError() error {
	return fmt.Errorf("[%s]: %s", r.Status, r.Description)
}

func (r *Response) SetDescription(description string) *Response {
	r.Description = description
	return r
}
func (r *Response) SetStatus(status ResponseStatus) *Response {
	r.Status = status
	return r
}
