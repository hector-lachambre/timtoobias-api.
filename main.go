package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {

	client := http.Client{}

	cache := &Cache{}

	cache.updateStreamDatas(client)
	cache.updateYoutubeDatas(client, YT_HuzId_main, true)
	cache.updateYoutubeDatas(client, YT_HuzId_second, false)

	http.HandleFunc("/datas", cache.provideDatas)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (cache *Cache) provideDatas(w http.ResponseWriter, r *http.Request) {

	client := http.Client{}

	w.Header().Set("Content-Type", "application/json")

	if time.Since(cache.StreamContainer.DateSync).Seconds() > 30 {

		cache.updateStreamDatas(client)
	}

	if time.Since(cache.VideosContainer.DateSync).Seconds() > 60*2 {

		cache.updateYoutubeDatas(client, YT_HuzId_main, true)
		cache.updateYoutubeDatas(client, YT_HuzId_second, false)
	}

	output, err := json.Marshal(cache)

	if err != nil {

		log.Println("Impossible de transformer le cache en JSON")
	}

	_, err = w.Write(output)

	if err != nil {

		log.Println("Impossible d'écrire la sortie")
	}
}
