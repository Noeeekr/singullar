package server

import (
	"strings"

	"github.com/Noeeekr/singullar/server/internal/api/server/handlers"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PrepareRouter(handlers *handlers.Handlers, env *types.Environment) (*gin.Engine, error) {
	middlewares := RouterMiddlewares{
		Environment: env,
	}

	r := gin.Default()
	if env.Environment != "production" {
		gin.SetMode(gin.DebugMode)
	}

	util.Info.Printf("Allowed Origins:\n\t%s", strings.ReplaceAll(env.AllowedOrigins, ",", "\n\t"))
	// User auth session and store for authentication
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(env.AllowedOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	// Auth routes
	r.POST("/api/auth/", handlers.SignIn) // LOGIN - VALIDATE VIA DATABASE, CREATE COOKIE
	r.GET("/api/auth/", handlers.SignOut) // LOGOUT - DELETE COOKIE

	// Is this even necessary?
	// r.GET("/api/user/auth", handlers.Authenticate) // For users and institutions

	// SELECT
	r.GET("/api/institution/", middlewares.Authenticate(models.STUDENT, models.ADMIN, models.TEACHER, models.SUPERVISOR), handlers.GetInstitution)
	r.POST("/api/institution/students", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.GetStudents)
	r.POST("/api/institution/users", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.GetUsers)

	// CREATE
	r.POST("/api/user/create/", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.CreateUser) // Institution admin creates users

	// r.POST("/api/class/create", middlewares.Authenticate, handlers.CreateClass)

	return r, nil
}

// People interested ask devs for a institution
// It is created with a admin user..
// Then the admin can log to do admin stuff, like creating users for teachers, supervisors, students, etc...
// Then users can login
