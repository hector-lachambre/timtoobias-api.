package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {

	http.HandleFunc("/datas", provideDatas)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func provideDatas(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cache := &Cache{}

	absPath, err := filepath.Abs(path)

	_, _ = fmt.Println(absPath)

	if err != nil {
		log.Fatal("Erreur à la lecture du chemin")
	}

	cache.absPath = absPath

	if _, err := os.Stat(cache.absPath); err == nil {
		cache.getCurrentCache(cache.absPath)
	}

	// res, _ := json.Marshal(cache)

	// log.Println(string(res))

	client := http.Client{}

	if time.Since(cache.StreamContainer.DateSync).Seconds() > 30 {

		cache.updateStreamDatas(client)
	}

	if time.Since(cache.VideosContainer.DateSync).Seconds() > 60*2 {

		cache.updateYoutubeDatas(client, YT_HuzId_main, true)
		cache.updateYoutubeDatas(client, YT_HuzId_second, false)
	}

	cache.save()

	test, _ := json.Marshal(cache)

	_, _ = w.Write(test)
}
