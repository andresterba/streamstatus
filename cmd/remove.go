package cmd

import (
	"errors"

	"github.com/andresterba/streamstatus/internal"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove existing streamer",
	Long: `Remove streamer with name.

streamstatus remove [name] -> remove streamer from config.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkRemoveArgs(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		internal.RemoveExistingStreamer(args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func checkRemoveArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("Please provide name of streamer that should be deleted.")
	}
	return nil
}
