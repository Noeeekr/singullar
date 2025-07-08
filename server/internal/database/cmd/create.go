package cmd

import (
	"github.com/Noeeekr/singullar/server/internal/database/cmd/create"

	"github.com/spf13/cobra"
)

var createCommand *cobra.Command = &cobra.Command{
	Use:   "create [ institution ]",
	Short: "Create allows inserting institutions into the database via command-line-interface (CLI).",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	createCommand.AddCommand(create.InstitutionCmd)

	rootCmd.AddCommand(createCommand)
}
