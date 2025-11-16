package handlers

import (
	"errors"
	"net/http"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetSubjects(ctx *gin.Context) {
	unsignedUser, _ := ctx.Get(types.REQUEST_USER_TOKEN)
	user := unsignedUser.(models.Users)

	subjects, err := h.databaseOperations.SelectSubjects(user.InstitutionId)
	if err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusFound, types.NewServerResponse([]models.Subjects{}))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao procurar matérias"))
		return
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(subjects))
}

func (h *Handlers) CreateSubject(ctx *gin.Context) {
	request := &models.CreateSubjectsRequest{}

	requester, _ := h.GetRequestUserInformation(ctx)

	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(request)) {
		return
	}

	if subject, err := h.databaseOperations.InsertSubject(requester.InstitutionId, request); err != nil {
		if errors.Is(err, borm.ErrFound) {
			ctx.JSON(http.StatusBadRequest, types.NewServerResponse(nil, "Matéria já existe."))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao criar a questão"))
	} else {
		ctx.JSON(http.StatusCreated, types.NewServerResponse(subject))
	}
}
