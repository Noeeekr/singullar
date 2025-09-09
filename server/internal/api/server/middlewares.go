package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type RouterMiddlewares struct {
	*types.Environment
}

func (m *RouterMiddlewares) Authenticate(roles ...models.UserRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Check if cookie exists
		cookie, err := ctx.Cookie("auth")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, types.NewServerResponse("", "Falha ao authenticar o usuário"))
			ctx.Abort()
			return
		}

		// Parse cookie default claims and custom claims (values) using the same method used to encrypt
		claims := &types.AuthClaims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("failed to parse token")
			}
			return []byte(m.JwtSecret), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, types.NewServerResponse("", "Falha ao authenticar o usuário"))
			ctx.Abort()
			return
		}

		// Checks the token
		if token.Valid {
			var allowed bool = false

			// Check if the role of the user is allowed
			for i := range roles {
				if claims.User.Role == roles[i] {
					allowed = true
				}
			}

			if !allowed {
				ctx.JSON(http.StatusUnauthorized, types.NewServerResponse("", "Falha ao authenticar o usuário"))
				ctx.Abort()
				return
			}

			// Update the cookie and save it in ctx
			ctx.SetCookie(
				"auth",
				cookie,
				3600,
				"/",
				m.Domain,
				false,
				true,
			)

			fmt.Println("cookie renewed")
			ctx.Set("user", claims.User)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, types.NewServerResponse("", "Falha ao authenticar o usuário"))
			ctx.Abort()
		}
	}
}
