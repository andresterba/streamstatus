package cmd

import (
	"fmt"
	"github.com/andresterba/streamstatus/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
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
	streamers, err := internal.ReadStreamersFromFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Println("Your streams are currently :")

	for i, streamer := range streamers {

		internal.UpdateStreamerStatus(&streamer)
		fmt.Printf("[%2d] %-16s %-7s", i, streamer.Name, streamer.Category)
		if streamer.Status == "offline" {
			color.Red(streamer.Status)
		} else {
			color.Green(streamer.Status)
			fmt.Println(streamer.CurrentTitle)
		}
	}
}
