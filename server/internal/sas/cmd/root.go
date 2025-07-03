package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sas [command]",
	Short: "Singullar Static Server (SINGSS) provides static files.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("port", "p", "80", "Defines the port the server will listen to. Defaults to 80")
	serveCmd.Flags().BoolP("debug", "d", false, "Defines if server will start in debug mode. Defaults to false")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
