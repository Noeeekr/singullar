package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/api/types"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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
	if h.HandleBadJsonRequest(ctx, &request) {
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
