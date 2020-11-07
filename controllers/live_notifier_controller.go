package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/timtoobias-projects/timtoobias-api/presenters"
	"gitlab.com/timtoobias-projects/timtoobias-core/ports"
	"gitlab.com/timtoobias-projects/timtoobias-core/usecases"
)

const (
	// YoutubeMainChannelID is the main youtube channel identifier
	YoutubeMainChannelID = "UClTaTsOdHo6UNSCYAPN1YNQ"

	// YoutubeSecondaryChannelID is the secondary youtube channel identifier
	YoutubeSecondaryChannelID = "UC0pREpRaQjcC71d_4Db40iQ"

	// TwitchChannelID is the twitch channel identifier
	TwitchChannelID = "42428057"
)

type LiveNotifierController struct {
	TwitchRepository  ports.StreamingRepository
	YoutubeRepository ports.VideosRepository
}

func (controller *LiveNotifierController) Get(w http.ResponseWriter, r *http.Request) {
	input := &usecases.GetMediaInfosInputMessage{
		StreamingChannelID:       TwitchChannelID,
		MainVideosChannelID:      YoutubeMainChannelID,
		SecondaryVideosChannelID: YoutubeSecondaryChannelID,
	}
	output := &presenters.LiveNotifierPresenter{}

	interactor := &usecases.GetMediaInfosInteractor{
		StreamingRepository: controller.TwitchRepository,
		VideosRepository:    controller.YoutubeRepository,
		Presenter:           output,
	}

	interactor.Execute(input)

	makeJSONResponseFrom(output.GetViewModel(), w)
}

func makeJSONResponseFrom(object interface{}, w http.ResponseWriter) {
	output, err := json.Marshal(object)

	if err != nil {

		log.Println("Impossible de transformer le cache en JSON")
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(output)

	if err != nil {

		log.Println("Impossible d'écrire la sortie")
	}
}
