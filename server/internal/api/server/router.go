package server

import (
	"strings"

	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/server/handlers"
	"github.com/Noeeekr/singullar/server/internal/api/server/middlewares"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PrepareRouter(handlers *handlers.Handlers, env *types.Environment) (*gin.Engine, error) {
	middlewares := middlewares.Middlewares{
		Environment: env,
	}

	r := gin.Default()
	if env.Environment != "production" {
		gin.SetMode(gin.DebugMode)
	}

	logs.Info.Printf("Allowed Origins:\n\t%s", strings.ReplaceAll(env.AllowedOrigins, ",", "\n\t"))
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
	r.POST("/api/auth/", handlers.SignIn) // Validate through database, create validation cookie
	r.GET("/api/auth/", handlers.SignOut) // Delete validation cookie

	// Select data related
	// Institution data related routes
	r.GET("/api/subjects", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.GetSubjects)
	r.GET("/api/dashboard", middlewares.Authenticate(models.SUPERVISOR, models.ADMIN), handlers.GetDashboard)
	r.GET("/api/institution", middlewares.Authenticate(models.STUDENT, models.ADMIN, models.TEACHER, models.SUPERVISOR), handlers.GetInstitution)
	// Question data related routes
	r.GET("/api/question/list/difficulties", middlewares.Authenticate(models.STUDENT, models.TEACHER, models.ADMIN, models.SUPERVISOR), handlers.GetQuestionListDifficulties)
	r.GET("/api/question/list", middlewares.Authenticate(models.STUDENT, models.TEACHER, models.SUPERVISOR, models.ADMIN), handlers.GetQuestionList)
	r.POST("/api/question", middlewares.Authenticate(models.STUDENT, models.TEACHER, models.SUPERVISOR, models.ADMIN), handlers.GetQuestions)
	// User data related routes
	r.POST("/api/students", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.GetStudents)
	r.POST("/api/users", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.GetUsers)
	// Class data related rotues
	r.POST("/api/classes", middlewares.Authenticate(models.STUDENT, models.TEACHER, models.ADMIN, models.SUPERVISOR), handlers.GetClasses)

	// Create data related
	r.POST("/api/subjects/create", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.CreateSubject)
	r.POST("/api/user/create", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.CreateUser) // Institution admin creates users
	r.POST("/api/class/create", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.CreateClass)
	r.POST("/api/question/create", middlewares.Authenticate(models.ADMIN, models.SUPERVISOR), handlers.CreateQuestion)
	// r.POST("/api/class/create", middlewares.Authenticate, handlers.CreateClass)
	return r, nil
}
