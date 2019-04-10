package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/andresterba/streamstatus/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all streamers and their current status",
	Long:  `Show all streamers and their current status.`,
	Run: func(cmd *cobra.Command, args []string) {
		showAllStreamers()

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func showAllStreamers() {
	streamers, streamersArray := internal.ReadStreamersFromFile()

	if len(streamers) == 0 {
		fmt.Println("please add a config file")
		os.Exit(0)
	}

	fmt.Println("Your streams are currently :")

	for i, streamer := range streamers {
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

	//streamersStruct.UpdateStatus()
	//internal.UpdateStatus(&streamersArray)

	fmt.Println(streamersArray)

	for i, streamer := range streamersArray {

		internal.UpdateStreamerStatus(&streamer)
		fmt.Printf("[%2d] %-16s %-7s %-7s \n", i, streamer.Name, streamer.Category, streamer.Status)
	}
}
