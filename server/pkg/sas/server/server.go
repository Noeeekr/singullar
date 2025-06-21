package server

import (
	"net/http"

	"github.com/Noeeekr/singullar/server/pkg/common/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	// Static files path
	staticsfolder string

	// Port the server will listen to
	port string
}

// New returns a Static File Server
func New(staticsfolder, port string, debugMode bool) *Server {
	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		staticsfolder,
		port,
	}
}

func (s *Server) ServeAndListen() error {
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
	r.Static("/", s.staticsfolder)
	r.NoRoute(func(ctx *gin.Context) {
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusPermanentRedirect)
	})

	server := http.Server{
		Addr:     ":" + s.port,
		Handler:  r,
		ErrorLog: logs.Error,
	}

	logs.Info.Println("Static file server running on " + server.Addr)
	return server.ListenAndServe()
}
