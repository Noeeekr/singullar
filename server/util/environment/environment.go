package environment

import (
	"os"

	"github.com/Noeeekr/singullar/server/util"
	"github.com/caarlos0/env"
	dotenv "github.com/joho/godotenv"
)

type EnvironmentMode string

// Environment represents the environment configuration for the application.
// Extend this with the necessary environment variables
type Environment struct {
	Mode EnvironmentMode `env:"APPLICATION_ENVIRONMENT,required"` // The environment mode. It can be "development", "production" or "test".
}

const (
	PRODUCTION  EnvironmentMode = "production"
	DEVELOPMENT EnvironmentMode = "development"
)

var environment *Environment = &Environment{
	// Default values
	Mode: PRODUCTION,
}

func init() {
	if err := Scan(environment); err != nil {
		res := util.NewResponse().
			WithDescription("Unable to start program: Environment variables not set correctly.").
			WithStatus(util.StatusInternalError)
		util.Error.Println(res.String() + "\n\t" + err.String())
		os.Exit(1)
	}
}

func (e *Environment) IsProduction() bool {
	return e.Mode == PRODUCTION
}
func (e *Environment) IsDevelopment() bool {
	return e.Mode == DEVELOPMENT
}
func (e *Environment) ApplicationMode() EnvironmentMode {
	return e.Mode
}
func (e *Environment) SetApplicationMode(mode EnvironmentMode) {
	e.Mode = mode
}

// Set "key" to "value" if "key" is empty. If "value" is empty mantains "key" as it is. If "key" is not empty ignores it. Returns the value present in the key at the end or an empty string if an error was found.
func OverrideEmpty(key string, newValue string) string {
	originalValue := os.Getenv(key)
	if newValue != "" && originalValue == "" {
		if err := os.Setenv(key, newValue); err != nil {
			return ""
		}
		return newValue
	}
	return originalValue
}
func Parse(files ...string) *util.Response {
	if len(files) == 0 {
		return nil
	}

	for _, file := range files {
		stat, err := os.Stat(file)
		if err != nil {
			return util.NewResponse().
				WithStatus(util.StatusNotFound).
				WithDescription("Environment file not found.")
		}
		if stat.IsDir() {
			return util.NewResponse().
				WithStatus(util.StatusInvalidRequest).
				WithDescription("Path doesn't lead to an actual file.")
		}
	}

	if err := dotenv.Load(files...); err != nil {
		return util.NewResponse().
			WithStatus(util.StatusInternalError).
			WithDescription("Failed to load environment file")
	}

	return nil
}

// Populates v with all variables. V must be a struct with all fields to be parsed
func Scan(v any) *util.Response {
	if err := env.Parse(v); err != nil {
		return util.NewResponse().
			WithStatus(util.StatusInternalError).
			WithDescription("Failed to retrieve environment variables to application: " + err.Error())
	}
	return nil
}

func Settings() *Environment {
	return environment
}
