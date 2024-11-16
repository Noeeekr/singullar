package server

import (
	"errors"
	"net/http"

	"github.com/noeeekr/sch-server/config"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type RouterMiddlewares struct {
	env *config.Configuration
}

// Need to be tested : Redirect users that are not logged from protected routes.
func (m *RouterMiddlewares) Authenticate(ctx *gin.Context) {
	cookie, err := ctx.Cookie("auth")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get auth cookie",
			"data":  nil,
		})
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["User"];

		ctx.Set("User",user);
		ctx.Next();
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid auth token",
			"data":  nil,
		})
	}
}