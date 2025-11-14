package handlers

import (
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
		if err == borm.ErrNotFound {
			ctx.JSON(http.StatusFound, types.NewServerResponse([]models.Subjects{}))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao procurar matérias"))
		return
	}

	ctx.JSON(http.StatusFound, types.NewServerResponse(subjects))
}
