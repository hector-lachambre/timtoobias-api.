package main

import (
	"encoding/json"
	"github.com/hector-lachambre/huzlive-api/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	absPath         string
}

func (c *Cache) getCurrentCache(path string) {

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Impossible d'ouvrir le cache à partir de \"%s\"", path)
	}

	content, err := ioutil.ReadAll(file)

	if err := json.Unmarshal(content, &c); err != nil {
		log.Fatal("Impossible de transformer le cache en structure")
	}

	_ = file.Close()

}

func (c *Cache) updateStreamDatas(client http.Client) {

	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_id="+Twitch_HuzId, nil)

	if err != nil {
		log.Fatal("La requête à l'API distante à échouée")
	}

	req.Header.Add("Client-ID", Twitch_key)

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API status %v", resp.StatusCode)
	}

	structuredResponse := model.TwitchResponseContainer{}

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &structuredResponse)

	if len(structuredResponse.Datas) != 0 {

		c.StreamContainer.Stream = &Stream{
			Title: structuredResponse.Datas[0].Title,
			Date:  structuredResponse.Datas[0].StartedAt,
		}
	} else {
		c.StreamContainer.Stream = nil
	}

	c.StreamContainer.DateSync = time.Now()
}

func (c *Cache) updateYoutubeDatas(client http.Client, channelId string, isMain bool) {

	req, err := http.NewRequest(
		"GET",
		"https://www.googleapis.com/youtube/v3/search?key="+
			YT_key+
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
		log.Fatalf("API status %v", resp.StatusCode)
	}


	body, _ := ioutil.ReadAll(resp.Body)

	var structuredResponse *model.YoutubeResponseContainer

	_ = json.Unmarshal(body, &structuredResponse)

	if c.VideosContainer.Videos == nil {
		c.VideosContainer.Videos = &Videos{}
	}

	video := Video{
		Id:          structuredResponse.Datas[0].Id.Id,
		Title:       structuredResponse.Datas[0].Snippet.Title,
		Description: structuredResponse.Datas[0].Snippet.Description,
		Date:        structuredResponse.Datas[0].Snippet.PublishedAt,
		Thumbnail:   structuredResponse.Datas[0].Snippet.Thumbnails.Default.Url,
	}

	if isMain {
		c.VideosContainer.Videos.Main = video
	} else {
		c.VideosContainer.Videos.Second = video
	}

	c.VideosContainer.DateSync = time.Now()
}

func (c *Cache) save() {

	test, _ := json.Marshal(c)

	_ = ioutil.WriteFile(c.absPath, test, 0777)
}
