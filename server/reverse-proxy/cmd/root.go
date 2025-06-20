package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "srv",
	Short: "Singullar Reverse Proxy (SRV) handles starting other services for Singullar platform.",
	Long: `Singullar Reverse Proxy (SRV) handles starting other services for Singullar platform.
			You can start it and the main api. Additionally you can also start a static file server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Help message")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
