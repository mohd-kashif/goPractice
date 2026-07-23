package repositories

import (
	"errors"
	"goPractice/notificationSystem/models"
)

type UserRepo struct {
	Users map[int64]models.User
}

func NewUserRepo() UserRepo {
	return UserRepo{
		Users: make(map[int64]models.User),
	}
}

func (u *UserRepo) AddUser(email string, contact string) (int64, error) {

	user := models.NewUser(email, contact)
	return user.ID, nil
}

func (u *UserRepo) GetUser(ID int64) (*models.User, error) {
	user, exist := u.Users[ID]
	if !exist {
		return nil, errors.New("User does not exist")
	}

	return &user, nil
}
