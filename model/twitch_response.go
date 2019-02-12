package model

import "time"

type TwitchDatas struct {
	Title     string    `json:"title"`
	StartedAt time.Time `json:"started_at"`
}
type TwitchResponseContainer struct {
	Datas []TwitchDatas `json:"data"`
}
