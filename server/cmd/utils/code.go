package main

import (
	"flag"
	"math/big"

	"golang.org/x/crypto/bcrypt"

	"crypto/rand"

	configs "github.com/noeeekr/sch-server/config"
	logs "github.com/noeeekr/sch-server/internal/core/log"

	"github.com/noeeekr/sch-server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*?"

func main() {
	// FLAGS
	logs.InitLoggers()
	env_path := flag.String("env-path", "./config/dev.env", "sets the absolute enviroment path for server setup.")
	email := flag.String("email", "", "Create a institution and insert an admin user.")

	flag.Parse()

	// ENV CONFIG
	cfg, err := configs.NewConfig(*env_path)
	if err != nil {
		logs.LogErr.Fatal("Found error while configuring env")
	}

	if email != nil && *email == "" {
		logs.LogErr.Fatal("Invalid e-mail")
	}

	logs.LogInfo.Println("Generating password")
	strongPassword := make([]byte, 52)

	for i := range strongPassword {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			logs.LogErr.Fatal(err)
		}
		strongPassword[i] = chars[num.Int64()]
	}

	logs.LogInfo.Println("Hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword(strongPassword, 16)
	if err != nil {
		logs.LogErr.Fatalf("Failed to hash password %s", err.Error())
	}

	logs.LogInfo.Println("Connecting to database")

	// DATABASE SETUP
	db, err := gorm.Open(postgres.Open(cfg.DatabaseWR_ConnectString))
	if err != nil {
		logs.LogErr.Fatal(err)
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback on panic
		}
	}()

	logs.LogInfo.Println("Creating Institution")

	inst := &models.Institutions{
		Name:          "Nome não definido",
		ProfileImgUrl: "./assets/defaultpfp.jpg",
	}

	if err := tx.Model(&models.Institutions{}).Create(&inst).Error; err != nil {
		tx.Rollback()
		logs.LogErr.Fatalf("Fail in transaction: %s", err.Error())
	}

	logs.LogInfo.Println("Creating administrator user")

	user := &models.Users{
		CreateUsers: models.CreateUsers{
			Name:          "Administrador",
			Email:         *email,
			Password:      string(hashedPassword),
			Role:          "admin",
			InstitutionId: inst.ID,
		},
		Institutions:  []models.Institutions{*inst},
		ProfileImgUrl: "./assets/defaultpfp.jpg",
	}

	if err := tx.Model(&user).Create(user).Error; err != nil {
		tx.Rollback()
		logs.LogErr.Fatalf("Fail in transaction: %s", err.Error())
	}

	if err := tx.Model(&user).Association("Institutions").Append(&inst); err != nil {
		tx.Rollback()
		logs.LogErr.Fatalf("Fail while trying to oficialize transaction: %s", err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		logs.LogErr.Fatalf("Fail while trying to oficialize transaction: %s", err.Error())
	}

	logs.LogInfo.Println("Success, created an institution and admin user..")
	logs.LogInfo.Printf("E-mail: %s", *email)
	logs.LogInfo.Printf("Password: %s", string(strongPassword))
	logs.LogInfo.Println("")
}

/*func createInstitution(user models.CreateUsers) (*models.Institutions, error) {
	formattedUser := models.Institutions{
		CreateInstitutions: models.CreateInstitutions{
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		},
		Name:          user.Name,
		ProfileImgUrl: "images/userDefaultPic.png",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formattedUser.Password), 16)

	if err != nil {
		return &formattedUser, err
	}

	formattedUser.Password = string(hashedPassword)

	return &formattedUser, nil
}*/
