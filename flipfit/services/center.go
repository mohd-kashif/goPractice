package services

import (
	"errors"
	"goPractice/flipfit/constant"
	"goPractice/flipfit/models"
	"goPractice/flipfit/repositories"
)

type CenterService struct {
	CenterRepo repositories.CenterRepo
}

func NewCenterService() CenterService {
	return CenterService{
		CenterRepo: repositories.CenterRepo{},
	}
}

func (c *CenterService) AddCenter(cityID int, name string, latitude float32, longitude float32, numberOfSlots int) int {
	recommendationStrategy := models.RecommendBasedOnSlotTimeAndWorkoutType{}
	centerID := c.CenterRepo.AddCenter(name, latitude, longitude, numberOfSlots, cityID, recommendationStrategy)
	return centerID
}

func (c *CenterService) IsCenterExist(centerID int) bool {
	return c.CenterRepo.IsCenterExist(centerID)
}

func (c *CenterService) AddSlot(centerID int, workoutType constant.WORKOUT_TYPE, startHour int, startMinute int, date int, month int, capacity int, waitingListSize int) (int, error) {
	center, err := c.CenterRepo.GetCenterByID(centerID)
	if err != nil {
		return 0, err
	}

	if center.ActiveSlots >= center.NumberOfSlots {
		return 0, errors.New("slots full for the day")
	}
	startTime := startHour*60 + startMinute
	workoutSlots := center.Slots[workoutType]
	for _, workoutSlot := range workoutSlots {
		if (workoutSlot.EndTime) > startTime && workoutSlot.Status == 1 {
			return 0, errors.New("workout with same type is overlapping")
		}
	}
	slotID, err := c.CenterRepo.AddSlot(center.ID, workoutType, startHour, startMinute, date, month, capacity, waitingListSize)
	if err != nil {
		return 0, err
	}
	return slotID, nil
}

func (c *CenterService) GetAvailableSlots(centerID int) (map[constant.WORKOUT_TYPE][]models.Slot, error) {
	return c.CenterRepo.DetailsOfWorkoutsSlots(centerID)
}
