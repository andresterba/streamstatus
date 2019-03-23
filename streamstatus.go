package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {

	streamers := readStreamersFromFile()

	if len(streamers) == 0 {
		fmt.Println("please add a config file")
		os.Exit(0)
	}

	args := os.Args[1:]

	if len(args) < 1 {
		showAllStreamers(streamers)
	} else {
		category := args[0]
		showStreamersOfCategory(streamers, category)
	}
}

func showAllStreamers(streamers []string) {
	fmt.Println("Your streams are currently :")

	for i, streamer := range streamers {

		s := strings.Split(streamer, " ")
		name, category := s[0], s[1]
		id := GetUserId(name)
		status := GetStreamStatus(id)

		fmt.Printf("[%2d] %-16s %-7s", i, name, category)

		if status == "offline" {
			color.Red(status)

		} else {
			color.Green(status)

		}

	}

}

func showStreamersOfCategory(streamers []string, categoryFromUser string) {

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
			id := GetUserId(name)
			status := GetStreamStatus(id)

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
