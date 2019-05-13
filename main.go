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
)

func main() {

	client := http.Client{}

	config, err := readConfig("config.ini")

	if err != nil {
		log.Fatalln(err.Error())
	}

	application := &Application{
		Cache:  &Cache{},
		Config: config,
	}

	application.updateStreamDatas(client)
	application.updateYoutubeDatas(client, YT_HuzId_main, true)
	application.updateYoutubeDatas(client, YT_HuzId_second, false)

	http.HandleFunc("/datas", application.provideDatas)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
