package server

import (
	"net/http"

	"github.com/noeeekr/sch-server/config"
	logs "github.com/noeeekr/sch-server/internal/core/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/noeeekr/sch-server/pkg/models/pgsql"
	"gorm.io/gorm"
)

func getRouter(database *gorm.DB) (*gin.Engine, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	middlewares := RouterMiddlewares{}
	handlers := RouterHandlers{
		LogErr:  logs.LogErr,
		LogInfo: logs.LogInfo,
		env:     cfg,
		users:   &pgsql.UserModel{DB: database},
	}

	r := gin.Default()

	// User auth session and store for authentication
	store := cookie.NewStore([]byte(cfg.UserAuthStoreSecret))
	store.Options(sessions.Options{
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	r.Use(sessions.Sessions("userId", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Static Routes
	r.NoRoute(middlewares.disableSession, handlers.StaticsHandler)

	// Auth routes
	r.POST("/api/auth/signin", handlers.SigninHandler)
	r.POST("/api/auth/signout", handlers.SignupHandler)
	r.GET("/api/user/auth", handlers.Authenticate)

	// Private routes
	r.GET("/home", middlewares.AuthMiddleware, handlers.StaticsHandler)

	return r, nil
}
