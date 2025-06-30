package main

type EnvironmentVariables struct {
	// The port to listen to. The default is internet 80 port.
	Port string `env:"PORT" envDefault:"8700"`
}
