package repositories

import "goPractice/hackathon/models"

type UserRepo struct {
	Users map[int]models.User
}

func NewUserRepo() UserRepo {
	return UserRepo{
		Users: map[int]models.User{},
	}
}

func (u *UserRepo) AddUser(name string, department string) models.User {
	user := models.NewUser(name, department)

	u.Users[user.ID] = user
	return user
}
