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
