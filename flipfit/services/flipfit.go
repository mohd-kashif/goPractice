package services

import (
	"errors"
	"goPractice/flipfit/constant"
	"goPractice/flipfit/models"
	"goPractice/flipfit/repositories"
)

type FlipfitService struct {
	FlipfitRepo   repositories.CityRepo
	UserService   UserService
	CenterService CenterService
	SlotService   SlotService
}

type SlotTime struct {
	StartHour   int
	StartMinute int
	EndHour     int
	EndMinute   int
}

func NewFlipfitService() FlipfitService {
	return FlipfitService{
		FlipfitRepo: repositories.CityRepo{},
	}
}

func (f *FlipfitService) AddCity(cityName string) (int, error) {
	cityID, err := f.FlipfitRepo.AddCity(cityName)
	if err != nil {
		return 0, err
	}

	return cityID, nil
}

func (f *FlipfitService) AddCenter(cityID int, name string, latitude float32, longitude float32, numberOfSlots int) (int, error) {
	if !f.IsCityExist(cityID) {
		return 0, errors.New("city does not exist")
	}
	centerID := f.CenterService.AddCenter(cityID, name, latitude, longitude, numberOfSlots)

	return centerID, nil
}

func (f *FlipfitService) IsCityExist(cityID int) bool {
	return f.FlipfitRepo.IsCityExist(cityID)
}

func (f *FlipfitService) AddSlot(centerID int, workoutType string, startHour int, startMinute int, capacity int, date int, month int, waitingListSize int) (int, error) {
	_, validWorkoutType := constant.ValidWorkouts[workoutType]
	if !validWorkoutType {
		return 0, errors.New("workout type not allowed")
	}
	return f.CenterService.AddSlot(centerID, constant.WORKOUT_TYPE(workoutType), startHour, startMinute, date, month, capacity, waitingListSize)
}

func (f *FlipfitService) GetAvailableSlots(centerID int) (map[constant.WORKOUT_TYPE][]models.Slot, error) {
	return f.CenterService.GetAvailableSlots(centerID)
}

func (f *FlipfitService) RegisterUser(name string, cityID int) (int, error) {
	if !f.IsCityExist(cityID) {
		return 0, errors.New("city does not exist")
	}
	userID := f.UserService.AddUser(name, cityID)
	return userID, nil
}

func (f *FlipfitService) BookSlot(userID int, centerID int, slotID int) (bool, int, error) {
	user, err := f.UserService.GetUser(userID)
	if err != nil {
		return false, 0, err
	}

	slot, err := f.SlotService.GetSlot(slotID)
	if err != nil {
		return false, 0, err
	}
	return f.UserService.BookSlot(user, slot, centerID)
}

func (f *FlipfitService) RecommendSlot(userID int, centerID int, workoutType constant.WORKOUT_TYPE, slotTime SlotTime) []models.Slot {
	return nil
}
