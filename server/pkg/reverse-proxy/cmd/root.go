package cmd

import (
	"fmt"
	"os"

	"github.com/Noeeekr/singullar/server/pkg/common/logs"
	"github.com/Noeeekr/singullar/server/pkg/reverse-proxy/proxy"
	"github.com/spf13/cobra"
)

// First argument's the path of the
var rootCmd = &cobra.Command{
	Use:   "singsrv",
	Short: "Singullar Reverse Proxy (SRV) handles starting other services for Singullar platform.",
	Long: `Singullar Reverse Proxy (SRV) handles starting other services for Singullar platform.
			You need to configure the execution in a proxy.json and pass its folder path in the first parameter.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Use --help for a detailed help message")
			fmt.Println("Please provide the path to the folder containing the proxy.json file")
			return
		}

		logs.Info.Println("Parsing json configuration file..")
		ReverseProxyConfig, err := proxy.ParseReverseProxyConfig(args[0])
		if err != nil {
			logs.Info.Println("Failed to parse json configuration file.")
			logs.Error.Println(err.Error())
			return
		}
		logs.Info.Println("Json configuration file parsed successfully.")
		logs.Info.Println("Starting reverse proxy.")

		ReverseProxy := proxy.NewReverseProxy()
		ReverseProxy.StartWith(ReverseProxyConfig)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
