package config

import (
	"log"
	"path/filepath"

	"github.com/caarlos0/env"
	dotenv "github.com/joho/godotenv"
	abspath "github.com/noeeekr/sch-server/pkg/absolutepath"
)

type Configuration struct {
	Addr        string `env:"ADDR" envDefault:":8000"`
	StaticsPath string `env:"STATICS_PATH" envDefault:"./static"`
}

func NewConfig(files ...string) (*Configuration, error) {
	if len(files) == 0 {
		log.Fatal("No .env file(s) provided to setup environment.")
	}

	_files := filePathsToAbs(files)

	err := dotenv.Load(_files...)
	if err != nil {
		log.Println("Failed to load enviroment files. It is not secure to start the program without proper setup. Aborting..")
		log.Fatal(err)
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
		files[i] = filepath.Join(abspath.Root, f)
	}

	return files
}
