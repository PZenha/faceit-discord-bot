package webhooks

import "time"

type WebhookPayloadClientCustomServer struct {
	Country string `json:"country"`
	Port    string `json:"port"`
	IP      string `json:"ip"`
}

type WebhookPayloadClientCustom struct {
	ServerIP   string                           `json:"server_ip"`
	Server     WebhookPayloadClientCustomServer `json:"server"`
	ServerPort int                              `json:"server_port"`
	Map        string                           `json:"map"`
	MatchID    string                           `json:"match_id"`
}

type WebhookPayloadTeamRoster struct {
	ID                string `json:"id"`
	Nickname          string `json:"nickname"`
	Avatar            string `json:"avatar"`
	GameID            string `json:"game_id"`
	GameName          string `json:"game_name"`
	GameSkillLevel    int    `json:"game_skill_level"`
	Membership        string `json:"membership"`
	AnticheatRequired bool   `json:"anticheat_required"`
}

type WebhookPayloadTeam struct {
	ID         string                     `json:"id"`
	Name       string                     `json:"name"`
	Type       string                     `json:"type"`
	Avatar     string                     `json:"avatar"`
	LeaderID   string                     `json:"leader_id"`
	CoLeaderID string                     `json:"co_leader_id"`
	Roster     []WebhookPayloadTeamRoster `json:"roster"`
}

type WebhookPayloadEntity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type WebhookPayload struct {
	ID           string                     `json:"id"`
	OrganizerID  string                     `json:"organizer_id"`
	Region       string                     `json:"region"`
	Game         string                     `json:"game"`
	Version      int                        `json:"version"`
	Entity       WebhookPayloadEntity       `json:"entity"`
	Teams        []WebhookPayloadTeam       `json:"teams"`
	ClientCustom WebhookPayloadClientCustom `json:"client_custom"`
	CreatedAt    time.Time                  `json:"created_at"`
	UpdatedAt    time.Time                  `json:"updated_at"`
}
type Webhook struct {
	TransactionID string         `json:"transaction_id"`
	Event         string         `json:"event"`
	EventID       string         `json:"event_id"`
	ThirdPartyID  string         `json:"third_party_id"`
	AppID         string         `json:"app_id"`
	Timestamp     time.Time      `json:"timestamp"`
	RetryCount    int            `json:"retry_count"`
	Version       int            `json:"version"`
	Payload       WebhookPayload `json:"payload"`
}

type WebhooksRepository interface {
	MessageWebhook(webhook Webhook) error
}

type WebhooksService interface {
	ProcessWebhook(webhook Webhook) error
}
