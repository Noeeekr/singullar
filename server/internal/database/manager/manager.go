package manager

import "github.com/Noeeekr/borm"

// DatabaseManager organizes and separates database operation logic. For asyncronous implementations a new instance of DatabaseManager must be created for each goroutine.
type DatabaseManager struct {
	*borm.Commiter
}

// Returns an instance of DatabaseManager. It is recommended to check [type DatabaseManager] for further instructions on how to use it.
func New(commiter *borm.Commiter) *DatabaseManager {
	return &DatabaseManager{
		Commiter: commiter,
	}
}

// Instantiate a new operator to handle the transaction
func (ops *DatabaseManager) NewTransactionOperator() (*Operator, error) {
	tx, err := ops.Commiter.StartTx()
	return NewOperator(tx), err
}
