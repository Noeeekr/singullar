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

	subjects, err := h.databaseManager.SelectSubjects(user.InstitutionId)
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

	if h.HandleBadJsonRequest(ctx, request) {
		return
	}

	operator, err := h.databaseManager.NewTransactionOperator()
	if err != nil {
		h.internalError(ctx, "Falha ao iniciar requisição", err)
		return
	}

	subject, err := operator.InsertSubject(requester.InstitutionId, request)

	if err != nil {
		if errors.Is(err, borm.ErrFound) {
			ctx.JSON(http.StatusBadRequest, types.NewServerResponse(nil, "Matéria já existe."))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao criar a questão"))
	}

	if err = operator.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewFailedResponse("Falha ao salvar mudanças"))
		return
	}

	ctx.JSON(http.StatusCreated, types.NewServerResponse(subject))

}
