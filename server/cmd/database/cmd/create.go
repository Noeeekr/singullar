package cmd

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common/configs"
	"github.com/Noeeekr/singullar/server/internal/database"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var createCommand *cobra.Command = &cobra.Command{
	Use:   "create [--email EMAIL] [--name NAME] [--password PASSWORD] [--environment file]",
	Short: "Create allows inserting institutions into the database via command-line-interface (CLI).",
	Long: fmt.Sprint(
		"-n --name: The name of the institution to be created.",
		"-e --environment: The path to the file containig the connection string. It must be in a variable called POSTGRES_CONNECTION_STRING.",
	),
	Run: func(cmd *cobra.Command, args []string) {
		// FLAGS
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		path, _ := cmd.Flags().GetString("environment")
		password, _ := cmd.Flags().GetString("password")

		if res := configs.Parse(path); res != nil {
			fmt.Println(res.ParseToError().Error())
			return
		}

		var env database.PostgresEnvironment
		if res := configs.Scan(&env); res != nil {
			fmt.Println(res.ParseToError().Error())
			return
		}

		// CONNECT TO DATABASE
		db, err := database.Connect(env.POSTGRES_CONNECTION_STRING)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// HASH PASSWORD
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		ops := operations.New(db)
		// CHECK IF USER EXISTS
		_, res := ops.SelectUserByEmail(email)
		if res == nil {
			fmt.Println("User with specified email already exists, please choose other.")
			return
		} else if res.Status != transactions.StatusNotFound {
			fmt.Println(res.Description)
			return
		}

		// CREATE INSTITUTION
		user, res := ops.InsertInstitution(name, email, string(pwd))
		if res != nil {
			fmt.Println(res.ParseToError().Error())
			return
		}

		// PRINT USER INFORMATION
		fmt.Println("Create new institution.")
		fmt.Println("Access it using the following user:")
		fmt.Println("Institution: ", name)
		fmt.Println("Name: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("Password: ", password)
	},
}

/*
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*?"

func main() {
	var environmentFilePath *string = nil
	var email *string = nil

	flag.StringVar(environmentFilePath, "env-path", "./config/dev.env", "sets the absolute enviroment path for server setup.")
	flag.StringVar(email, "email", "", "Create a institution and insert an admin user.")
	flag.Parse()

	err := configs.Parse(*environmentFilePath)
	if err != nil {
		fmt.Println("Found error while configuring env")
		return
	}

	if email != nil && *email == "" {
		fmt.Println("Invalid e-mail")
		return
	}

	logs.Info.Println("Generating password")
	strongPassword := make([]byte, 52)

	for i := range strongPassword {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			fmt.Println(err)
			return
		}
		strongPassword[i] = chars[num.Int64()]
	}

	logs.Info.Println("Hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword(strongPassword, 10)
	if err != nil {
		fmt.Printlnf("Failed to hash password %s", err.Error())
		return
	}

	logs.Info.Println("Connecting to database")

	// DATABASE SETUP
	db, err := gorm.Open(postgres.Open(cfg.DatabaseWR_ConnectString))
	if err != nil {
		fmt.Println(err)
		return
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback on panic
		}
	}()

	logs.Info.Println("Creating Institution")

	inst := &models.Institutions{
		Name:          "Nome não definido",
		ProfileImgUrl: "./assets/defaultpfp.jpg",
	}

	if err := tx.Model(&models.Institutions{}).Create(&inst).Error; err != nil {
		tx.Rollback()
		fmt.Printlnf("Fail in transaction: %s", err.Error())
		return
	}

	logs.Info.Println("Creating administrator user")

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
		fmt.Printlnf("Fail in transaction: %s", err.Error())
		return
	}

	if err := tx.Model(&user).Association("Institutions").Append(&inst); err != nil {
		tx.Rollback()
		fmt.Printlnf("Fail while trying to oficialize transaction: %s", err.Error())
		return
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Printlnf("Fail while trying to oficialize transaction: %s", err.Error())
		return
	}

	logs.Info.Println("Success, created an institution and admin user..")
	logs.Info.Printf("E-mail: %s", *email)
	logs.Info.Printf("Password: %s", string(strongPassword))
	logs.Info.Println("")
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
