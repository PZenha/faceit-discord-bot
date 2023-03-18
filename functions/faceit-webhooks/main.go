package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/projects/faceit-discord/handlers"
	"github.com/projects/faceit-discord/pkg/external/discord"
	"github.com/projects/faceit-discord/services"
)

func main() {
	discordToken, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		panic("Need TABLE environment variable")
	}

	repo, err := discord.NewDiscord(discordToken)
	if err != nil {
		fmt.Printf("%v", err)
	}

	webhookService := services.NewWebhooksService(repo)
	handler := handlers.NewWebhooksHandlers(webhookService)

	lambda.Start(handler.HandleWebhook)
}
