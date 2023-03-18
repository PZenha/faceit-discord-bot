package discord

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/projects/faceit-discord/domain/webhooks"
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

func (d *discord) MessageWebhook(webhook webhooks.Webhook) error {
	if webhook.Event != "match_status_ready" {
		return nil
	}

	err := d.session.Open()
	if err != nil {
		return errors.New("failed to open a connection to discord")
	}

	channelID := "809940308162248747"

	embedMessage := createEmbebMessageFromWebhook(webhook)

	_, err = d.session.ChannelMessageSendEmbed(channelID, embedMessage)
	if err != nil {
		return errors.New("failed to open a connection to discord")
	}

	d.session.Close()
	return nil
}

func createEmbebMessageFromWebhook(webhook webhooks.Webhook) *discordgo.MessageEmbed {
	colorInt, err := strconv.ParseInt("0cb333", 16, 32)
	if err != nil {
		fmt.Println("Error parsing hex color:", err)
	}

	title := webhook.Payload.Teams[0].Name + " VS " + webhook.Payload.Teams[1].Name + " match is ready!"

	//@TODO Add the Discord GUIDs to a database and the MessageEmbedThumbnail to S3.
	// No warm in being hardcoded for now
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: "connect " + webhook.Payload.ClientCustom.ServerIP + ":" + webhook.Payload.ClientCustom.Server.Port,
		Color:       int(colorInt),
		Fields: []*discordgo.MessageEmbedField{
			{
				Value:  "<@" + "282674991911141378" + ">",
				Inline: true,
			},
			{
				Value:  "<@" + "326003983011414017" + ">",
				Inline: true,
			},
			{
				Value:  "<@" + "188243972362207233" + ">",
				Inline: true,
			},
			{
				Value:  "<@" + "346705438995251211" + ">",
				Inline: true,
			},
			{
				Value:  "<@" + "236267881925574656" + ">",
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.gyazo.com/892bd58ff3f7cc44f0602abf07763415.png",
		},
	}
}
