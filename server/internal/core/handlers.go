package core

import (
	"net/http"
	"os"
	"path/filepath"

	abspath "github.com/noeeekr/sch-server/pkg/absolutepath"
)

// Serves the static files in $STATICS_PATH or $STATICS_PATH/index.html if the file
// doesn't exist, so react-router-dom can handle the redirect to not found page.

func StaticsHandler() http.Handler {
	staticsDir := filepath.Join(abspath.Root, os.Getenv("STATICS_PATH"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticsDir, r.URL.Path)

		_, err := os.Stat(path)
		if err != nil {
			http.ServeFile(w, r, filepath.Join(staticsDir, "index.html"))
			return
		}

		http.ServeFile(w, r, path)
	})
}
