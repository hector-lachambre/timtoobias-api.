// 	Copyright (C) 2019 Hector Lachambre
//
// 	This file is part of huzlive-api.
//
//  Foobar is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  Foobar is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with Foobar.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitlab.com/timtoobias-projects/timtoobias-api/model"
)

// VideoContainer represents the video part of the api response
type VideoContainer struct {
	Videos   *Videos   `json:"datas"`
	DateSync time.Time `json:"dateSync"`
}

// Stream represents the stream content of the api response
type Stream struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

// StreamContainer represents the stream part of the api response
type StreamContainer struct {
	Stream   *Stream   `json:"datas"`
	DateSync time.Time `json:"dateSync"`
}

// Videos represents both channel last Video
type Videos struct {
	Main   Video `json:"main"`
	Second Video `json:"second"`
}

// Video represents a video in api response
type Video struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Thumbnail   string    `json:"thumbnail"`
}

// Cache represents the root of response
type Cache struct {
	StreamContainer StreamContainer `json:"stream"`
	VideosContainer VideoContainer  `json:"videos"`
}

// Config datas representation
type Config struct {
	Mode         string
	YoutubeKey   string
	TwitchClient string
	TwitchSecret string
}

// Application root structure
type Application struct {
	Cache  *Cache
	Config *Config
}

func (a *Application) updateStreamDatas(client http.Client) {

	log.Println("Actualisation des données Twitch en cours...")

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://id.twitch.tv/oauth2/token?client_id=%v&client_secret=%v&grant_type=client_credentials",
			a.Config.TwitchClient,
			a.Config.TwitchSecret,
		),
		nil,
	)

	if err != nil {
		log.Fatal("La requête à l'API distante à échouée")
	}

	resp, err := client.Do(req)

	var data map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &data)

	bearer := data["access_token"]

	req, err = http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_id="+TwitchChannelID, nil)

	if err != nil {
		log.Fatal("La requête à l'API distante à échouée")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", bearer))
	req.Header.Add("Client-ID", a.Config.TwitchClient)

	resp, err = client.Do(req)

	if resp.StatusCode != http.StatusOK {

		log.Printf("API status %v, echec de la mise à jour des données\n", resp.StatusCode)

		return
	}

	structuredResponse := model.TwitchResponseContainer{}

	body, _ = ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &structuredResponse)

	if len(structuredResponse.Datas) != 0 {

		a.Cache.StreamContainer.Stream = &Stream{
			Title: structuredResponse.Datas[0].Title,
			Date:  structuredResponse.Datas[0].StartedAt,
		}
	} else {
		a.Cache.StreamContainer.Stream = nil
	}

	a.Cache.StreamContainer.DateSync = time.Now()

	log.Println("Les données Twitch ont été mise à jour")
}

func (a *Application) updateYoutubeDatas(client http.Client, channelID string, isMain bool) {

	log.Println("Actualisation des données Youtube en cours...")

	req, err := http.NewRequest(
		"GET",
		"https://www.googleapis.com/youtube/v3/search?key="+
			a.Config.YoutubeKey+
			"&channelId="+
			channelID+
			"&part=snippet,id&order=date&maxResults=1",
		nil,
	)

	if err != nil {
		log.Fatal("La requête à l'API Youtube à échouée")
	}

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {

		log.Printf("API Youtube status %v, echec de la mise à jour des données\n", resp.StatusCode)

		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var structuredResponse *model.YoutubeResponseContainer

	_ = json.Unmarshal(body, &structuredResponse)

	if a.Cache.VideosContainer.Videos == nil {
		a.Cache.VideosContainer.Videos = &Videos{}
	}

	video := Video{
		ID:          structuredResponse.Datas[0].Id.Id,
		Title:       structuredResponse.Datas[0].Snippet.Title,
		Description: structuredResponse.Datas[0].Snippet.Description,
		Date:        structuredResponse.Datas[0].Snippet.PublishedAt,
		Thumbnail:   structuredResponse.Datas[0].Snippet.Thumbnails.Default.Url,
	}

	if isMain {
		a.Cache.VideosContainer.Videos.Main = video
	} else {
		a.Cache.VideosContainer.Videos.Second = video
	}

	a.Cache.VideosContainer.DateSync = time.Now()

	log.Println("Les données Youtube ont été mise à jour")
}
