package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/projects/faceit-discord/domain/webhooks"
)

type WebhookHandlers struct {
	service webhooks.WebhooksService
}

func NewWebhooksHandlers(service webhooks.WebhooksService) *WebhookHandlers {
	return &WebhookHandlers{
		service: service,
	}
}

func (wh *WebhookHandlers) HandleWebhook(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//body := c.Request.Body
	b := []byte(request.Body)
	var webhook webhooks.Webhook
	json.Unmarshal(b, &webhook)

	wh.service.ProcessWebhook(webhook)

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}
