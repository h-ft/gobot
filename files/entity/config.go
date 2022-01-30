package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token      string `json:"token"`
	YoutubeKey string `json:"youtube_api_key"`
}

func ReadConfig(path string) ([]byte, Config) {
	cfg, _ := ioutil.ReadFile(path)
	data := Config{}

	return cfg, data
}

func InitConfig() *Config {
	// Getting the bot's token from a json file
	byteData, cfg := ReadConfig("config/config.json")

	err := json.Unmarshal([]byte(byteData), &cfg)
	if err != nil {
		fmt.Println("[connection.go: readConfig] error unmarshalling config: ", err)
	}
	return &cfg
}
