package configs

import (
	"errors"

	"github.com/caarlos0/env"
	dotenv "github.com/joho/godotenv"
)

// Extend this with the necessary environment variables
type EnvironmentVariables struct {
	// Optional
	// Addr string `env:"ADDR" envDefault:"80"` // The port to listen to. The default is internet 80 port.

	// FrontendUrl string `env:"FRONTEND_URL,required"` // For cors security. This represents the permitted url.
	// StaticsPath string `env:"STATICS_PATH" envDefault:"./static"` // Client files.

	// DatabaseWR_ConnectString string `env:"DB_WR_CONNECTION_STR,required"` // For reading and creating data only.

	// JwtSecret           string `env:"JWT_SECRET,required"` // Secret for json web token for client auth.
	// UserAuthStoreSecret string `env:"USER_AUTH_STORE_SECRET,required"`
}

func Parse(files ...string) error {
	if len(files) == 0 {
		return errors.New(" No env file location provided. Starting a server without proper configuration is dangeous. Please provide a config file in --env flag, or create a default config file in ./config/* as dev.env. ")
	}

	if err := dotenv.Load(files...); err != nil {
		return errors.New(" Failed to load enviroment file. " + err.Error())
	}
	return nil
}

// Populates v with all variables. V must be a struct with all fields to be parsed
func Scan(v any) error {
	if err := env.Parse(v); err != nil {
		return err
	}
	return nil
}
