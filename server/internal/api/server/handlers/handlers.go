package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// COOKIE BASED AUTH
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common"
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

// Errors
//
//	borm.ErrNotFound
//	bcrypt.ErrMismatchedPasswords
func (h *Handlers) checkUserPassword(email string, password string) (*models.Users, common.ResponseStatus, error) {
	user, err := h.databaseOperations.SelectUserByEmail(email)
	if err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			return nil, common.StatusNotFound, err
		}
		return nil, common.StatusInternalError, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, common.StatusNotEqual, err
		}
		return nil, common.StatusInternalError, err
	}

	user[0].Password = ""
	return user[0], common.StatusEmpty, nil
}

// SingInHandler gets a SignInRequest, check the user email and password agaisnt database.
// If access is granted it creates an auth cookie and returns a json
// If access is not granted it returns a json
func (h *Handlers) SignIn(ctx *gin.Context) {
	var request types.SignInRequest
	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindJSON(&request)) {
		return
	}

	user, status, err := h.checkUserPassword(request.Email, request.Password)
	if err != nil {
		switch status {
		default:
			h.internalError(ctx, "Falha ao checar se o usuario existe. ", err)
			return
		case common.StatusNotFound:
			h.clientError(ctx, "Usuario não existe")
			return
		case common.StatusNotEqual:
			h.internalError(ctx, "Senha incorreta.", err)
			return
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.AuthClaims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    fmt.Sprintf("%d", user.Id),
		},
	})

	stringifiedToken, err := token.SignedString([]byte(h.JwtSecret))
	if err != nil {
		h.internalError(ctx, " Falha ao validar o usuario. ", err)
	}

	ctx.SetCookie(
		"auth",
		stringifiedToken,
		3600,
		"/",
		h.Environment.Domain,
		false, // SHOULD BE TRUE IN HTTPS
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  user,
	})
}
func (h *Handlers) SignOut(ctx *gin.Context) {
	ctx.SetCookie(
		"auth",
		"",
		-1,
		"/",
		h.Environment.Domain,
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  nil,
		"error": nil,
	})
}

func (h *Handlers) CreateQuestion(ctx *gin.Context) {
	request := &models.CreateQuestionRequest{}

	requester, _ := h.GetRequestUserInformation(ctx)

	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(request)) {
		return
	}

	if _, err := h.databaseOperations.InsertQuestion(requester.InstitutionId, request); err != nil {
		println(err.Error())
		ctx.JSON(http.StatusInternalServerError, types.NewServerResponse(nil, "Falha ao criar a questão"))
	}
}
func (h *Handlers) CreateClass(ctx *gin.Context) {
	var request *models.CreateClassRequest = &models.CreateClassRequest{}

	unsignedUser, _ := ctx.Get(types.REQUEST_USER_TOKEN)
	requester := unsignedUser.(models.Users)

	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(request)) {
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
	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindJSON(&request)) {
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
	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(&request)) {
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
	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(&filters)) {
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
	if h.HandleBadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(&request)) {
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

func (h *Handlers) GetQuestionList(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, types.NewServerResponse([]any{}, "Nenhuma lista de questões encontrada"))
}

func (h *Handlers) GetQuestionListDifficulties(ctx *gin.Context) {
	difficulties := []*models.QuestionDifficulty{}
	query := models.TableQuestionDifficulty.
		Select("difficulty_name", "difficulty_level").
		Scanner(scan.QuestionDifficulties(&difficulties)).
		OrderAscending("difficulty_level")

	if err := h.databaseOperations.Do(query); err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			ctx.JSON(http.StatusOK, types.NewServerResponse([]any{}))
			return
		}
		h.internalError(ctx, "Falha ao procurar as dificuldades da lista de questões", err)
		return
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(difficulties))
}
