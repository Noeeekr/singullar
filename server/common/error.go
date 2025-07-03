package common

import "fmt"

type ErrorStatus string

const (
	StatusNotFound       ErrorStatus = "StatusNotFound"
	StatusInternalError  ErrorStatus = "StatusInternalError"
	StatusInvalidRequest ErrorStatus = "StatusInvalidRequest"
	StatusEmpty          ErrorStatus = "StatusEmpty"
)

type Error struct {
	Status      ErrorStatus
	Description string
}

func (e *Error) ParseToError() error {
	return fmt.Errorf("[%s]: %s", e.Status, e.Description)
}
