package connections

var Development ConnectionEnvironment = "development"

type PostgresDevelopmentConnection struct {
	HOST     string `env:"POSTGRES_HOST,required"`
	USER     string `env:"SINGULLAR_TEST_USER,required"`
	PASSWORD string `env:"SINGULLAR_TEST_PASSWORD,required"`
	DB       string `env:"SINGULLAR_TEST_DB,required"`
}

func (e *PostgresDevelopmentConnection) User() string {
	return e.USER
}
func (e *PostgresDevelopmentConnection) Password() string {
	return e.PASSWORD
}
func (e *PostgresDevelopmentConnection) Database() string {
	return e.DB
}
func (e *PostgresDevelopmentConnection) Host() string {
	return e.HOST
}
