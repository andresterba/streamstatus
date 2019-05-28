package cmd

import (
	"github.com/andresterba/streamstatus/internal"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all streamers and their current status",
	Long:  `Show all streamers and their current status.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.ShowAllStreamers()

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
