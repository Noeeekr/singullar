package connections

import "github.com/Noeeekr/singullar/server/common/environment"

var POSTGRES environment.EnvironmentMode = "postgres"

type PostgresConnection struct {
	HOST     string `env:"POSTGRES_HOST,required"`
	USER     string `env:"POSTGRES_USER,required"`
	PASSWORD string `env:"POSTGRES_PASSWORD,required"`
	DB       string `env:"POSTGRES_DB,required"`
}

func (e *PostgresConnection) User() string {
	return e.USER
}
func (e *PostgresConnection) Password() string {
	return e.PASSWORD
}
func (e *PostgresConnection) Database() string {
	return e.DB
}
func (e *PostgresConnection) Host() string {
	return e.HOST
}
