package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

func Connect(connString string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", connString)
	if err != nil {
		return db, errors.New("Unable to connect to postgres: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return db, errors.New("Unable to ping postgres. " + err.Error())
	}
	return db, err
}
