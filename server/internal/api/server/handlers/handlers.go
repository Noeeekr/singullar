package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// COOKIE BASED AUTH
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/api/server/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type Handlers struct {
	LogInfo *log.Logger
	LogErr  *log.Logger

	operations *operations.Operations

	env *types.ApiEnvironment
}

func New(ops *operations.Operations, env *types.ApiEnvironment) *Handlers {
	return &Handlers{
		LogInfo:    logs.Info,
		LogErr:     logs.Error,
		operations: ops,
		env:        env,
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

func (h *Handlers) checkUserPassword(email string, password string) (*models.Users, *transactions.Response) {
	user, res := h.operations.SelectUserByEmail(email)
	if res != nil {
		return nil, res
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, transactions.NewResponse().SetDescription(err.Error()).SetStatus(transactions.StatusNotEqual)
		}
		return nil, transactions.NewResponse().SetDescription("Failed to validate users").SetStatus(transactions.StatusFailedTransaction)
	}

	user.Password = ""
	return user, nil
}

// SingInHandler gets a SignInRequest, check the user email and password agaisnt database.
// If access is granted it creates an auth cookie and returns a json
// If access is not granted it returns a json
func (h *Handlers) SignInHandler(ctx *gin.Context) {
	var request types.SignInRequest
	if h.BadJsonRequest(ctx, ctx.ShouldBindJSON(&request)) {
		return
	}

	user, res := h.checkUserPassword(request.Email, request.Password)
	if res != nil {
		switch res.Status {
		case transactions.StatusNotFound:
			h.clientError(ctx, "Usuario não existe")
			return
		case transactions.StatusNotEqual:
			h.internalError(ctx, "Senha incorreta.", res.ParseToError())
			return
		default:
			h.internalError(ctx, "Falha ao checar se o usuario existe. ", res.ParseToError())
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

	stringifiedToken, err := token.SignedString([]byte(h.env.JwtSecret))
	if err != nil {
		h.internalError(ctx, " Falha ao validar o usuario. ", err)
	}

	ctx.SetCookie(
		"auth",
		stringifiedToken,
		3600,
		"/",
		h.env.FrontendUrl,
		false, // SHOULD BE TRUE IN HTTPS
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  user,
	})
}

func (h *Handlers) SignoutHandler(ctx *gin.Context) {
	ctx.SetCookie(
		"auth",
		"",
		-1,
		"/",
		h.env.FrontendUrl,
		false,
		true,
	)

	ctx.Redirect(http.StatusFound, "/auth")
}

func (h *Handlers) CreateUserHandler(ctx *gin.Context) {
	var user *models.Users = &models.Users{}

	// Only admins can create other users so you can use admin.InstitutionId to attribute created users ids

	var request models.CreateUsers
	if h.BadJsonRequest(ctx, ctx.ShouldBindJSON(&request)) {
		return
	}

	_, res := h.operations.SelectUserByEmail(request.Email)
	if res.Status != transactions.StatusNotFound {
		if res.Status != transactions.StatusSuccess {
			h.internalError(ctx, "Falha ao checar se o email já está em uso.", res.ParseToError())
			return
		}
		h.clientError(ctx, "O e-mail já está em uso.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		h.internalError(ctx, "Falha ao gerar o usuario.", err)
		return
	}

	userRequest := &models.CreateUsers{
		Name:          request.Name,
		Email:         request.Email,
		Password:      string(hashedPassword),
		Role:          request.Role,
		InstitutionId: user.InstitutionId,
	}

	user, detailedError := h.operations.InsertUser(userRequest)
	if detailedError != nil {
		h.internalError(ctx, "Falha ao criar o usuario.", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  user,
		"error": nil,
	})
}

func (h *Handlers) GetInstitution(ctx *gin.Context) {
	unsignedUser, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": nil, "error": "No authentication data  found for the user"})
		return
	}

	user, ok := unsignedUser.(models.Users)
	if !ok {
		h.internalError(ctx, "Failed to get user", nil)
		return
	}

	institution, res := h.operations.SelectInstitutionById(user.InstitutionId)
	if res.Status == transactions.StatusNotFound {
		h.clientError(ctx, "Nenhuma instituição encontrada")
		return
	} else if res != nil {
		h.internalError(ctx, "Falha ao buscar as instituições", res.ParseToError())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  institution,
		"error": nil,
	})
}

func (h *Handlers) GetUsersByInstitutionId(ctx *gin.Context) {
	var request SelectByInstitutionRequest
	if h.BadJsonRequest(ctx, ctx.ShouldBindBodyWithJSON(&request)) {
		return
	}

	users, res := h.operations.SelectUsersByInstitutionId(request.InstitutionId)
	if res.Status == transactions.StatusNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": "Nenhum usuário encontrado.",
		})
		return
	} else if res != nil {
		h.internalError(ctx, "Falha ao procurar usuários", res.ParseToError())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  users,
		"error": nil,
	})
}

/*
	Refactor the whole cookie logic
	The cookie store the institution of the user so I can use it later to check certain data
	Create a cookie package that
		-> stores cookie names for fast use with no name mismatch
		-> handle basic cookie shit like auth and validations
*/
