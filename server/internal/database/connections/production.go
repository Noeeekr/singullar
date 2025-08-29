package connections

type PostgresProductionConnection struct {
	HOST     string `env:"POSTGRES_HOST,required"`
	USER     string `env:"SINGULLAR_USER,required"`
	PASSWORD string `env:"SINGULLAR_PASSWORD,required"`
	DB       string `env:"SINGULLAR_DB,required"`
}

func (e *PostgresProductionConnection) User() string {
	return e.USER
}
func (e *PostgresProductionConnection) Password() string {
	return e.PASSWORD
}
func (e *PostgresProductionConnection) Database() string {
	return e.DB
}
func (e *PostgresProductionConnection) Host() string {
	return e.HOST
}
