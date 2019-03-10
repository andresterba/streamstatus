package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Streamer struct {
	ID                string `json:"id"`
	Login             string `json:"login"`
	Display_name      string `json:"display_name"`
	Type              string `json:"type"`
	Broadcaster_type  string `json:"broadcaster_type"`
	Description       string `json:"description"`
	Profile_image_url string `json:"profile_image_url"`
	Offline_image_url string `json:"offline_image_url"`
	View_count        int    `json:"view_count"`
}

type Response struct {
	Data []Streamer `json:"data"`
}

type Status struct {
	Stream string `json:"stream"`
	Links  struct {
		Self    string `json:"self"`
		Channel string `json:"channel"`
	} `json:"_links"`
}

type Foo struct {
	X map[string]interface{} `json:"-"` // Rest of the fields should go here.
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

func GetStreamStatus(userid string) string {

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

	f := Foo{}
	if err := json.Unmarshal([]byte(body), &f.X); err != nil {
		panic(err)
	}

	stream_status := f.X["stream"]

	if stream_status == nil {
		return "offline"

	} else {
		return "online"

	}
}
