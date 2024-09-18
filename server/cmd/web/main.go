package main

import (
	"flag"
	"log"

	config "github.com/noeeekr/sch-server/config"
	"github.com/noeeekr/sch-server/internal/core"
)

func main() {
	// Flags setup
	env_path := flag.String("env-path", "./config/dev.env", "sets the absolute enviroment path for server setup.")

	flag.Parse()

	// Env config
	_, err := config.NewConfig(*env_path)
	if err != nil {
		log.Fatalf("An error happenned in the setup of environment. %q", err)
	}

	// Server and Routing
	log.Fatal(core.ServeAndListen())
}
