package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

func readStreamersFromFile() []string {
	input_file := getUserHomeDir()
	var Streamers_array []string

	b, err := ioutil.ReadFile(input_file)
	if err != nil {
		//log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	for _, line := range lines {

		if len(line) == 0 {
		} else {

			Streamers_array = append(Streamers_array, line)
		}
	}

	return Streamers_array

}

func getUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.streamstatus"

}
