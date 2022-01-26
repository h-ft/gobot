package main

import (
	"main.go/files/entity"
)

var config entity.Config

func main() {
	entity.InitConfig(&config)
	connect(config.Token)
}
