package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/projects/faceit-discord/pkg/external/discord"
	"github.com/projects/faceit-discord/services"
)

func Start(){
	token := ""
	repo, err := discord.NewDiscord(token)
	if err != nil {
		fmt.Printf("%v", err)
	}
	webhookService := services.NewWebhooksService(repo)
	wh := NewWebhooksHandlers(webhookService)


	r := gin.Default()
  r.POST("/webhook", wh.handleWebhook)
	
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}