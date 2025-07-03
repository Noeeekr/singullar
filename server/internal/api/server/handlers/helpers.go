package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Internal server errors sends a json as response with InternalServerError response code.
//
// It accepts a main "errType" and insert the error object inside it to be parsed into JSON for frontend.
// It also accepts an err for debug porpuses.
func (h *Handlers) internalError(ctx *gin.Context, message string, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	h.LogErr.Output(2, trace)

	// err types: validation and default
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
		"data":  nil,
	})
}

func (h *Handlers) clientError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": message,
		"data":  nil,
	})
}

// Handles the client message and returns true if error happens is in incorrect format.
func (h *Handlers) BadJsonRequest(ctx *gin.Context, err error) bool {
	if err != nil {
		h.clientError(ctx, "Dados em formato incorreto")
		return true
	}
	return false
}
