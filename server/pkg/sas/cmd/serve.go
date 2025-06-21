package cmd

import (
	"github.com/Noeeekr/singullar/server/pkg/common/logs"
	"github.com/Noeeekr/singullar/server/pkg/sas/server"
	"github.com/spf13/cobra"
)

var staticsfolder string
var port string
var debugMode bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a server that serves static files from a designated path on the designated port",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var errors []error

		for _, err := range errors {
			if err != nil {
				logs.Error.Println(err.Error())
				return
			}
		}

		err := server.New(staticsfolder, port, debugMode).ServeAndListen()
		if err != nil {
			logs.Error.Println(err.Error())
			return
		} else {
			logs.Info.Println("Server stopped successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&staticsfolder, "path", "p", "", "The absolute path of the folder that contains the assets to be served")
	serveCmd.MarkFlagRequired("path")

	serveCmd.Flags().StringVar(&port, "port", "", "The port the server will listen to")
	serveCmd.MarkFlagRequired("port")

	serveCmd.Flags().BoolVarP(&debugMode, "debug", "d", true, "Defines if server will start in debug mode. Defaults to true")
}
