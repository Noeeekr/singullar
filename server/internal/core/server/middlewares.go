package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RouterMiddlewares struct{}

// Need to be tested : Disable session in context for the next handlers.
func (m *RouterMiddlewares) disableSession(ctx *gin.Context) {

	ctx.Set("auth", nil)
	ctx.Next()

}

// Need to be tested : Redirect users that are not logged from protected routes.
func (m *RouterMiddlewares) AuthMiddleware(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("auth")

	fmt.Println(ctx.Request.URL)

	if user != nil {
		fmt.Println("PROTECTED ROUTE : USER AUTHORIZED : NO REDIRECTED : USER IS ", user)

		ctx.Next()
	}

	fmt.Println("PROTECTED ROUTE : USER UNAUTHORIZED : REDIRECT : USER IS ", user)
	ctx.Redirect(http.StatusUnauthorized, "/")
}
