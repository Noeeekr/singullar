package core

import (
	"log"
	"net/http"
	"os"
)

func ServeAndListen() error {
	var port string = os.Getenv("PORT")
	r := Router()

	server := http.Server{
		Addr:    port,
		Handler: r,
	}

	log.Printf("Server is running on port http://localhost:%s", port)

	return server.ListenAndServe()
}
