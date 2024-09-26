package server

import (
	"errors"
	"os"
	"path/filepath"
	"runtime/debug"

	"fmt"
	"log"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/noeeekr/sch-server/config"
	"github.com/noeeekr/sch-server/pkg/forms"
	"github.com/noeeekr/sch-server/pkg/models"
	pqsql "github.com/noeeekr/sch-server/pkg/models/pgsql"
	"github.com/noeeekr/sch-server/pkg/paths"
)

type RouterHandlers struct {
	LogInfo *log.Logger
	LogErr  *log.Logger
	users   *pqsql.UserModel
	env     *config.Configuration
}

// HTTP ERROR RESPONSE CODE REPLIERS

// Internal server errors sends a json as response with InternalServerError response code.
//
// It accepts a main "errType" and insert the error object inside it to be parsed into JSON for frontend.
// It also accepts an err for debug porpuses.
func (h *RouterHandlers) internalServerErr(ctx *gin.Context, message string, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	h.LogErr.Output(2, trace)

	// err types: validation and default
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
		"data":  nil,
	})
}

func (h *RouterHandlers) clientError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": message,
		"data":  nil,
	})
}

// STATICS RELATED HANDLERS

// Serves the static files located in environment variable "$STATICS_PATH".
//
// In case the file doesn't exist, serves environment variable "$STATICS_PATH/index.html",
// so react-router-dom (frontend) can manage the requests to not found page (404).
func (h *RouterHandlers) StaticsHandler(ctx *gin.Context) {
	path := filepath.Join(paths.Root, h.env.StaticsPath, ctx.Request.URL.Path)

	_, err := os.Stat(path)
	if err != nil {
		http.ServeFile(ctx.Writer, ctx.Request, filepath.Join(paths.Root, h.env.StaticsPath, "index.html"))
		return
	}

	http.ServeFile(ctx.Writer, ctx.Request, path)
}

// AUTH RELATED HANDLERS

// Checks if user email and password matches in database.
//
// For architetural porpuses this function doesn't handle query and bcrypt errors
// instead, they're returned to be handled in the main function.
// Such errors include Gorm.ErrRecordNotFound and Bcrypt.ErrMismatchedHashAndPassword.
func (h *RouterHandlers) CreateUser(user models.Users) (*models.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	if err != nil {
		return &user, err
	}

	user.Password = string(hashedPassword)
	user.ProfileImgUrl = "images/userDefaultPic.png"

	newUser, err := h.users.Insert(&user)
	if err != nil {
		return &user, err
	}

	return newUser, nil
}
func (h *RouterHandlers) checkUserPassword(email string, password string) (id string, e error) {
	var result struct {
		Hashed_password string `gorm:"column:password"`
		Id              string `gorm:"column:id"`
	}

	err := h.users.DB.Model(&models.Users{}).Select("password, id").Where("email = ?", email).Scan(&result).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Hashed_password), []byte(password))
	if err != nil {
		return "", err
	}

	return result.Id, nil
}

func (h *RouterHandlers) Authenticate(ctx *gin.Context) {
	s := sessions.Default(ctx)
	auth := s.Get("auth")

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  auth,
	})
}

// Handles user sign-in and returns custom graceful JSON objects for client form ui.
func (h *RouterHandlers) SigninHandler(ctx *gin.Context) {
	var user models.Users

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		h.internalServerErr(ctx, "Falha ao tentar converter a requisição.", err)
		return
	}

	form := forms.New(&user)
	form.SetField("Email").IsValidEmail()
	form.SetField("Password").Length(8, 0)

	if form.IsValid() != nil {
		h.clientError(ctx, "Senha ou e-mail invalidos.")
		return
	}

	id, err := h.checkUserPassword(user.Email, user.Password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h.clientError(ctx, "Usuário não existe.")
		return
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		h.clientError(ctx, "As senhas não coincidem.")
		return
	} else if err != nil {
		h.internalServerErr(ctx, "Falha ao checar se o usuário existe.", err)
		return
	}

	s := sessions.Default(ctx)
	s.Set("auth", id)
	s.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  id,
	})
}

func (h *RouterHandlers) SignoutHandler(ctx *gin.Context) {
	sessions.Default(ctx).Delete("userId")

	http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusPermanentRedirect)
}

func (h *RouterHandlers) SignupHandler(ctx *gin.Context) {
	// Validate user
	var user models.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.internalServerErr(ctx, "Falha ao tentar converter a requisição", err)
		return
	}

	form := forms.New(&user)
	form.Required("Name", "Surname", "Email", "Password")
	form.SetFields("Name", "Surname").Length(6, 0)
	form.SetField("Email").IsValidEmail()
	form.SetField("Password").Length(8, 0)

	if err := form.IsValid(); err != nil {
		h.clientError(ctx, err.Error())
		return
	}

	// Check if user already exists
	_, exists, err := h.users.GetByEmail(user.Email)
	if err != nil {
		h.internalServerErr(ctx, "Falha ao checar se o email já está em uso.", err)
		return
	} else if exists {
		h.clientError(ctx, "O e-mail já está em uso.")
		return
	}

	// Create new user
	newUser, err := h.CreateUser(user)

	if err != nil {
		h.internalServerErr(ctx, "Falha ao criar o usuário.", err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data":  newUser,
		"error": nil,
	})
}
