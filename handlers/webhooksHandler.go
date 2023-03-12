package handlers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/projects/faceit-discord/domain/webhooks"
)


type WebhookHandlers struct {
	service webhooks.WebhooksService
}


func NewWebhooksHandlers(service webhooks.WebhooksService) *WebhookHandlers{
	return &WebhookHandlers{
		service: service,
	}
}

func (wh *WebhookHandlers) handleWebhook(c *gin.Context) {
	//body := c.Request.Body
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	wh.service.ProcessWebhook(webhooks.Webhook{Event: "From Handler"})

	c.JSON(200, gin.H{
		"ok": 1,
	})
}