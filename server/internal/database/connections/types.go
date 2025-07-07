package connections

type Connection interface {
	User() string
	Password() string
	Database() string
	Host() string
}

// Defines the environment to be used for database connection. The environment may change the required user, password, database and host from environment variable
type ConnectionEnvironment string
