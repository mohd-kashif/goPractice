package models

var userID = 0

type User struct {
	ID   int
	Name string

	CityID int
}

func NewUser(name string, cityID int) User {
	id := getNextUserID()
	return User{
		ID:     id,
		Name:   name,
		CityID: cityID,
	}
}

func getNextUserID() int {
	userID++
	return userID
}
