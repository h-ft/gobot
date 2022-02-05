package handler

import (
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
			logrus.Error("error getting user ", ctx.User.Username, err)
			return
		}
		str := usr.Username + "#" + usr.Discriminator
		userString = &str
	}

	resp, err := covid.GetCountryInfo(ctx.Args[1])
	if err != nil {
		logrus.Error("error getting country info: ", err)
		return
	}

	ctx.Reply(resp)
}
