package repositories

import (
	"errors"
	"goPractice/flipfit/constant"
	"goPractice/flipfit/models"
)

type CenterRepo struct {
	Centers map[int]models.Center
}

func NewCenterRepo() CenterRepo {
	return CenterRepo{
		Centers: make(map[int]models.Center),
	}
}

func (c *CenterRepo) AddCenter(name string, latitude float32, longitude float32, numberOfSlots int, cityID int, recommendationStrategy models.RecommendationStrategy) int {
	center := models.NewCenter(name, latitude, longitude, numberOfSlots, cityID, recommendationStrategy)
	c.Centers[center.ID] = center

	return center.ID
}

func (c *CenterRepo) IsCenterExist(centerID int) bool {
	_, found := c.Centers[centerID]
	return found
}

func (c *CenterRepo) AddSlot(centerID int, workoutType constant.WORKOUT_TYPE, startHour int, startMinute int, date int, month int, capacity int, waitingListSize int) (int, error) {
	center, found := c.Centers[centerID]
	if !found {
		return 0, errors.New("center not found")
	}
	return center.AddSlot(workoutType, startHour, date, month, startMinute, capacity, waitingListSize)
}

func (c *CenterRepo) DetailsOfWorkoutsSlots(centerID int) (map[constant.WORKOUT_TYPE][]models.Slot, error) {
	center, found := c.Centers[centerID]
	if !found {
		return nil, errors.New("center not found")
	}
	return center.DetailsOfWorkoutsSlots(), nil
}

func (c *CenterRepo) GetCenterByID(centerID int) (*models.Center, error) {
	center, found := c.Centers[centerID]
	if !found {
		return nil, errors.New("center not found")
	}
	return &center, nil
}
