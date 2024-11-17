package server

import (
	"github.com/noeeekr/sch-server/config"
	logs "github.com/noeeekr/sch-server/internal/core/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/noeeekr/sch-server/internal/models/pgsql"
	"gorm.io/gorm"
)

func getRouter(database *gorm.DB) (*gin.Engine, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	handlers := RouterHandlers{
		LogErr:       logs.LogErr,
		LogInfo:      logs.LogInfo,
		env:          cfg,
		users:        &pgsql.UserModel{DB: database},
		institutions: &pgsql.InstitutionModel{DB: database},
	}

	middlewares := RouterMiddlewares{
		env: cfg,
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
	r.POST("/api/user/signin", handlers.SigninHandler)  // For users and institutions
	r.GET("/api/user/signout", handlers.SignoutHandler) // For users and institutions

	r.GET("/api/user/auth", handlers.Authenticate) // For users and institutions

	// CREATE
	r.POST("/api/user/signup", handlers.UserSignupHandler) /// for users and institutions

	// READ
	r.GET("/api/user/institutions", middlewares.Authenticate, handlers.GetInstitutions)

	// r.GET("/api/institution/users", middlewares.Authenticate.GetInstitutionUsers)
	// r.GET("/api/institution/classes", middlewares.Authenticate.GetInstitutionClasses)

	// r.POST("/api/auth/notifications", handlers.CreateNotificationsHandler)
	// r.GET("/api/auth/notifications", handlers.GetNotificationsHandler)

	// r.Post("/api/auth/classes", handlers.CreateClassesHandler)
	// r.GET("/api/auth/classes", handlers.GetClassesHandler)

	return r, nil
}
