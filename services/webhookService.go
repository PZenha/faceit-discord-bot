package services

import (
	"errors"

	"github.com/projects/faceit-discord/domain/webhooks"
)

type WebhookServive struct {
	repo webhooks.WebhooksRepository
}

func NewWebhooksService(repo webhooks.WebhooksRepository) *WebhookServive {
	return &WebhookServive{
		repo: repo,
	}
}

func (wh *WebhookServive) ProcessWebhook(webhook webhooks.Webhook) error {
	err := wh.repo.MessageWebhook(webhook)
	if err != nil {
		return errors.New("failed to write to channel")
	}
	return nil
}
