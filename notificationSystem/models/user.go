package models

var UserID int64

type User struct {
	ID      int64
	Email   string
	Contact string
}

func NewUser(email string, contact string) User {
	return User{
		ID:      getUserID(),
		Email:   email,
		Contact: contact,
	}
}

func getUserID() int64 {
	UserID++
	return UserID
}
