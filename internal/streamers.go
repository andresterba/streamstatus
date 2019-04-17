package internal

import (
	"errors"
)

func UpdateStreamerStatus(streamer *Streamer) {
	streamerUserID := GetUserId(streamer.Name)
	status := GetStreamStatus(streamerUserID)
	streamer.Status = status
}

func FilterForCategory(streamers []Streamer, category string) ([]Streamer, error) {
	var filteredStreamers []Streamer

	for _, streamer := range streamers {
		if streamer.Category == category {
			filteredStreamers = append(filteredStreamers, streamer)
		}
	}

	if len(filteredStreamers) == 0 {
		return filteredStreamers, errors.New("Can't find any streamer in this category. Please add a new streamer to this category.")
	}

	return filteredStreamers, nil

}
