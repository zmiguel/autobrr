package domain

import (
	"context"
	"time"
)

type NotificationRepo interface {
	List(ctx context.Context) ([]Notification, error)
	//FindByType(ctx context.Context, nType NotificationType) ([]Notification, error)
	FindByID(ctx context.Context, id int) (*Notification, error)
	Store(ctx context.Context, notification Notification) (*Notification, error)
	Delete(ctx context.Context, notificationID int) error
}

type Notification struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	Type      NotificationType    `json:"type"`
	Enabled   bool                `json:"enabled"`
	Events    []NotificationEvent `json:"events"`
	Token     string              `json:"token"`
	APIKey    string              `json:"api_key"`
	Webhook   string              `json:"webhook"`
	Title     string              `json:"title"`
	Icon      string              `json:"icon"`
	Username  string              `json:"username"`
	Host      string              `json:"host"`
	Password  string              `json:"password"`
	Channel   string              `json:"channel"`
	Rooms     string              `json:"rooms"`
	Targets   string              `json:"targets"`
	Devices   string              `json:"devices"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

type NotificationType string

const (
	NotificationTypeDiscord    NotificationType = "DISCORD"
	NotificationTypeIFTTT      NotificationType = "IFTTT"
	NotificationTypeJoin       NotificationType = "JOIN"
	NotificationTypeMattermost NotificationType = "MATTERMOST"
	NotificationTypeMatrix     NotificationType = "MATRIX"
	NotificationTypePushBullet NotificationType = "PUSH_BULLET"
	NotificationTypePushover   NotificationType = "PUSHOVER"
	NotificationTypeRocketChat NotificationType = "ROCKETCHAT"
	NotificationTypeSlack      NotificationType = "SLACK"
	NotificationTypeTelegram   NotificationType = "TELEGRAM"
)

type NotificationEvent string

const (
	NotificationEventPushApproved    NotificationEvent = "PUSH_APPROVED"
	NotificationEventPushRejected    NotificationEvent = "PUSH_REJECTED"
	NotificationEventUpdateAvailable NotificationEvent = "UPDATE_AVAILABLE"
	NotificationEventIRCHealth       NotificationEvent = "IRC_HEALTH"
)
