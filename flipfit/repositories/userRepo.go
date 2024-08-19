package repositories

import (
	"errors"
	"goPractice/flipfit/models"
)

type UserRepo struct {
	Users       map[int]models.User
	BookedSlots map[int]map[int][]models.Slot //map of userID to centerID to list slots
}

func NewUserRepo() UserRepo {
	return UserRepo{
		Users:       make(map[int]models.User),
		BookedSlots: make(map[int]map[int][]models.Slot),
	}
}

func (u *UserRepo) AddUser(name string, cityID int) int {
	user := models.NewUser(name, cityID)
	u.Users[user.ID] = user
	return user.ID
}

func (u *UserRepo) GetUser(id int) (*models.User, error) {
	user, found := u.Users[id]
	if !found {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
