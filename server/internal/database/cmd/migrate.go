package cmd

import (
	"github.com/Noeeekr/singullar/server/internal/database/cmd/migrate"

	"github.com/spf13/cobra"
)

var migrateCmd *cobra.Command = &cobra.Command{
	Use:   "migrate [command]",
	Short: "Allows migrating tables to Postgresql database.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	migrateCmd.AddCommand(migrate.EnvironmentCmd)
	migrateCmd.AddCommand(migrate.RelationsCmd)

	rootCmd.AddCommand(migrateCmd)
}
