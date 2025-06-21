package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/Noeeekr/singullar/server/api/pkg/paths"

	"github.com/gin-gonic/gin"

	// COOKIE BASED AUTH
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Noeeekr/singullar/server/api/config"
	"github.com/Noeeekr/singullar/server/api/internal/core/models"
	pqsql "github.com/Noeeekr/singullar/server/api/internal/core/models/pgsql"
	"gorm.io/gorm"
)

type RouterHandlers struct {
	LogInfo      *log.Logger
	LogErr       *log.Logger
	users        *pqsql.UserModel
	institutions *pqsql.InstitutionModel
	env          *config.Configuration
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

	h.LogInfo.Println("STATIC HANDLER CALLED")
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

func (h *RouterHandlers) Authenticate(ctx *gin.Context) {
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
			return nil, errors.New("failed to parse token")
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

func (h *RouterHandlers) checkUserPassword(email string, password string) (user *models.Users, e error) {
	var _user models.Users

	err := h.users.DB.Model(&models.Users{}).Where("email = ?", email).First(&_user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(_user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	_user.Password = ""
	return &_user, nil
}

// Handles user sign-in and returns custom graceful JSON objects for client form ui.
type AuthClaims struct {
	User models.Users
	jwt.RegisteredClaims
}

func (h *RouterHandlers) SigninHandler(ctx *gin.Context) {
	var user models.Signin

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		h.clientError(ctx, "Campos em formato incorreto.")
		return
	}
	usr, err := h.checkUserPassword(user.Email, user.Password)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.clientError(ctx, "Usuario não existe")
			return
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			h.internalServerErr(ctx, "Senha incorreta. ", err)
			return
		} else {
			h.internalServerErr(ctx, "Falha ao checar se o usuario existe. ", err)
			return
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		User: *usr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    fmt.Sprintf("%d", usr.ID),
		},
	})

	stringifiedToken, err := token.SignedString([]byte(h.env.JwtSecret))
	if err != nil {
		h.internalServerErr(ctx, " Falha ao validar o usuario. ", err)
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
		"data":  usr,
	})
}

func (h *RouterHandlers) SignoutHandler(ctx *gin.Context) {
	ctx.SetCookie(
		"auth",
		"",
		0,
		"/",
		h.env.FrontendUrl,
		false,
		true,
	)

	ctx.Redirect(http.StatusFound, "/auth")
}

// SIGNUP RELATED HANDLERS
func createUser(user models.CreateUsers) (*models.Users, error) {
	formattedUser := models.Users{
		CreateUsers: models.CreateUsers{
			Name:          user.Name,
			Email:         user.Email,
			Password:      user.Password,
			Role:          user.Role,
			InstitutionId: user.InstitutionId,
		},
		ProfileImgUrl: "images/userDefaultPic.png",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formattedUser.Password), 10)

	if err != nil {
		return &formattedUser, err
	}

	formattedUser.Password = string(hashedPassword)

	return &formattedUser, nil
}

func (h *RouterHandlers) UserSignupHandler(ctx *gin.Context) {
	var user models.CreateUsers

	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.internalServerErr(ctx, "Dados em formato incorreto", err)
		return
	}

	if user.Role != models.Student && user.Role != models.Teacher && user.Role != models.Supervisor {
		h.clientError(ctx, "Cargo de usuário invalido.")
		return
	}

	// Check if user already exists
	var exists bool
	var err error

	_, exists, err = h.users.GetByEmail(user.Email)
	if err != nil {
		h.internalServerErr(ctx, "Falha ao checar se o email já está em uso.", err)
		return
	} else if exists {
		h.clientError(ctx, "O e-mail já está em uso.")
		return
	}

	// Create new user
	newUser, err := createUser(user)
	if err != nil {
		h.internalServerErr(ctx, "Falha ao gerar o usuario.", err)
		return
	}

	resUser, err := h.users.Insert(newUser)
	if err != nil {
		h.internalServerErr(ctx, "Falha ao criar o usuario.", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  resUser,
		"error": nil,
	})
}

func (h *RouterHandlers) GetInstitutions(ctx *gin.Context) {
	user, exists := ctx.Get("User")
	if !exists {
		h.internalServerErr(ctx, "Falha ao encontrar os dados do usuário", errors.New(" Failed to get user via from cookie, from context passed by auth middleware"))
		return
	}

	usr, ok := user.(models.Users)
	if !ok {
		h.internalServerErr(ctx, "Dados de usuário no formato errado", errors.New(" Failed to type assert the given context"))
		return
	}

	Institutions, exists, err := h.institutions.FindByUserId(usr.ID)
	if err != nil {
		h.internalServerErr(ctx, "Falha ao buscar as instituições", err)
		return
	} else if !exists {
		h.clientError(ctx, "Nenhuma instituição encontrada")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  Institutions,
		"error": nil,
	})
}

type GetStudentsRequest struct {
	Students      []models.CreateStudent `json:"students" binding:"required"`
	InstitutionId uint                   `json:"institutionId" binding:"required"`
}

func (h *RouterHandlers) GetStudents(ctx *gin.Context) {
	var Req GetStudentsRequest

	if err := ctx.ShouldBindJSON(&Req); err != nil {
		h.internalServerErr(ctx, "Dados em formato incorreto", err)
		return
	}

	var StudentsNames []string

	for i := 0; i < len(Req.Students); i++ {
		StudentsNames = append(StudentsNames, Req.Students[i].Name)
	}

	usrs, exist, err := h.users.QueryByNames(StudentsNames, Req.InstitutionId)
	if err != nil {
		h.internalServerErr(ctx, "Dados em formato incorreto", err)
		return
	}

	if !exist {
		ctx.JSON(201, gin.H{
			"data":  nil,
			"error": "Nenhum usuário encontrado.",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"data":  usrs,
		"error": nil,
	})
}
