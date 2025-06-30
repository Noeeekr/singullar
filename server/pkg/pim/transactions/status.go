package transactions

import "reflect"

type ResponseStatus int

const (
	StatusFailedTransactionStart ResponseStatus = iota + 999
	StatusSuccess
	StatusFailedTransactionRollback
	StatusFailedTransactionCommit
	StatusFailedTransactionScan
	StatusFailedTransaction
	StatusUnregisteredMigration
	StatusUnregisteredMethod
	StatusAlreadyExists
	StatusInvalidResponse
	StatusInvalidSyntax
)

func (r ResponseStatus) ToString() string {
	return reflect.TypeOf(r).Name()
}
