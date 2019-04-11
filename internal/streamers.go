package internal

import ()

func UpdateStreamerStatus(streamer *Streamer) {
	streamerUserID := GetUserId(streamer.Name)
	status := GetStreamStatus(streamerUserID)
	streamer.Status = status
}
