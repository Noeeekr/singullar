package managers

import "database/sql"

type PostgrestManagerEnvironment struct {
	POSTGRES_USER          string `env:"POSTGRES_USER,required"`
	POSTGRES_USER_PASSWORD string `env:"POSTGRES_USER_PASSWORD,required"`

	POSTGRES_TEST_USER          string `env:"POSTGRES_TEST_USER,required"`
	POSTGRES_TEST_USER_PASSWORD string `env:"POSTGRES_TEST_USER_PASSWORD,required"`

	POSTGRES_CONTAINER_NAME string `env:"PG_CONTAINER_NAME,required"`
}

type Response struct {
	Description string
	Status      ResponseStatus
}

type TransactionResponse struct {
	*Response
	Rows *sql.Rows
}

type ResponseStatus int

const (
	StatusFailedTransactionStart ResponseStatus = iota + 999
	StatusSuccess
	StatusAlreadyExists
	StatusFailedTransactionRollback
	StatusFailedTransactionCommit
	StatusFailedTransaction
	StatusInvalidSyntax
	StatusUnregisteredMigration
	StatusUnregisteredMethod
)

type Query struct {
	// Wether or not that query returns rows
	Returns bool
	Query   string
	Args    []any
}
