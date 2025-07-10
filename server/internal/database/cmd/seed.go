package cmd

import (
	"fmt"

	"github.com/Noeeekr/singullar/server/common/environment"
	"github.com/Noeeekr/singullar/server/internal/database/connections"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/operations"
	"github.com/Noeeekr/singullar/server/internal/database/seeder"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
	"github.com/spf13/cobra"
)

var seedCmd *cobra.Command = &cobra.Command{
	Use:   "seed [ -f ENVIRONMENT_FILES... ] [ -n QUANTITY ] [ -i ID ] { production | development }",
	Args:  cobra.MinimumNArgs(1),
	Short: "Seeds all tables with objects related to a specific institution defined by the id flag..",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		quantity, _ := cmd.Flags().GetInt("quantity")
		institutionId, _ := cmd.Flags().GetInt("id")
		mode := args[0]

		files, _ := cmd.Flags().GetStringArray("environmentFiles")
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

		// To make everything in a single transaction
		tx := transactions.NewManager(db).Start()
		ops := operations.New(db)

		createdUsers := seeder.CreateUserRequests(quantity, institutionId)
		users, tx := ops.InsertManyUsers(createdUsers...)
		if tx.Response != nil {
			fmt.Println(tx.Response.ParseToString())
			return
		}

		var teachers []*models.Users
		for _, user := range users {
			if user.Role == models.Teacher {
				teachers = append(teachers, user)
			}
		}
		var students []*models.Users
		for _, user := range users {
			if user.Role == models.Student {
				students = append(students, user)
			}
		}

		for _, teacher := range teachers {
			for _, student := range students {
				notifications := seeder.CreateNotificationRequests(4, teacher.Id, student.Id, student.Role)
				res := ops.InsertNotifications(tx, notifications...).Response
				if res != nil {
					fmt.Println(res.ParseToString())
					return
				}
			}
		}
		// Create users
		//   -> If teacher
		//       -> Create more ten users
		//			-> Create notifications for them
		//	 -> If student nothing
		//
		//
	},
}

func init() {
	seedCmd.Flags().StringArrayP("environmentFiles", "f", []string{}, "Defines environment files to parse the required environment variables.")

	seedCmd.Flags().IntP("quantity", "n", 50, "The amount of times to seed the table. Defaults to 50")

	seedCmd.Flags().IntP("id", "i", 0, "The id of the institution to seed")
	seedCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(seedCmd)
}
