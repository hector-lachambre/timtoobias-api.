// 	Copyright (C) 2019 Hector Lachambre
//	Copyright (C) Unknwon for go-ini project
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
	"log"
	"net/http"

	"gitlab.com/timtoobias-projects/timtoobias-api/controllers"
	"gitlab.com/timtoobias-projects/timtoobias-datas/configuration"
	"gitlab.com/timtoobias-projects/timtoobias-datas/repositories"
)

func main() {

	configurationManager := &configuration.CredentialsManager{}
	client := &http.Client{}

	controller := &controllers.LiveNotifierController{
		TwitchRepository: &repositories.TwitchRepository{
			CM:     configurationManager,
			Client: client,
		},
		YoutubeRepository: &repositories.YoutubeRepository{
			CM:     configurationManager,
			Client: client,
		},
	}

	http.HandleFunc("/datas", controller.Get)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
