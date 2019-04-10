package internal

import (
	"fmt"
)

// func UpdateStatus(streamers *Streamers) {
// 	for _, streamer := range streamers.Streamer {
// 		// streamer.updateStreamerStatus()
// 		//status := updateStreamerStatus(streamer)
//
// 		streamer.Status = "wasd"
//
// 		fmt.Println(streamer)
// 	}
// }

func UpdateStreamerStatus(streamer *Streamer) {
	streamerUserID := GetUserId(streamer.Name)
	status := GetStreamStatus(streamerUserID)

	fmt.Printf("%s %s \n", streamerUserID, status)

	streamer.Status = status
}
