package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type discord struct {
	session *discordgo.Session
}

var (
	Token string
)

func NewDiscord(token string) (*discord, error) {
	discordSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.New("failed to set up discord bot")
	}

	discord := &discord{
		session: discordSession,
	}

	return discord, nil
}

func (d *discord) WriteToChannel(message string) error {
	err := d.session.Open()
	if err != nil {
		return errors.New("failed to open a connection to discord")
	}

	channelID := "809940308162248747"

	_, err = d.session.ChannelMessageSend(channelID, message)
	if err != nil {
		return errors.New("failed to open a connection to discord")
	}

	d.session.Close()
	return nil
}
