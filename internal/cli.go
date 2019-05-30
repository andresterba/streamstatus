package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// ShowStreamersOfCategoryFiltered shows all streamers of a specified category with their current state.
func ShowStreamersOfCategoryFiltered(selectedCategory string) {

	streamers, err := ReadStreamersFromFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	filteredStreamers, err := FilterForCategory(streamers, selectedCategory)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Println("Your streams are currently :")

	for i, streamer := range filteredStreamers {

		err := UpdateStreamerStatus(&streamer)

		if err != nil {
			fmt.Printf("Can't find user %s \n", streamer.Name)
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

// ShowAllStreamers shows all streamers with their current state.
func ShowAllStreamers() {
	streamers, err := ReadStreamersFromFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Println("Your streams are currently :")

	for i, streamer := range streamers {

		err := UpdateStreamerStatus(&streamer)

		if err != nil {
			fmt.Printf("Can't find user %s \n", streamer.Name)
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
