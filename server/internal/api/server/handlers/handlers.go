package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// COOKIE BASED AUTH

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
)

type Handlers struct {
	LogInfo *log.Logger
	LogErr  *log.Logger

	databaseOperations *operations.Operations

	*types.Environment
}

func New(ops *operations.Operations, env *types.Environment) *Handlers {
	return &Handlers{
		LogInfo:            logs.Info,
		LogErr:             logs.Error,
		databaseOperations: ops,

		Environment: env,
	}
}

// HTTP ERROR RESPONSE CODE REPLIERS

// AUTH RELATED HANDLERS

// Checks if user email and password matches in database.
//
// For architetural porpuses this function doesn't handle query and bcrypt errors
// instead, they're returned to be handled in the main function.
// Such errors include Gorm.ErrRecordNotFound and Bcrypt.ErrMismatchedHashAndPassword.

/*
func (h *Handlers) Authenticate(ctx *gin.Context) {
	cookie, err := ctx.Cookie("auth")
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			gin.H{"data": nil, "error": "No authentication data found for the user"},
		)
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(" Failed to parse token. Unexpected token method. ")
		}
		return []byte(h.env.JwtSecret), nil
	})

	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			gin.H{"data": nil, "error": "Token in wrong format."},
		)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["User"]

		ctx.JSON(http.StatusOK, gin.H{"data": user, "error": nil})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token.",
			"data":  nil,
		})
	}
}
*/

func (h *Handlers) GetDashboard(ctx *gin.Context) {
	user := &models.Users{}
	{
		unsignedUser, _ := ctx.Get(types.REQUEST_USER_TOKEN)

		*user = (unsignedUser).(models.Users)
	}

	institutions := []*models.Institutions{}
	query := models.TableInstitutions.
		Select("created_at", "updated_at", "deleted_at", "name", "id").
		Scanner(scan.Institutions(&institutions))
	query.Where(query.Field("id").IsEqual(user.InstitutionId))
	if err := h.databaseOperations.Do(query); err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, types.NewServerResponse(nil, "Falha ao encontrar a instituição"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao procurar dados"))
		return
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(institutions[0]))
}
func (h *Handlers) CreateClass(ctx *gin.Context) {
	request := &models.CreateClassRequest{}
	requester, _ := h.GetRequestUserInformation(ctx)

	if h.HandleBadJsonRequest(ctx, request) {
		return
	}

	if _, err := h.databaseOperations.SelectUsersById(requester.Id, request.TeacherId); err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusBadRequest, types.NewServerResponse(nil, "Professor não encontrado"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao encontrar o professor requisitado"))
		return
	}
	if h.databaseOperations.StartTransaction() != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao iniciar criação da classe"))
		return
	}
	class, err := h.databaseOperations.InsertClass(&models.CreateClasses{
		Segment:       request.Segment,
		Series:        request.Series,
		Name:          request.Name,
		InstitutionId: requester.Id,
		TeacherId:     request.TeacherId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao criar a classe"))
		return
	}

	if len(request.StudentsIds) == 0 {
		err = h.databaseOperations.CommitTransaction()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao salvar alterações"))
			return
		}
		ctx.JSON(http.StatusCreated, types.NewServerResponse(class))
		return
	}

	if err := h.databaseOperations.InsertClassStudents(class.Id, &request.StudentsIds); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao adicionar estudantes. Por favor, tente novamente depois"))
		return
	}
	if h.databaseOperations.CommitTransaction() != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao salvar mudanças"))
		return
	}
	ctx.JSON(http.StatusCreated, types.NewServerResponse(class, ""))
}
func (h *Handlers) CreateUser(ctx *gin.Context) {
	var user *models.Users = &models.Users{}

	// Only admins can create other users so you can use admin.InstitutionId to attribute created users ids

	var request models.CreateUsers
	if h.HandleBadJsonRequest(ctx, &request) {
		return
	}

	_, err := h.databaseOperations.SelectUserByEmail(request.Email)
	if errors.Is(err, borm.ErrNotFound) {
		if err != nil {
			h.internalError(ctx, "Falha ao checar se o email já está em uso.", err)
			return
		}
		h.clientError(ctx, "O e-mail já está em uso.")
		return
	}

	userRequest := &models.CreateUsers{
		Name:          request.Name,
		Email:         request.Email,
		Password:      request.Password,
		Role:          request.Role,
		InstitutionId: user.InstitutionId,
	}

	users, err := h.databaseOperations.InsertManyUsers(nil, userRequest)
	if err != nil {
		h.internalError(ctx, "Falha ao criar o usuario.", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  users[0],
		"error": nil,
	})
}

func (h *Handlers) GetInstitution(ctx *gin.Context) {
	unsignedUser, ok := ctx.Get(types.REQUEST_USER_TOKEN)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": nil, "error": "No authentication data found for the user"})
		return
	}

	user, ok := unsignedUser.(models.Users)
	if !ok {
		h.internalError(ctx, "Failed to get user", nil)
		return
	}

	institution, err := h.databaseOperations.SelectInstitutionById(user.InstitutionId)
	if errors.Is(err, borm.ErrNotFound) {
		h.clientError(ctx, "Nenhuma instituição encontrada")
		return
	} else if err != nil {
		h.internalError(ctx, "Falha ao buscar as instituições", err)
		return
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(institution, ""))
}

func (h *Handlers) GetStudents(ctx *gin.Context) {
	var request []operations.FilterStudentOptions
	if h.HandleBadJsonRequest(ctx, &request) {
		return
	}

	unsignedRequestUser, exists := ctx.Get(types.REQUEST_USER_TOKEN)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, types.NewServerResponse(nil, "Usuário não autorizado"))
	}

	requestUser := unsignedRequestUser.(models.Users)
	students, err := h.databaseOperations.SelectStudents(requestUser.InstitutionId, &request, false)
	if err != nil {
		h.internalError(ctx, "Falha ao buscar estudantes sem turmas", err)
		return
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(students, ""))
}

func (h *Handlers) GetClasses(ctx *gin.Context) {
	filters := []*operations.FilterClassesOptions{}
	if h.HandleBadJsonRequest(ctx, &filters) {
		ctx.JSON(http.StatusBadRequest, types.NewServerResponse(nil, "Dados em formato incorreto"))
		return
	}
	unsignedUser, _ := ctx.Get(types.REQUEST_USER_TOKEN)
	user := unsignedUser.(models.Users)

	classes, err := h.databaseOperations.SelectClasses(user.Id, filters...)
	if err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusOK, types.NewServerResponse([]any{}))
			return
		}
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao procurar as classes solicitadas"))
		return
	}
	ctx.JSON(http.StatusOK, types.NewServerResponse(classes))
}

func (h *Handlers) GetUsers(ctx *gin.Context) {
	request := &operations.SelectUsersOptions{}
	if h.HandleBadJsonRequest(ctx, &request) {
		return
	}

	unsignedUser, ok := ctx.Get(types.REQUEST_USER_TOKEN)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, types.NewServerResponse(nil, "Usuário não autorizado"))
		return
	}
	user := unsignedUser.(models.Users)

	users, err := h.databaseOperations.SelectUsers(user.Id, request)
	if err != nil {
		h.internalError(ctx, "Falha ao	 procurar usuários", err)
		return
	}

	for _, user := range users {
		user.Password = ""
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(users, ""))
}
