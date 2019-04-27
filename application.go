package main

import (
	"encoding/json"
	"github.com/hector-lachambre/huzlive-api/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type VideoContainer struct {
	Videos   *Videos   `json:"datas"`
	DateSync time.Time `json:"dateSync"`
}

type Stream struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type StreamContainer struct {
	Stream   *Stream   `json:"datas"`
	DateSync time.Time `json:"dateSync"`
}

type Videos struct {
	Main   Video `json:"main"`
	Second Video `json:"second"`
}

type Video struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Thumbnail   string    `json:"thumbnail"`
}

type Cache struct {
	StreamContainer StreamContainer `json:"stream"`
	VideosContainer VideoContainer  `json:"videos"`
}

type Config struct {
	Mode       string
	YoutubeKey string
	TwitchKey  string
}

type Application struct {
	Cache  *Cache
	Config *Config
}

func (a *Application) updateStreamDatas(client http.Client) {

	log.Println("Actualisation des données Twitch en cours...")

	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_id="+Twitch_HuzId, nil)

	if err != nil {
		log.Fatal("La requête à l'API distante à échouée")
	}

	req.Header.Add("Client-ID", a.Config.TwitchKey)

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {

		log.Println("API status %v, echec de la mise à jour des données", resp.StatusCode)

		return
	}

	structuredResponse := model.TwitchResponseContainer{}

	body, _ := ioutil.ReadAll(resp.Body)

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

func (a *Application) updateYoutubeDatas(client http.Client, channelId string, isMain bool) {

	log.Println("Actualisation des données Youtube en cours...")

	req, err := http.NewRequest(
		"GET",
		"https://www.googleapis.com/youtube/v3/search?key="+
			a.Config.YoutubeKey+
			"&channelId="+
			channelId+
			"&part=snippet,id&order=date&maxResults=1",
		nil,
	)

	if err != nil {
		log.Fatal("La requête à l'API Youtube à échouée")
	}

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {

		log.Println("API Youtube status %v, echec de la mise à jour des données", resp.StatusCode)

		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var structuredResponse *model.YoutubeResponseContainer

	_ = json.Unmarshal(body, &structuredResponse)

	if a.Cache.VideosContainer.Videos == nil {
		a.Cache.VideosContainer.Videos = &Videos{}
	}

	video := Video{
		Id:          structuredResponse.Datas[0].Id.Id,
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
