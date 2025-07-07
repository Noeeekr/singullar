package connections

import (
	"database/sql"
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	_ "github.com/lib/pq"
)

func ScanEnvironmentForConnection(connectionEnvironment ConnectionEnvironment) (connection Connection, res *common.Response) {
	switch connectionEnvironment {
	case Production:
		production := &PostgresProductionConnection{}
		res = environment.Scan(production)
		connection = production
	case Development:
		development := &PostgresDevelopmentConnection{}
		res = environment.Scan(development)
		connection = development
	case Postgres:
		postgres := &PostgresConnection{}
		res = environment.Scan(postgres)
		connection = postgres
	default:
		return nil, common.NewResponse().
			WithStatus(common.StatusNotFound).
			WithDescription("Environment not defined: " + string(connectionEnvironment))
	}

	if res != nil {
		return nil, res
	}

	return connection, nil
}

func ParseConnectionString(connection Connection) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		connection.User(), connection.Password(), connection.Host(), connection.Database(),
	)
}
func NewConnectionString(user, password, host, database string) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		user, password, host, database,
	)
}

// Scans the environment variables, uses them to parse the connection string to connect to Postgres.
// Check the package for more info about the required environment variables.
func ConnectWithEnvironment(environment ConnectionEnvironment) (db *sql.DB, err *common.Response) {
	connection, err := ScanEnvironmentForConnection(environment)
	if err != nil {
		return nil, err
	}
	return Connect(ParseConnectionString(connection))
}

func Connect(connString string) (*sql.DB, *common.Response) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return db, common.NewResponse().
			WithStatus(common.StatusInternalError).
			WithDescription("Unable to connect to postgres: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return db, common.NewResponse().
			WithStatus(common.StatusInternalError).
			WithDescription("Unable to ping postgres. " + err.Error())
	}

	return db, nil
}
