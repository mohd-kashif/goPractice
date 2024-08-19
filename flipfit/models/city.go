package models

var cityID = 0

type City struct {
	ID   int
	Name string
}

func NewCity(name string) City {
	id := getNextCityID()
	return City{
		ID:   id,
		Name: name,
	}
}

func getNextCityID() int {
	cityID++
	return cityID
}
