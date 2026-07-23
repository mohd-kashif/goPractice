package controllers

import (
	"goPractice/notificationSystem/constants"
	"goPractice/notificationSystem/models"
	"goPractice/notificationSystem/services"
)

type NotificationController struct {
	NotificationService services.NotificationService
	UserService         services.UserService
}

func NewNotificationController() NotificationController {
	return NotificationController{
		NotificationService: services.NewNotificationService(),
		UserService:         services.NewUserService(),
	}
}

func (nc *NotificationController) GetAllNotification() (*[]models.Notification, error) {
	return nc.NotificationService.GetAllNotification()
}

func (nc *NotificationController) GetNotification(notificationID int64) (*models.Notification, error) {
	return nc.NotificationService.GetNotification(notificationID)
}

func (nc *NotificationController) SendNotification(recipient int64, message string, channel constants.Channel, priority constants.Priority) error {
	notification, err := nc.NotificationService.CreateNotification(recipient, message, channel, priority)
	if err != nil {
		return err
	}

	return nc.NotificationService.SendNotification(*notification)
}

func (nc *NotificationController) RetrySend(notificationID int64) error {
	notification, err := nc.GetNotification(notificationID)
	if err != nil {
		return err
	}

	return nc.NotificationService.SendNotification(*notification)
}
