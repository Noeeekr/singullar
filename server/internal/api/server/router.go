package server

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/internal/api/server/handlers"
	"github.com/Noeeekr/singullar/server/internal/api/server/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PrepareRouter(db *sql.DB, env *types.ApiEnvironment) (*gin.Engine, error) {
	handlers := handlers.New(operations.New(db), env)

	middlewares := RouterMiddlewares{
		env: env,
	}

	r := gin.Default()

	// User auth session and store for authentication
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{env.ClientUrl}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Auth routes
	r.POST("/api/auth", handlers.SignIn) // LOGIN - VALIDATE VIA DATABASE, CREATE COOKIE
	r.GET("/api/auth", handlers.SignOut) // LOGOUT - DELETE COOKIE

	// Is this even necessary?
	// r.GET("/api/user/auth", handlers.Authenticate) // For users and institutions

	// SELECT
	r.GET("/api/institution", middlewares.Authenticate(models.Student, models.Admin, models.Teacher, models.Supervisor), handlers.GetInstitution)
	r.GET("/api/user", middlewares.Authenticate(models.Admin, models.Supervisor), handlers.GetUsersByInstitutionId)

	// CREATE
	r.POST("/api/user/create", middlewares.Authenticate(models.Admin, models.Supervisor), handlers.CreateUser) // Institution admin creates users

	// r.POST("/api/class/create", middlewares.Authenticate, handlers.CreateClass)

	return r, nil
}

// People interested ask devs for a institution
// It is created with a admin user..
// Then the admin can log to do admin stuff, like creating users for teachers, supervisors, students, etc...
// Then users can login
