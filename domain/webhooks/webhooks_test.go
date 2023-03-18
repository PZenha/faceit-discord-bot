package webhooks_test

import (
	"testing"

	"github.com/projects/faceit-discord/domain/webhooks"
	"github.com/projects/faceit-discord/services"
)

type MockWebhooksRepo struct{}

func (m *MockWebhooksRepo) MessageWebhook(webhook webhooks.Webhook) error {
	return nil
}

func TestProcessWebhooks(t *testing.T) {
	mockRepo := &MockWebhooksRepo{}
	webhookService := services.NewWebhooksService(mockRepo)

	err := webhookService.ProcessWebhook(webhooks.Webhook{Event: "mockEvent"})
	if err != nil {
		t.Errorf("Unexpected error, %v", err)
	}
}
