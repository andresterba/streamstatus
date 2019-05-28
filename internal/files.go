package internal

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

// Streamer represents a twitch streamer with propertys that will be displayed to the user.
type Streamer struct {
	Name         string
	Category     string
	Status       string
	CurrentTitle string
}

// ReadStreamersFromFile reads the .streamstatus config and loads the streamers.
func ReadStreamersFromFile() ([]Streamer, error) {
	inputFile := getConfigPath()
	var StreamersArray []string
	var streamers []Streamer

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return []Streamer{}, errors.New("can't open streamstatus config,please add .streamstatus to your home directory")
	}
	lines := strings.Split(string(b), "\n")

	for _, currentLine := range lines {

		if len(currentLine) == 0 {
		} else {
			StreamersArray = append(StreamersArray, currentLine)
			s := strings.Fields(currentLine)
			streamerName := s[0]
			streamerCategory := strings.ToLower(s[1])
			streamers = append(streamers, Streamer{streamerName, streamerCategory, "offline", "currently coding on streamstatus"})
		}
	}

	return streamers, nil
}

func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.streamstatus"
}

func AddNewStreamer(args []string) {
	streamer := args[0]
	category := args[1]

	// If the file doesn't exist create it. If it exists append to the file.
	f, err := os.OpenFile(getConfigPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(streamer + " " + category + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
