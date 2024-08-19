package models

var userID int = 0

type User struct {
	ID             int
	name           string
	Department     string
	Score          float32
	SolvedProblems map[int]bool
}

func NewUser(name string, department string) User {
	userID := getNextUserID()
	return User{
		ID:             userID,
		name:           name,
		Department:     department,
		Score:          0,
		SolvedProblems: map[int]bool{},
	}
}

func getNextUserID() int {
	userID++
	return userID
}
