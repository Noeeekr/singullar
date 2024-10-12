package server

import (
	"github.com/noeeekr/sch-server/config"
	logs "github.com/noeeekr/sch-server/internal/core/log"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendUrl}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Static Routes
	r.NoRoute(handlers.StaticsHandler)

	// Auth routes
	r.POST("/api/auth/signin", handlers.SigninHandler)
	r.POST("/api/auth/signup", handlers.SignupHandler)
	r.GET("/api/auth/signout", handlers.SignoutHandler)
	r.GET("/api/user/auth", handlers.Authenticate)

	// Statics private routes session renew
	r.GET("/home", middlewares.AuthMiddleware, handlers.StaticsHandler)

	return r, nil
}
