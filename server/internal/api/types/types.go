package types

import (
	"strings"

	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/golang-jwt/jwt/v5"
)

type Environment struct {
	JwtSecret string `env:"API_JWT_SECRET,required"`

	Domain         string `env:"API_DOMAIN"`
	Port           string `env:"API_PORT,required"`
	AllowedOrigins string `env:"API_ALLOWED_ORIGINS,required"`

	Environment string `env:"API_ENVIRONMENT"`
}

type SignInRequest struct { // FOR JSON
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=2"`
}

// Handles user sign-in and returns custom graceful JSON objects for client form ui.
type AuthClaims struct {
	User models.Users
	jwt.RegisteredClaims
}

type ServerResponse map[string]any

const (
	// REQUEST_USER_TOKEN refers to the token used to store the user struct
	REQUEST_USER_TOKEN string = "user"
)

func NewServerResponse(data any, err ...string) *ServerResponse {
	return &ServerResponse{
		"data":  data,
		"error": strings.Join(err, ": "),
	}
}

func NewSuccessResponse(data any) *ServerResponse {
	return &ServerResponse{
		"data":  data,
		"error": "",
	}
}

func NewFailedResponse(err ...string) *ServerResponse {
	return &ServerResponse{
		"data":  nil,
		"error": strings.Join(err, ": "),
	}
}
