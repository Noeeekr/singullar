package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "database",
	Short: "Database is a utility tool that allows inserting rows into singullar.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	createCommand.Flags().String("email", "", "The email of the user that will be given administrator role")
	createCommand.MarkFlagRequired("email")

	createCommand.Flags().StringP("environment", "e", "", "The path to the file containing the connection string.")
	createCommand.MarkFlagRequired("environment")

	createCommand.Flags().StringP("name", "n", "", "The name of the institution to be created.")
	createCommand.MarkFlagRequired("name")

	createCommand.Flags().StringP("password", "p", "", "The password of the administrator.")
	createCommand.MarkFlagRequired("password")

	rootCmd.AddCommand(createCommand)

	migrateCmd.Flags().StringP("environment", "e", "", "The path to the file containing the connection string.")
	migrateCmd.MarkFlagRequired("environment")

	rootCmd.AddCommand(migrateCmd)
}
