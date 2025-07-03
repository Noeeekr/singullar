package database

type PostgresEnvironment struct {
	POSTGRES_CONNECTION_STRING string `env:"POSTGRES_CONNECTION_STRING,required"`
}

type PostgresTestEnvironment struct {
	POSTGRES_TEST_CONNECTION_STRING string `env:"POSTGRES_TEST_CONNECTION_STRING,required"`
}
