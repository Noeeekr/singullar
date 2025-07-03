package cmd

import (
	"os"

	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/server/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve --port [PORT] --debug [FALSE] [FOLDER]",
	Short: "Serves the files inside the folder specified in the first argument after the flags.",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		port, _ := cmd.Flags().GetString("port")

		if !debug {
			os.Setenv("GIN_MODE", "release")
		}
		os.Setenv("PORT", port)
		os.Setenv("FOLDER", args[0])

		server := server.New()
		if err := server.ServeFolder(); err != nil {
			logs.Error.Println(err.Error())
		}
	},
}
