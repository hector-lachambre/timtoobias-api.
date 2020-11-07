package viewmodels

import (
	"gitlab.com/timtoobias-projects/timtoobias-core/entities"
)

type Videos struct {
	Main      *entities.Video `json:"main"`
	Secondary *entities.Video `json:"second"`
}

// VideoContainer represents the video part of the api response
type VideoContainer struct {
	Videos Videos `json:"datas"`
	// DateSync time.Time `json:"dateSync"`
}

// StreamContainer represents the stream part of the api response
type StreamContainer struct {
	Stream *entities.StreamingStatus `json:"datas"`
	// DateSync time.Time                 `json:"dateSync"`
}

// LiveNotifierViewModel represents the api response for live notifier
type LiveNotifierViewModel struct {
	StreamContainer StreamContainer `json:"stream"`
	VideosContainer VideoContainer  `json:"videos"`
}
