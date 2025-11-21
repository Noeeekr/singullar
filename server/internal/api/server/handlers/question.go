package handlers

import (
	"errors"
	"net/http"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetQuestionList(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, types.NewServerResponse([]any{}, "Nenhuma lista de questões encontrada"))
}

func (h *Handlers) GetQuestionListDifficulties(ctx *gin.Context) {
	difficulties := []*models.QuestionDifficulty{}
	query := models.TableQuestionDifficulty.
		Select("difficulty_name", "difficulty_level").
		Scanner(scan.QuestionDifficulties(&difficulties)).
		OrderAscending("difficulty_level")

	if err := h.databaseManager.Do(query); err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusOK, types.NewServerResponse([]any{}))
			return
		}
		h.internalError(ctx, "Falha ao procurar as dificuldades da lista de questões", err)
		return
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(difficulties))
}

type GetQuestionRequest struct {
	Filters *[]*models.GetQuestionRequest `json:"filters" binding:"required"`
}

func (h *Handlers) GetQuestions(ctx *gin.Context) {
	filters := GetQuestionRequest{}
	requester, _ := h.GetRequestUserInformation(ctx)

	if h.HandleBadJsonRequest(ctx, &filters) {
		return
	}

	questions, err := h.databaseManager.SelectQuestions(requester.InstitutionId, filters.Filters)
	if err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusOK, types.NewServerResponse([]any{}))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao procurar questões"))
		return
	}

	ctx.JSON(http.StatusCreated, types.NewServerResponse(questions))
}

func (h *Handlers) CreateQuestionList(ctx *gin.Context) {
	request := &models.CreateQuestionListRequest{}
	requester, _ := h.GetRequestUserInformation(ctx)

	if h.HandleBadJsonRequest(ctx, request) {
		return
	}

	operator, err := h.databaseManager.NewTransactionOperator()
	if err != nil {
		h.internalError(ctx, "Falha ao executar transação", err)
		return
	}

	list, err := operator.InsertQuestionList(requester.Id, request)
	if err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusNotAcceptable, types.NewFailedResponse("Falha ao encontrar as questões selecionadas"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewFailedResponse("Falha ao criar a lista de questões"))
		return
	}

	if err = operator.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewFailedResponse("Falha ao salvar mudanças"))
		return
	}

	ctx.JSON(http.StatusCreated, types.NewSuccessResponse(list))
}

func (h *Handlers) CreateQuestion(ctx *gin.Context) {
	request := &models.CreateQuestionRequest{}

	requester, _ := h.GetRequestUserInformation(ctx)

	if h.HandleBadJsonRequest(ctx, request) {
		return
	}

	operator, err := h.databaseManager.NewTransactionOperator()
	if err != nil {
		h.internalError(ctx, "Falha ao executar transação", err)
		return
	}

	question, err := operator.InsertQuestion(requester.InstitutionId, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao criar a questão"))
		return
	}

	if err = operator.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewFailedResponse("Falha ao salvar mudanças"))
		return
	}

	ctx.JSON(http.StatusCreated, types.NewServerResponse(question))
}
