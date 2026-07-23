package services

import (
	"errors"
	"goPractice/notificationSystem/constants"
	"goPractice/notificationSystem/models"
	"goPractice/notificationSystem/repositories"
)

type NotificationService struct {
	NotificationRepo repositories.NotificationRepo
	UserService      UserService
}

func NewNotificationService() NotificationService {
	return NotificationService{
		NotificationRepo: repositories.NewNotificationRepo(),
		UserService:      NewUserService(),
	}
}

func (n *NotificationService) CreateNotification(recipient int64, message string, channel constants.Channel, priority constants.Priority) (*models.Notification, error) {
	err := validateMessage(message)
	if err != nil {
		return nil, err
	}

	senderChannelID, err := n.UserService.GetSenderChannelID(recipient, string(channel))
	if err != nil {
		return nil, nil
	}

	notification := n.NotificationRepo.CreateNotification(recipient, message, channel, priority, senderChannelID)
	return &notification, nil
}

func (n *NotificationService) SendNotification(notification models.Notification) error {
	channel, err := NewChannel(notification.Channel)
	if err != nil {
		return err
	}

	isSent, err := channel.Send(notification.Message, notification.SenderChannelID)
	if err != nil {
		return err
	}
	if isSent {
		_, err := n.NotificationRepo.UpdateNotificationSent(notification.ID)
		if err != nil {
			return err
		}
	}

	_, err = n.NotificationRepo.UpdateNotificationFailed(notification.ID)
	if err != nil {
		return err
	}
	return errors.New("notification sending failed")
}

func (n *NotificationService) GetAllNotification() (*[]models.Notification, error) {
	return n.NotificationRepo.GetAllNotification()
}

func (n *NotificationService) GetNotification(notificationID int64) (*models.Notification, error) {
	return n.NotificationRepo.GetNotification(notificationID)
}

func validateMessage(message string) error {
	if len(message) > 250 {
		return errors.New("message can not be greater than 250 characters")
	}

	return nil
}
