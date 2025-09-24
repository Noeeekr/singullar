package create

import (
	"errors"
	"fmt"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/util"
	"github.com/Noeeekr/singullar/server/util/environment"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var InstitutionCmd *cobra.Command = &cobra.Command{
	Use:   "institution --email [EMAIL] --password [PASSWORD] --name [NAME] [ -f ENVIRONMENT_FILES... ] { development | production }",
	Short: "Create a institution and an administrator user",
	Args:  cobra.MinimumNArgs(1),
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Required information for creating an institution
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		// Required configuration flags
		if args[0] == "production" {
			environment.Settings().SetApplicationMode(environment.PRODUCTION)
		} else {
			environment.Settings().SetApplicationMode(environment.DEVELOPMENT)
		}

		files, _ := cmd.Flags().GetStringArray("environmentFiles")
		if err := environment.Parse(files...); err != nil {
			util.Error.Fatal(err)
			return
		}

		// HASH PASSWORD
		psswd, err2 := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err2 != nil {
			fmt.Println(err2.Error())
			return
		}

		commiter, err := borm.Connect(models.EnvironmentDatabase)
		if err != nil {
			util.Error.Fatal(err)
		}
		ops := operations.New(commiter)

		// CHECK IF USER EXISTS
		if _, err := ops.SelectUserByEmail(email); err == nil {
			fmt.Println("User with specified email already exists, please choose other.")
			return
		} else if !errors.Is(err, borm.ErrNotFound) {
			util.Error.Fatal(err)
			return
		}

		// CREATE INSTITUTION
		institution := operations.CreateInstitutionRequest(name, string(psswd), email)

		err = ops.StartTransaction()
		if err != nil {
			util.Error.Fatal(err)
		}

		users, err := ops.InsertInstitutions(institution)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			for _, user := range users {
				fmt.Println("Created new institution.")
				fmt.Println("Access it using the following user:")
				fmt.Println("Institution: ", name)
				fmt.Println("Name: ", user.Name)
				fmt.Println("Email: ", user.Email)
				fmt.Println("Password: ", password)
			}
		}

		err = ops.CommitTransaction()
		if err != nil {
			fmt.Println(err)
			return
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
