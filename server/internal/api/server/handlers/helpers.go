package handlers

import (
	"fmt"
	"net/http"

	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/gin-gonic/gin"
)

// Internal server errors sends a json as response with InternalServerError response code.
//
// It accepts a main "errType" and insert the error object inside it to be parsed into JSON for frontend.
// It also accepts an err for debug porpuses.
func (h *Handlers) internalError(ctx *gin.Context, message string, err error) {
	trace := fmt.Sprintf("%s : ", message)
	if err != nil {
		trace += err.Error() + "\n"
	}

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

// Is always success if done after an authentication middleware
func (h *Handlers) GetRequestUserInformation(ctx *gin.Context) (*models.Users, bool) {
	unsignedRequestUser, found := ctx.Get(types.REQUEST_USER_TOKEN)
	user := unsignedRequestUser.(models.Users)
	return &user, found
}

// Handles the client message and returns true if error happens is in incorrect format.
func (h *Handlers) HandleBadJsonRequest(ctx *gin.Context, request any) bool {
	if err := ctx.ShouldBindBodyWithJSON(request); err != nil {
		h.clientError(ctx, "Dados em formato incorreto")
		return true
	}
	return false
}
