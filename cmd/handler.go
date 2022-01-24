package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"main.go/files/helper"
	"main.go/files/request"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	logrus.Info("Processing msg : ", m.Content, " from ", m.Author.Username)

	err := helper.CheckPrefix(m.Content)
	if err != nil {
		logrus.Info("Message is not a command")
		return // Do nothing if the message is not a command
	}

	msg := strings.TrimSpace(m.Content)
	logrus.Info(msg)
	command := strings.Split(msg[1:], " ")
	logrus.Info(command)
	if len(command) == 1 {
		switch command[0] {
		case "btc", "bitcoin":
			s.ChannelMessageSend(m.ChannelID, request.GetTokenByID(90))
			return
		}
	}
}
