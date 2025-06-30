package pim

type PostgrestEnvironment struct {
	POSTGRES_USER          string `env:"POSTGRES_USER,required"`
	POSTGRES_USER_PASSWORD string `env:"POSTGRES_USER_PASSWORD,required"`

	POSTGRES_TEST_USER          string `env:"POSTGRES_TEST_USER,required"`
	POSTGRES_TEST_USER_PASSWORD string `env:"POSTGRES_TEST_USER_PASSWORD,required"`

	POSTGRES_CONTAINER_NAME string `env:"PG_CONTAINER_NAME,required"`
}
