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
