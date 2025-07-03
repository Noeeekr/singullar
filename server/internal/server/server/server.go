package server

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Noeeekr/singullar/server/common/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct{}

// New returns a Static File Server
func New() *Server {
	return &Server{}
}

func (s *Server) ServeFolder() error {
	port := os.Getenv("PORT")

	folder := os.Getenv("FOLDER")
	if folder == "" {
		return errors.New(" Please specify the path to a folder to serve on environment variable: FOLDER ")
	}

	folder, err := filepath.Abs(folder)
	if err != nil {
		return err
	}

	debug := os.Getenv("DEBUG")
	if debug != "" {
		gin.SetMode(gin.ReleaseMode)
	}

	logs.Info.Println("Starting server..")

	r := gin.New()

	// Cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Frontend URL
		AllowMethods:     []string{"GET", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(func(ctx *gin.Context) {
		logs.Info.Println(ctx.Request.Method, ctx.Request.URL)
	})
	r.Static("/", folder)
	r.NoRoute(func(ctx *gin.Context) {
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusPermanentRedirect)
	})

	server := http.Server{
		Addr:     "0.0.0.0:" + port,
		Handler:  r,
		ErrorLog: logs.Error,
	}

	logs.Info.Println("Server running..")
	logs.Info.Println("Folder:", folder)
	logs.Info.Println("Address: http://" + server.Addr)

	return server.ListenAndServe()
}
