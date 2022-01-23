package helper

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type IncomingMessage struct {
	*discordgo.MessageCreate
}

func (m *IncomingMessage) CheckPrefix() error {
	if !strings.HasPrefix(m.Content, "&") {
		return &WrongPrefixError{}
	}
	return nil
}
