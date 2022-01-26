package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type Config struct {
	Token      string `json:"token"`
	YoutubeKey string `json:"youtube_api_key"`
}

func ReadConfig(path string) (string, string) {
	cfg, _ := ioutil.ReadFile(path)
	data := Config{}

	err := json.Unmarshal([]byte(cfg), &data)
	if err != nil {
		fmt.Println("[connection.go: readConfig] error unmarshalling config: ", err)
	}

	return data.Token, data.YoutubeKey
}

func InitConfig(c *Config) {
	if c == nil {
		return
	}

	// Getting the bot's token from a json file
	token, youtubeKey := ReadConfig("config/config.json")

	c.Token = token
	c.YoutubeKey = youtubeKey
}

func (c Config) YoutubeClient() (*youtube.Service, error) {
	httpClient := &http.Client{
		Transport: &transport.APIKey{Key: c.YoutubeKey},
	}

	return youtube.New(httpClient)
}
