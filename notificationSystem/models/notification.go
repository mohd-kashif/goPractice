package models

import (
	"goPractice/notificationSystem/constants"
	"time"
)

var NotificationID int64

type Notification struct {
	ID              int64
	Recipient       int64
	Message         string
	Channel         constants.Channel
	SenderChannelID string
	Priority        constants.Priority
	Status          constants.Status
	Timestamp       time.Time
}

func NewNotification(recipient int64, message string, channel constants.Channel, priority constants.Priority, senderChannelID string) Notification {
	return Notification{
		ID:              getNotificationID(),
		Recipient:       recipient,
		Message:         message,
		Channel:         channel,
		Priority:        priority,
		Status:          constants.StatusPending,
		SenderChannelID: senderChannelID,
		Timestamp:       time.Now(),
	}
}

func UpdateNotificationSent(notification Notification) *Notification {
	return &Notification{
		ID:              notification.ID,
		Recipient:       notification.Recipient,
		Message:         notification.Message,
		Channel:         notification.Channel,
		Priority:        notification.Priority,
		Status:          constants.StatusSent,
		SenderChannelID: notification.SenderChannelID,
		Timestamp:       notification.Timestamp,
	}
}

func UpdateNotificationFailed(notification Notification) *Notification {
	return &Notification{
		ID:              notification.ID,
		Recipient:       notification.Recipient,
		Message:         notification.Message,
		Channel:         notification.Channel,
		Priority:        notification.Priority,
		Status:          constants.StatusFailed,
		SenderChannelID: notification.SenderChannelID,
		Timestamp:       notification.Timestamp,
	}
}

func getNotificationID() int64 {
	NotificationID++
	return NotificationID
}
