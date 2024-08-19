package repositories

import (
	"errors"
	"goPractice/flipfit/models"
)

type CityRepo struct {
	Cities    map[int]models.City
	cityNames map[string]bool
}

func NewFlipRepo() CityRepo {
	return CityRepo{
		Cities: make(map[int]models.City),
	}
}

func (f *CityRepo) AddCity(cityName string) (int, error) {
	_, found := f.cityNames[cityName]
	if found {
		return 0, errors.New("city is already present")
	}

	city := models.NewCity(cityName)
	f.Cities[city.ID] = city
	return city.ID, nil
}

func (f *CityRepo) IsCityExist(ID int) bool {
	_, found := f.Cities[ID]
	return found
}
