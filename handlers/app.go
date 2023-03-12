package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/projects/faceit-discord/pkg/external/discord"
	"github.com/projects/faceit-discord/services"
)

func Start(){
	token := "ODA5OTM3MDIwMjE4NTA3MzI1.Gn3_Yw.S_TOk0uXkadVfxnbM-0J02CqaKiKJH2245VncU"
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