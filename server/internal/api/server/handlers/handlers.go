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
)

type Handlers struct {
	LogInfo *log.Logger
	LogErr  *log.Logger

	operations *operations.Operations

	*types.Environment
}

func New(ops *operations.Operations, env *types.Environment) *Handlers {
	return &Handlers{
		LogInfo:    logs.Info,
		LogErr:     logs.Error,
		operations: ops,

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

// Errors
//
//	borm.ErrNotFound
//	bcrypt.ErrMismatchedPasswords
func (h *Handlers) checkUserPassword(email string, password string) (*models.Users, common.ResponseStatus, error) {
	user, err := h.operations.SelectUserByEmail(email)
	if err != nil {
		if errors.Is(err, borm.ErrNotFound) {
			return nil, common.StatusNotFound, err
		}
		return nil, common.StatusInternalError, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, common.StatusNotEqual, err
		}
		return nil, common.StatusInternalError, err
	}

	user.Password = ""
	return user, common.StatusEmpty, nil
}

// SingInHandler gets a SignInRequest, check the user email and password agaisnt database.
// If access is granted it creates an auth cookie and returns a json
// If access is not granted it returns a json
func (h *Handlers) SignIn(ctx *gin.Context) {
	var request types.SignInRequest
	if h.BadJsonRequest(ctx, ctx.ShouldBindJSON(&request)) {
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

func (h *Handlers) CreateUser(ctx *gin.Context) {
	var user *models.Users = &models.Users{}

	// Only admins can create other users so you can use admin.InstitutionId to attribute created users ids

	var request models.CreateUsers
	if h.BadJsonRequest(ctx, ctx.ShouldBindJSON(&request)) {
		return
	}

	_, err := h.operations.SelectUserByEmail(request.Email)
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

	users, err := h.operations.InsertManyUsers(nil, userRequest)
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
	unsignedUser, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": nil, "error": "No authentication data found for the user"})
		return
	}

	user, ok := unsignedUser.(models.Users)
	if !ok {
		h.internalError(ctx, "Failed to get user", nil)
		return
	}

	institution, err := h.operations.SelectInstitutionById(user.InstitutionId)
	if errors.Is(err, borm.ErrNotFound) {
		h.clientError(ctx, "Nenhuma instituição encontrada")
		return
	} else if err != nil {
		h.internalError(ctx, "Falha ao buscar as instituições", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  institution,
		"error": nil,
	})
}

type GetUsersRequest struct {
	TargetRoles []models.UserRole `json:"target_roles" binding:"required"`
}

func (h *Handlers) GetUsers(ctx *gin.Context) {
	var request GetUsersRequest
	if h.BadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(&request)) {
		return
	}

	unsignedUser, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, types.NewServerResponse(nil, "Usuário não autorizado"))
		return
	}
	user := unsignedUser.(models.Users)

	users, err := h.operations.SelectUsers(user.Id, request.TargetRoles...)
	if err != nil {
		h.internalError(ctx, "Falha ao procurar usuários", err)
		return
	}

	for _, user := range users {
		user.Password = ""
	}

	ctx.JSON(http.StatusOK, types.NewServerResponse(users, ""))
}
