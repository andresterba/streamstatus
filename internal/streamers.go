package internal

import (
	"fmt"
)

func UpdateStatus(streamers *Streamers) {
	for _, streamer := range streamers.Streamer {
		// streamer.updateStreamerStatus()
		status := updateStreamerStatus(streamer)

		s := &streamer

		s.Status = status
	}
}

func updateStreamerStatus(streamer Streamer) string {
	streamerUserID := GetUserId(streamer.Name)
	status := GetStreamStatus(streamerUserID)

	fmt.Printf("%s %s \n", streamerUserID, status)

	return status
}
