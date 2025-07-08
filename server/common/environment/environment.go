package environment

import (
	"os"

	"github.com/Noeeekr/singullar/server/common"
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

// Set "key" to "value" if "key" is empty. If "value" is empty mantains "key" as it is. If "key" is not empty ignores it. Returns the value present in the key at the end.
func SetIfNotEmpty(key string, value string) (string, error) {
	if value != "" && os.Getenv(key) == "" {
		return value, os.Setenv(key, value)
	}
	return os.Getenv(key), nil
}
func Parse(files ...string) *common.Response {
	if len(files) == 0 {
		return nil
	}

	for _, file := range files {
		stat, err := os.Stat(file)
		if err != nil {
			return common.NewResponse().
				WithStatus(common.StatusNotFound).
				WithDescription("Environment file not found.")
		}
		if stat.IsDir() {
			return common.NewResponse().
				WithStatus(common.StatusInvalidRequest).
				WithDescription("Path doesn't lead to an actual file.")
		}
	}

	if err := dotenv.Load(files...); err != nil {
		return common.NewResponse().
			WithStatus(common.StatusInternalError).
			WithDescription("Failed to load environment file")
	}

	return nil
}

// Populates v with all variables. V must be a struct with all fields to be parsed
func Scan(v any) *common.Response {
	if err := env.Parse(v); err != nil {
		return common.NewResponse().
			WithStatus(common.StatusInternalError).
			WithDescription("Failed to scan environment variables to interface. " + err.Error())
	}
	return nil
}
