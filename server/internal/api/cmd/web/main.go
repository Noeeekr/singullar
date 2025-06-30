package main

import (
	"flag"

	config "github.com/Noeeekr/singullar/server/api/config"
	logs "github.com/Noeeekr/singullar/server/api/internal/core/log"
	server "github.com/Noeeekr/singullar/server/api/internal/core/server"
)

func main() {
	// Need modification to become an array
	env_path := flag.String("env-path", "./config/dev.env", "sets the absolute enviroment path for server setup.")
	flag.Parse()

	logs.InitLoggers()

	// Env config
	_, err := config.NewConfig(*env_path)
	if err != nil {
		logs.LogInfo.Fatalf("An error happenned in the setup of environment: %q", err)
	}

	// Server and Routing
	err = server.ServeAndListen()
	if err != nil {
		logs.LogErr.Fatal(err)
	}

	// implement a graceful shutdown for server
}
