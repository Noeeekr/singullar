package cmd

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/configs"
	"github.com/Noeeekr/singullar/server/internal/database"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/spf13/cobra"
)

var createCommand *cobra.Command = &cobra.Command{
	Use:   "create [--email EMAIL] [--environment file]",
	Short: "Create allows inserting institutions into the database via command-line-interface (CLI).",
	Long: fmt.Sprint(
		"-n --name: The name of the institution to be created.",
		"-e --environment: The path to the file containig the connection string. It must be in a variable called POSTGRES_CONNECTION_STRING.",
	),
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		path, err := cmd.Flags().GetString("environment")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if email == "" {
			fmt.Println("Name flag is required. Usage: --email='email'")
		}

		if err := configs.Parse(path); err != nil {
			fmt.Println(err.Status, err.Description)
			return
		}

		var env database.PostgresEnvironment
		if err := configs.Scan(&env); err != nil {
			fmt.Println(err.ParseToError())
			return
		}

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

		ops := operations.New(db)
		user, res := ops.InsertInstitution("administrator", email, common.GenerateRandomStrings(10))
		if res != nil {
			fmt.Println(res.ParseToError())
			return
		}

		fmt.Println("Create new institution.")
		fmt.Println("Access it using the following user:")
		fmt.Println("Name: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("Password: ", user.Password)
	},
}

func init() {
	createCommand.Flags().StringP("email", "n", "", "The name of the institution to be created")
	createCommand.Flags().StringP("environment", "e", "", "The path to the file containing the connection string.")
	rootCmd.AddCommand(createCommand)
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
		logs.Error.Fatal("Found error while configuring env")
	}

	if email != nil && *email == "" {
		logs.Error.Fatal("Invalid e-mail")
	}

	logs.Info.Println("Generating password")
	strongPassword := make([]byte, 52)

	for i := range strongPassword {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			logs.Error.Fatal(err)
		}
		strongPassword[i] = chars[num.Int64()]
	}

	logs.Info.Println("Hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword(strongPassword, 10)
	if err != nil {
		logs.Error.Fatalf("Failed to hash password %s", err.Error())
	}

	logs.Info.Println("Connecting to database")

	// DATABASE SETUP
	db, err := gorm.Open(postgres.Open(cfg.DatabaseWR_ConnectString))
	if err != nil {
		logs.Error.Fatal(err)
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
		logs.Error.Fatalf("Fail in transaction: %s", err.Error())
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
		logs.Error.Fatalf("Fail in transaction: %s", err.Error())
	}

	if err := tx.Model(&user).Association("Institutions").Append(&inst); err != nil {
		tx.Rollback()
		logs.Error.Fatalf("Fail while trying to oficialize transaction: %s", err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		logs.Error.Fatalf("Fail while trying to oficialize transaction: %s", err.Error())
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
