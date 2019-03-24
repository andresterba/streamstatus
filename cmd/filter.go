package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"github.com/andresterba/streamstatus/internal"
	"github.com/fatih/color"
	"os"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
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

func checkCategory(args []string) error{
	if len(args) < 1 {
		return errors.New("Please provide a valid category to filter for")
	}
	return nil
}

func showStreamersOfCategory(categoryFromUser string) {

	streamers := internal.ReadStreamersFromFile()

	if len(streamers) == 0 {
		fmt.Println("please add a config file")
		os.Exit(0)
	}
	var filteredStreamers []string

	for _, streamer := range streamers {

		s := strings.Split(streamer, " ")
		_, category := s[0], s[1]

		if category == categoryFromUser {
			filteredStreamers = append(filteredStreamers, streamer)
		}
	}

	if len(filteredStreamers) != 0 {
		fmt.Println("Your streams are currently :")

		for i, streamer := range filteredStreamers {
			s := strings.Split(streamer, " ")
			name, category := s[0], s[1]
			id := internal.GetUserId(name)
			status := internal.GetStreamStatus(id)

			fmt.Printf("[%2d] %-16s %-7s", i, name, category)

			if status == "offline" {
				color.Red(status)

			} else {
				color.Green(status)

			}

		}
	} else {
		fmt.Println("Seems like you haven't added any streamers to category " + categoryFromUser + " yet!")
	}
}