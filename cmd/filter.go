package cmd

import (
	"errors"

	"github.com/andresterba/streamstatus/internal"
	"github.com/spf13/cobra"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter [category]",
	Short: "Filter for specific category",
	Long: `Filter for specific category. For example:

streamstatus filter code -> shows all streamers in category code.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkCategory(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		internal.ShowStreamersOfCategoryFiltered(args[0])
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
}

func checkCategory(args []string) error {
	if len(args) < 1 {
		return errors.New("Please provide a valid category to filter for")
	}
	return nil
}
