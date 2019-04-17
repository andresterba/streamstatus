package internal

import (
	"errors"
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

func ReadStreamersFromFile() ([]Streamer, error) {
	inputFile := GetUserHomeDir()
	var StreamersArray []string
	var streamers []Streamer

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return []Streamer{}, errors.New("Can't open streamstatus config. Please add .streamstatus to your home directory.")
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

	return streamers, nil
}

func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.streamstatus"
}
