package models

type CreateDatabaseUser struct {
	Name     string
	Password string
	Database string
}

func NewDatabaseUser(name, password, database string) *CreateDatabaseUser {
	return &CreateDatabaseUser{
		Name:     name,
		Password: password,
		Database: database,
	}
}
