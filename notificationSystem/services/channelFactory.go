package services

import (
	"errors"
	"goPractice/notificationSystem/constants"
)

func NewChannel(channel constants.Channel) (Channel, error) {
	switch channel {
	case constants.ChannelEmail:
		return &ChannelEmail{}, nil
	case constants.ChannelPushNotification:
		return &ChannelPushNotification{}, nil
	case constants.ChannelSMS:
		return &ChannelSMS{}, nil
	default:
		return nil, errors.New("invalid channel")
	}
}
