package repositories

import (
	"errors"
	"goPractice/notificationSystem/constants"
	"goPractice/notificationSystem/models"
)

type NotificationRepo struct {
	Notifications map[int64]models.Notification
}

func NewNotificationRepo() NotificationRepo {
	return NotificationRepo{
		Notifications: make(map[int64]models.Notification),
	}
}

func (n *NotificationRepo) CreateNotification(recipient int64, message string, channel constants.Channel, priority constants.Priority, senderChannelID string) models.Notification {

	notification := models.NewNotification(recipient, message, channel, priority, senderChannelID)
	return notification
}

func (n *NotificationRepo) GetAllNotification() (*[]models.Notification, error) {
	if len(n.Notifications) == 0 {
		return nil, errors.New("no notification in the DB")
	}

	notifications := []models.Notification{}
	for _, notification := range n.Notifications {
		notifications = append(notifications, notification)
	}

	return &notifications, nil
}

func (n *NotificationRepo) GetNotification(notificationID int64) (*models.Notification, error) {
	notification, exist := n.Notifications[(notificationID)]
	if !exist {
		return nil, errors.New("notification not found")
	}

	return &notification, nil
}

func (n *NotificationRepo) UpdateNotificationFailed(notificationID int64) (*models.Notification, error) {
	notification, exist := n.Notifications[(notificationID)]
	if !exist {
		return nil, errors.New("notification not found")
	}

	return models.UpdateNotificationFailed(notification), nil
}

func (n *NotificationRepo) UpdateNotificationSent(notificationID int64) (*models.Notification, error) {
	notification, exist := n.Notifications[(notificationID)]
	if !exist {
		return nil, errors.New("notification not found")
	}

	return models.UpdateNotificationSent(notification), nil
}
