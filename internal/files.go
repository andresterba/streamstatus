package internal

import (
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

func ReadStreamersFromFile() []string {
	inputFile := GetUserHomeDir()
	var StreamersArray []string

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	for _, currentLine := range lines {

		if len(currentLine) == 0 {
		} else {

			StreamersArray = append(StreamersArray, currentLine)
		}
	}

	return StreamersArray

}

func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.streamstatus"

}
