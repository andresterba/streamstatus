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

		err := internal.UpdateStreamerStatus(&streamer)

		if err != nil {
			fmt.Printf("Can not find user %s \n", streamer.Name)
			continue
		}

		if streamer.Status == "offline" {
			fmt.Printf("[%2d] %-16s %-7s", i, streamer.Name, streamer.Category)
			color.Red(streamer.Status)
		} else {
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Printf("[%2d] %-16s %-7s %-7s %s \n", i, streamer.Name, streamer.Category, green(streamer.Status), streamer.CurrentTitle)
		}
	}
}
