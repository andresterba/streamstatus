package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long:  `Show current version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("streamstatus 0.7.5")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
