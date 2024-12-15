package config

import (
	"errors"
	"path/filepath"

	"github.com/caarlos0/env"
	dotenv "github.com/joho/godotenv"
	"github.com/noeeekr/sch-server/pkg/paths"
)

type Configuration struct {
	FrontendUrl string `env:"FRONTEND_URL,required"` // For cors security. This represents the permitted url. 
	Addr        string `env:"ADDR" envDefault:"80"` // The port to listen to. The default is internet 80 port.

	StaticsPath string `env:"STATICS_PATH" envDefault:"./static"` // Client files.

	DatabaseWR_ConnectString string `env:"DB_WR_CONNECTION_STR,required"` // For reading and creating data only.

	JwtSecret           string `env:"JWT_SECRET,required"` // Secret for json web token for client auth.
	UserAuthStoreSecret string `env:"USER_AUTH_STORE_SECRET,required"` 
}

func NewConfig(files ...string) (*Configuration, error) {
	if len(files) == 0 {
		return nil, errors.New(" No env file location provided. Starting a server without proper configuration is dangeous. Please provide a config file in --env flag, or create a default config file in ./config/* as dev.env. ")
	}

	_files := filePathsToAbs(files)

	err := dotenv.Load(_files...)
	if err != nil {
		return nil, errors.New(" Failed to load enviroment files. It is not secure to start the program without proper setup. ")
	}

	return GetConfig()
}

func GetConfig() (*Configuration, error) {
	cfg := Configuration{}

	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func filePathsToAbs(files []string) []string {
	for i, f := range files {
		files[i] = filepath.Join(paths.Root, f)
	}

	return files
}
