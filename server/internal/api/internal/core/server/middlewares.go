package server

import (
	"errors"
	"net/http"

	"github.com/Noeeekr/singullar/server/api/config"
	"github.com/Noeeekr/singullar/server/api/internal/core/models"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type RouterMiddlewares struct {
	env *config.Configuration
}

// Need to be tested : Redirect users that are not logged from protected routes.
// ctx - gin default ctxx
// role - a role to check if user is part of before liberating access - if fails send json back with error
func (m *RouterMiddlewares) Authenticate(roles ...models.UserRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("auth")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to get auth cookie",
				"data":  nil,
			})
			return
		}
		claims := &AuthClaims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("failed to parse token")
			}
			return []byte(m.env.JwtSecret), nil
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to parse cookie",
				"data":  nil,
			})
			return
		}

		if token.Valid {
			var canPass bool = false

			for i := 0; i < len(roles); i++ {
				if claims.User.Role == roles[i] {
					canPass = true
				}
			}

			if !canPass {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Usuário não autorizado.",
					"data":  nil,
				})
				return
			}

			ctx.Set("User", claims.User)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid auth token",
				"data":  nil,
			})
		}
	}
}
