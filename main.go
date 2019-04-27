package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"gopkg.in/ini.v1"
)

func main() {

	client := http.Client{}

	filename, err := os.Executable()

	if err != nil {
		log.Fatalf("Impossible d'obtenir les information concernant l'environnement d'éxecution")
	}

	configPath := filepath.Join(path.Dir(filename), "config.ini")

	cfg, err := ini.Load(configPath)

	if err != nil {
		log.Fatalf("Fail to read configuration file: %v", err)
	}

	application := &Application{
		Cache: &Cache{},
		Config: &Config{
			Mode:       cfg.Section("GENERAL").Key("Mode").String(),
			YoutubeKey: cfg.Section("API").Key("YT_key").String(),
			TwitchKey:  cfg.Section("API").Key("Twitch_key").String(),
		},
	}

	application.updateStreamDatas(client)
	application.updateYoutubeDatas(client, YT_HuzId_main, true)
	application.updateYoutubeDatas(client, YT_HuzId_second, false)

	http.HandleFunc("/datas", application.provideDatas)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *Application) provideDatas(w http.ResponseWriter, r *http.Request) {

	client := http.Client{}

	w.Header().Set("Content-Type", "application/json")

	if time.Since(a.Cache.StreamContainer.DateSync).Seconds() > 30 {

		a.updateStreamDatas(client)
	}

	if time.Since(a.Cache.VideosContainer.DateSync).Seconds() > 60*2 {

		a.updateYoutubeDatas(client, YT_HuzId_main, true)
		a.updateYoutubeDatas(client, YT_HuzId_second, false)
	}

	output, err := json.Marshal(a.Cache)

	if err != nil {

		log.Println("Impossible de transformer le cache en JSON")
	}

	_, err = w.Write(output)

	if err != nil {

		log.Println("Impossible d'écrire la sortie")
	}
}
