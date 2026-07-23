package services

import "errors"

type Channel interface {
	Send(message string, senderChannelID string) (bool, error)
	IsChannelUp() bool
}

type ChannelEmail struct {
}

func (ce *ChannelEmail) Send(message string, email string) (bool, error) {
	isChannelUp := ce.IsChannelUp()
	if !isChannelUp {
		return false, errors.New("email channel is down")
	}

	return true, nil
}

func (ce *ChannelEmail) IsChannelUp() bool {
	//check if channel external service is up
	return true
}

type ChannelSMS struct {
}

func (cs *ChannelSMS) Send(message string, contact string) (bool, error) {
	isChannelUp := cs.IsChannelUp()
	if !isChannelUp {
		return false, errors.New("SMS channel is down")
	}

	return true, nil
}

func (cs *ChannelSMS) IsChannelUp() bool {
	//check if channel external service is up
	return true
}

type ChannelPushNotification struct {
}

func (cp *ChannelPushNotification) Send(message string, contact string) (bool, error) {
	isChannelUp := cp.IsChannelUp()
	if !isChannelUp {
		return false, errors.New("Push Notification channel is down")
	}

	return true, nil
}

func (cp *ChannelPushNotification) IsChannelUp() bool {
	//check if channel external service is up
	return true
}
