package types

import (
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/golang-jwt/jwt/v5"
)

type ApiEnvironment struct {
	JwtSecret   string `env:"JWT_SECRET,required"`
	Port        string `env:"PORT,required"`
	FrontendUrl string `env:"FRONTEND_URL,required"`
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
