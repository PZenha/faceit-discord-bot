package webhooks

type Webhook struct{
	Event string `json:"event"`
}

type WebhooksRepository interface {
	WriteToChannel(event string) error
}

type WebhooksService interface {
	ProcessWebhook(webhook Webhook) error
}
