package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"main.go/files/entity"
)

func main() {
	// Getting the bot's token from a json file
	token := readConfig("config/config.json")

	connect(token)
}

func readConfig(path string) string {
	cfg, _ := ioutil.ReadFile(path)
	data := entity.Config{}

	err := json.Unmarshal([]byte(cfg), &data)
	if err != nil {
		fmt.Println("[connection.go: readConfig] error unmarshalling config: ", err)
	}

	return data.Token
}
