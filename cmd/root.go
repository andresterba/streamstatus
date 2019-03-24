package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "streamstatus",
	Short: "cli tool to check if your favorite twitch streamers are online",
	Long: `cli tool to check if your favorite twitch streamers are online.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

