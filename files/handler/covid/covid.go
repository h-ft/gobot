package handler

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"main.go/files/framework"
	"main.go/files/request/covid"
)

var userString *string

func CovidTrackerCommand(ctx framework.Context) {
	logrus.Info("[CovidTrackerCommand] Executing Command")
	if userString == nil {
		usr, err := ctx.Discord.User(ctx.User.ID)
		if err != nil {
			fmt.Println("error getting user ", ctx.User.Username, err)
			return
		}
		str := usr.Username + "#" + usr.Discriminator
		userString = &str
	}

	ctx.Reply(covid.GetCountryInfo(ctx.Args[1]))
}
