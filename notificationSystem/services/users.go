package services

import (
	"errors"
	"fmt"
	"goPractice/notificationSystem/constants"
	"goPractice/notificationSystem/repositories"
	"net/mail"
	"regexp"
	"strings"
)

type UserService struct {
	UserRepo repositories.UserRepo
}

func NewUserService() UserService {
	return UserService{
		UserRepo: repositories.UserRepo{},
	}
}

func (u *UserService) AddUser(email string, contact string) (int64, error) {
	err := validateEmail(email)
	if err != nil {
		return 0, err
	}

	err = validateContact(contact)
	if err != nil {
		return 0, err
	}

	userID, err := u.UserRepo.AddUser(email, contact)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u *UserService) GetSenderChannelID(userID int64, channel string) (string, error) {
	user, err := u.UserRepo.GetUser(userID)
	if err != nil {
		return "", err
	}

	switch channel {
	case string(constants.ChannelEmail):
		return user.Email, nil

	case string(constants.ChannelPushNotification):
		return user.Contact, nil

	case string(constants.ChannelSMS):
		return user.Contact, nil

	default:
		return "", errors.New("unknown channel")
	}
}

func validateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("email address cannot be empty")
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email format: %w", err)
	}

	if addr.Address != email {
		return errors.New("email address must not contain a display name")
	}

	return nil
}

func validateContact(contact string) error {
	var contactRegex = regexp.MustCompile(`^[1-9]\d{9}$`)
	contact = strings.TrimSpace(contact)

	if contact == "" {
		return errors.New("contact number cannot be empty")
	}

	if !contactRegex.MatchString(contact) {
		return errors.New("contact must be a valid 10-digit number")
	}

	return nil
}
