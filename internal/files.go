package internal

import (
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

type Streamer struct {
	Name         string
	Category     string
	Status       string
	CurrentTitle string
}

func ReadStreamersFromFile() ([]string, []Streamer) {
	inputFile := GetUserHomeDir()
	var StreamersArray []string
	var streamers []Streamer

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	for _, currentLine := range lines {

		if len(currentLine) == 0 {
		} else {
			StreamersArray = append(StreamersArray, currentLine)
			s := strings.Fields(currentLine)
			streamers = append(streamers, Streamer{s[0], s[1], "offline", "currently coding on streamstatus"})
		}
	}

	return StreamersArray, streamers
}

func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.streamstatus"
}
