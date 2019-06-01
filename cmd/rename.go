package cmd

import (
	"errors"
	"fmt"

	"github.com/andresterba/streamstatus/internal"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename category/streamer [oldname] [newname]",
	Short: "Rename existing streamer or category",
	Long: `Rename streamer or category.

streamstatus rename category/streamer [oldname] [newname]`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkRenameArgs(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "category":
			internal.RenameCategory(args[1], args[2])
		case "streamer":
			internal.RenameStreamer(args[1], args[2])
		default:
			fmt.Println("Select category or streamer.")
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}

func checkRenameArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("Please provide category|streamer flag and old/new name of what should be renamed.")
	}
	return nil
}
