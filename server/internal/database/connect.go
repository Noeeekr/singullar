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

/*
	I'll end up building GORM myself, * is leading to this moment

	type TableManager struct {
		insert
		update
		select
		delete

		drop
		create
	}

	func (t *TableManager) Insert(name TableName) Fields(fields string...) VALUES(values[]) Where(expression string) {
		returns a TableStatement or TableQuery or string && response

		INSERT INTO name + (fields... separated by ,) + (VALUES[0]), (VALUES[1]), (VALUES[2]) + expression
	}

	tables.Insert(TableName)

*/
