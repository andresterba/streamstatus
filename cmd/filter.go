package cmd

import (
	"errors"
	"fmt"
	"github.com/andresterba/streamstatus/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
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
		showStreamersOfCategory(args[0])
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

func showStreamersOfCategory(categoryFromUser string) {

	streamers, err := internal.ReadStreamersFromFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	filteredStreamers, err := internal.FilterForCategory(streamers, categoryFromUser)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Println("Your streams are currently :")

	for i, streamer := range filteredStreamers {

		internal.UpdateStreamerStatus(&streamer)
		fmt.Printf("[%2d] %-16s %-7s", i, streamer.Name, streamer.Category)
		if streamer.Status == "offline" {
			color.Red(streamer.Status)
		} else {
			color.Green(streamer.Status)
		}
	}
}
