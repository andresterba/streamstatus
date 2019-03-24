package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/andresterba/streamstatus/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name] [category]",
	Short: "Add new streamer",
	Long: `Add new streamer with name and category.

streamstatus add [name] [category] -> shows all streamers in category code.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkArgs(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		addNewStreamer(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func checkArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("Please provide the name and category for streamer to be added")
	}
	return nil
}

func addNewStreamer(args []string) {
	streamer := args[0]
	category := args[1]

	// If the file doesn't exist create it. If it exists append to the file.
	f, err := os.OpenFile(internal.GetUserHomeDir(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(streamer + " " + category + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
