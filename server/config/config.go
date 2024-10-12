package config

import (
	"errors"
	"path/filepath"

	"github.com/caarlos0/env"
	dotenv "github.com/joho/godotenv"
	logs "github.com/noeeekr/sch-server/internal/core/log"
	"github.com/noeeekr/sch-server/pkg/paths"
)

type Configuration struct {
	FrontendUrl string `env:"FRONTEND_URL,required"`
	Addr        string `env:"ADDR" envDefault:"8000"`

	StaticsPath string `env:"STATICS_PATH" envDefault:"./static"`

	DatabaseWR_ConnectString string `env:"DB_WR_CONNECTION_STR,required"`

	JwtSecret           string `env:"JWT_SECRET,required"`
	UserAuthStoreSecret string `env:"USER_AUTH_STORE_SECRET,required"`
}

func NewConfig(files ...string) (*Configuration, error) {
	if len(files) == 0 {
		logs.LogInfo.Println(" No env file location provided. starting a server without proper configuration is dangeous. ")
		return nil, errors.New(" No env file location provided. starting a server without proper configuration is dangeous. ")
	}

	_files := filePathsToAbs(files)

	err := dotenv.Load(_files...)
	if err != nil {
		logs.LogErr.Println("Failed to load enviroment files. It is not secure to start the program without proper setup.")
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
