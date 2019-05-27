package internal

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type ApiStreamer struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	OfflineImageUrl string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
}

type Response struct {
	Data []ApiStreamer `json:"data"`
}

type Status struct {
	Stream string `json:"stream"`
	Links  struct {
		Self    string `json:"self"`
		Channel string `json:"channel"`
	} `json:"_links"`
}

func GetUserId(username string) string {

	baseUrl := "https://api.twitch.tv/helix/users?login="

	url := baseUrl + username

	resp, body, errs := gorequest.New().Get(url).
		Set("Client-ID", "qq88eegqplf3q8a5itj1qt3ohvhjo4").
		End()

	if errs != nil {
		fmt.Println("Resp:", resp)
		fmt.Println("Body:", body)
		fmt.Println("Err:", errs)
	}

	var s = new(Response)
	json.Unmarshal([]byte(body), &s)

	return s.Data[0].ID

}

func GetStreamStatus(userid string) (string, string) {

	baseUrl := "https://api.twitch.tv/kraken/streams/"

	url := baseUrl + userid

	resp, body, errs := gorequest.New().Get(url).
		Set("Client-ID", "qq88eegqplf3q8a5itj1qt3ohvhjo4").
		Set("Accept", "application/vnd.twitchtv.v5+json").
		End()

	if errs != nil {
		fmt.Println("Resp:", resp)
		fmt.Println("Body:", body)
		fmt.Println("Err:", errs)
	}

	var twitchAPIResponse map[string]interface{}

	if err := json.Unmarshal([]byte(body), &twitchAPIResponse); err != nil {
		panic(err)
	}

	streamStatus := twitchAPIResponse["stream"]

	//If stream is offline return offline and no current title
	if streamStatus == nil {
		return "offline", ""
	}

	streamerJSON := (twitchAPIResponse["stream"].(map[string]interface{}))
	channelJSON := (streamerJSON["channel"].(map[string]interface{}))
	currentStreamTitle := (channelJSON["status"].(string))

	return "online", currentStreamTitle
}
