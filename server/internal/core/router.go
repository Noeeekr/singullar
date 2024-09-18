package core

import (
	"net/http"
)

func Router() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("/", StaticsHandler())

	return r
}
