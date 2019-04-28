// 	Copyright (C) 2019 Hector Lachambre
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

package model

import "time"

type YoutubeDatasId struct {
	Id string `json:"videoId"`
}
type YoutubeDatasSnippetThumbnailsDefault struct {
	Url string `json:"url"`
}
type YoutubeDatasSnippetThumbnails struct {
	Default YoutubeDatasSnippetThumbnailsDefault `json:"default"`
}
type YoutubeDatasSnippet struct {
	Title       string                        `json:"title"`
	Description string                        `json:"description"`
	PublishedAt time.Time                     `json:"publishedAt"`
	Thumbnails  YoutubeDatasSnippetThumbnails `json:"thumbnails"`
}
type YoutubeDatas struct {
	Id      YoutubeDatasId      `json:"id"`
	Snippet YoutubeDatasSnippet `json:"snippet"`
}
type YoutubeResponseContainer struct {
	Datas []YoutubeDatas `json:"items"`
}
