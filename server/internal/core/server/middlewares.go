package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/noeeekr/sch-server/config"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type RouterMiddlewares struct{}

// Need to be tested : Redirect users that are not logged from protected routes.
func (m *RouterMiddlewares) AuthMiddleware(ctx *gin.Context) {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println("ERR 1", err)

		ctx.Redirect(http.StatusFound, "/")
	}
	cookie, err := ctx.Cookie("auth")
	if err != nil {
		fmt.Println("ERR 2", err)

		ctx.Redirect(http.StatusFound, "/")
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("failed to parse token")
		}
		return []byte(cfg.JwtSecret), nil
	})

	if err != nil {
		fmt.Println("ERR 3", err)

		ctx.Redirect(http.StatusFound, "/")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user1 := claims["user"]
		user2 := claims["User"]

		fmt.Println("USER1", user1)
		fmt.Println("USER2", user2)

		ctx.Next()
	} else {
		fmt.Println("ERR 4", token.Valid)

		ctx.Redirect(http.StatusFound, "/")
	}
}
