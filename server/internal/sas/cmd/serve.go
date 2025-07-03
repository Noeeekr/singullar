package cmd

import (
	"github.com/Noeeekr/singullar/server/common/logs"
	"github.com/Noeeekr/singullar/server/internal/sas/server"
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

		server := server.New()
		if !debug {
			server.DisableDebug()
		}

		if err := server.ServeFolder(args[0], port); err != nil {
			logs.Error.Println(err.Error())
		}
	},
}
