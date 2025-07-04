package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "api",
	Short: "Shows help for the api command line interface",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().BoolP("debug", "d", false, "Defines if the server should start in debug mode.")
	startCmd.Flags().StringP("port", "p", "80", "Defines the port the server will listen to.")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
