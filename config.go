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

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// Config datas representation
type Config struct {
	Mode         string
	YoutubeKey   string
	TwitchClient string
	TwitchSecret string
}

type ConfigurationError struct {
	Message string
}

func (e *ConfigurationError) Error() string {

	return e.Message
}

func ReadConfig(filename string) (*Config, error) {

	rootPath, err := os.Executable()

	if err != nil {

		return nil, &ConfigurationError{
			Message: "Can not get information about the runtime environment",
		}
	}

	configPath := filepath.Join(path.Dir(rootPath), filename)

	cfg, err := ini.Load(configPath)

	if err != nil {

		return nil, &ConfigurationError{
			Message: fmt.Sprintf("Fail to read configuration file: %v", err),
		}
	}

	return &Config{
		Mode:         cfg.Section("GENERAL").Key("Mode").String(),
		YoutubeKey:   cfg.Section("API").Key("YT_key").String(),
		TwitchClient: cfg.Section("API").Key("Twitch_client").String(),
		TwitchSecret: cfg.Section("API").Key("Twitch_secret").String(),
	}, nil
}
