package internal

import (
	"fmt"
)

func (streamers Streamers) UpdateStatus() {
    for _, streamer := range streamers.Streamer {
		streamer.updateStreamerStatus()
		fmt.Printf("Test")
	}
}


func (streamer Streamer) updateStreamerStatus() {
	streamerUserID := GetUserId(streamer.Name)
	status := GetStreamStatus(streamerUserID)

	fmt.Printf(status)

	streamer.Status = status
}