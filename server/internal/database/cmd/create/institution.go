package create

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var InstitutionCmd *cobra.Command = &cobra.Command{
	Use:   "institution [ --email EMAIL ] [ --password PASSWORD ] [ -f ENVIRONMENT_FILES... ] { development | production }",
	Short: "Create a institution and an administrator user",
	Args:  cobra.MinimumNArgs(1),
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Required creation flags
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		// Required connection flags
		files, _ := cmd.Flags().GetStringArray("environmentFiles")
		mode := args[0]

		if res := environment.Parse(files...); res != nil {
			fmt.Println(res.ParseToString())
			return
		}

		db, res := connections.ConnectWithEnvironment(connections.ConnectionEnvironment(mode))
		if res != nil {
			fmt.Println(res.ParseToString())
			return
		}
		defer db.Close()

		err := db.Ping()
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
		if _, res := ops.SelectUserByEmail(email); res == nil {
			fmt.Println("User with specified email already exists, please choose other.")
			return
		} else if res.Status != common.StatusNotFound {
			fmt.Println(res.Description)
			return
		}

		// CREATE INSTITUTION
		if user, tx := ops.InsertInstitution(nil, name, email, string(pwd)); tx.Response != nil {
			fmt.Println(tx.Response.ParseToString())
			return
		} else {
			// PRINT USER INFORMATION
			fmt.Println("Create new institution.")
			fmt.Println("Access it using the following user:")
			fmt.Println("Institution: ", name)
			fmt.Println("Name: ", user.Name)
			fmt.Println("Email: ", user.Email)
			fmt.Println("Password: ", password)
		}
	},
}

func init() {
	InstitutionCmd.Flags().String("email", "", "The email of the user that will be given administrator role")
	InstitutionCmd.MarkFlagRequired("email")

	InstitutionCmd.Flags().String("password", "", "The password of the administrator.")
	InstitutionCmd.MarkFlagRequired("password")

	InstitutionCmd.Flags().String("name", "", "The name of the institution to be created.")
	InstitutionCmd.MarkFlagRequired("name")

	InstitutionCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "The path to the files containing the required environment variables.")
}
