package managers

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func ParseConnectionString(host, user, passwd string) string {
	return fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		host, user, passwd, user,
	)
}

func Connect(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, errors.New("Unable to connect to postgres: " + err.Error())
	}

	return db, nil
}
