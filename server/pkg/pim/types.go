package pim

type PostgrestManagerEnvironment struct {
	POSTGRES_USER           string `env:"POSTGRES_USER,required"`
	POSTGRES_PASSWORD       string `env:"POSTGRES_PASSWORD,required"`
	POSTGRES_CONTAINER_NAME string `env:"PG_CONTAINER_NAME,required"`
}
