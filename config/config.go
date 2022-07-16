package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Url_shortener struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"url_shortener"`
}

func ReadConfig(path string) (*Config, error) {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var res Config
	if err := json.Unmarshal(config, &res); err != nil {
		return nil, err
	}
	return &res, err
}
