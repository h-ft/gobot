package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"main.go/files/entity"
	"main.go/files/framework"
	covid "main.go/files/handler/covid"
	"main.go/files/helper"
)

var (
	config         *entity.Config
	CommandHandler *framework.CommandHandler
	Sessions       *framework.SessionManager
)

func connect() {
	config = entity.InitConfig()
	// Creates a new command handler instance
	// and register our commands
	CommandHandler = framework.NewCommandHandler()
	registerCommands()

	Sessions = framework.NewSessionManager()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	defer func() {
		dg.Close()
	}()

	if err != nil {
		logrus.Info("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(commandHandler)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		logrus.Info("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	logrus.Info("Processing msg : ", m.Content, " from ", m.Author.Username)

	err := helper.CheckPrefix(m.Content)
	if err != nil {
		logrus.Info("[connect] Message is not a command")
		return // Do nothing if the message is not a command
	}

	msg := strings.TrimSpace(m.Content)
	logrus.Info(msg)
	name := strings.Split(msg[1:], " ")
	logrus.Info(name)
	command, found := CommandHandler.Get(name[0])
	if !found {
		logrus.Info("[connect] Command not found: ", name)
		return
	}
	logrus.Info(&command)
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		logrus.Info("[connect] Channel ID,", channel)
		logrus.Info("[connect] Error getting channel,", err)
		return
	}
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		logrus.Info("[connect] Error getting guild,", err)
		return
	}
	ctx := framework.NewContext(s, guild, channel, m.Author, m, config, CommandHandler, Sessions)
	ctx.Args = name
	c := *command
	logrus.Info(c)
	c(*ctx)
}

func registerCommands() {
	CommandHandler.Register("covid", covid.CovidTrackerCommand, "Gives you information about how fucked we are")
}
