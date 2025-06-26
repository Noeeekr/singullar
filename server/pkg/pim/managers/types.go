package managers

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

type ResponseStatus int

const (
	StatusSuccess ResponseStatus = iota + 999
	StatusFailedTransaction
	StatusFailedTransactionStart
	StatusFailedTransactionRollback
	StatusFailedTransactionCommit
	StatusUnregisteredMigration
	StatusUnregisteredMethod
)

type Query struct {
	Method QueryMethod
	Query  string
}

type QueryMethod int

const (
	INSERT QueryMethod = iota
	DROP
	SELECT
	CREATE
	UPDATE
)
