package cmd

import (
	"errors"

	"github.com/andresterba/streamstatus/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name] [category]",
	Short: "Add new streamer",
	Long: `Add new streamer with name and category.

streamstatus add [name] [category] -> shows all streamers in category code.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAddArgs(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		internal.AddNewStreamer(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func checkAddArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("Please provide the name and category for streamer to be added")
	}
	return nil
}
