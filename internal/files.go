package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

// Streamer represents a twitch streamer with properties that will be displayed to the user.
type Streamer struct {
	Name         string
	Category     string
	Status       string
	CurrentTitle string
}

// ReadStreamersFromFile reads the .streamstatus config and loads the streamers.
func ReadStreamersFromFile() ([]Streamer, error) {
	fileFile := getConfigPath()
	var StreamersArray []string
	var streamers []Streamer

	b, err := ioutil.ReadFile(fileFile)
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
			streamers = append(streamers, Streamer{streamerName, streamerCategory, "offline", "coding on streamstatus"})
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

// AddNewStreamer checks if streamer already exists and adds it if not.
func AddNewStreamer(streamerName string, category string) {
	streamerExists := checkIfStreamerExists(streamerName)
	if streamerExists {
		fmt.Println("Streamer " + streamerName + " already exists!")
		return
	}

	// If the file doesn't exist create it. If it exists append to the file.
	f, err := os.OpenFile(getConfigPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(streamerName + " " + category + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// RemoveExistingStreamer checks if streamer exists and deletes it from the config file.
func RemoveExistingStreamer(streamerName string) {
	streamerExists := checkIfStreamerExists(streamerName)
	if !streamerExists {
		fmt.Println("Streamer does not exist!")
		return
	}

	file, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, streamerName) {
			copy(lines[i:], lines[i+1:])
			lines[len(lines)-1] = ""
			lines = lines[:len(lines)-1]
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(getConfigPath(), []byte(output), 0600)
	if err != nil {
		log.Fatalln(err)
	}
}

func checkIfStreamerExists(streamerName string) bool {
	file, err := os.Open(getConfigPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()
		currentLineSplitted := strings.Split(currentLine, " ")
		currentStreamer := currentLineSplitted[0]

		if streamerName == currentStreamer {
			return true
		}
	}

	return false
}

// RenameCategory renames all occurrences of a category.
func RenameCategory(oldCategoryName string, newCategoryName string) {
	fmt.Println("Rename category " + oldCategoryName + " to " + newCategoryName + ".")

	file, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, oldCategoryName) {
			splittedLine := strings.Split(line, " ")
			streamerName := splittedLine[0]
			lines[i] = streamerName + " " + newCategoryName
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(getConfigPath(), []byte(output), 0600)
	if err != nil {
		log.Fatalln(err)
	}
}

// RenameStreamer renames an existing streamer.
func RenameStreamer(oldStreamerName string, newStreamerName string) {
	fmt.Println("Rename streamer " + oldStreamerName + " to " + newStreamerName + ".")
	streamerExists := checkIfStreamerExists(oldStreamerName)
	if !streamerExists {
		fmt.Println("Streamer does not exist!")
		return
	}

	file, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, oldStreamerName) {
			splittedLine := strings.Split(line, " ")
			categoryOfStreamer := splittedLine[1]
			lines[i] = newStreamerName + " " + categoryOfStreamer
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(getConfigPath(), []byte(output), 0600)
	if err != nil {
		log.Fatalln(err)
	}
}
