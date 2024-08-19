package models

import (
	"goPractice/flipfit/constant"
)

var centerID = 0

type Center struct {
	ID                     int
	Name                   string
	Latitude               float32
	Longitude              float32
	NumberOfSlots          int
	CityID                 int
	Slots                  map[constant.WORKOUT_TYPE][]Slot
	ActiveSlots            int
	RecommendationStrategy RecommendationStrategy
}

func NewCenter(name string, latitude float32, longitude float32, numberOfSlots int, cityID int, recommendationStrategy RecommendationStrategy) Center {
	id := getNextCenterID()
	return Center{
		ID:                     id,
		Name:                   name,
		Latitude:               latitude,
		Longitude:              longitude,
		NumberOfSlots:          numberOfSlots,
		CityID:                 cityID,
		Slots:                  make(map[constant.WORKOUT_TYPE][]Slot),
		ActiveSlots:            0,
		RecommendationStrategy: recommendationStrategy,
	}
}

func getNextCenterID() int {
	centerID++
	return centerID
}

func (c *Center) AddSlot(workoutType constant.WORKOUT_TYPE, startHour int, startMinute int, date int, month int, capacity int, waitingListSize int) (int, error) {
	slot, err := NewSlot(startHour, startMinute, capacity, workoutType, waitingListSize, date, month)
	if err != nil {
		return 0, nil
	}
	c.Slots[workoutType] = append(c.Slots[workoutType], *slot)
	return slot.ID, nil
}

func (c *Center) DetailsOfWorkoutsSlots() map[constant.WORKOUT_TYPE][]Slot {
	return c.Slots
}
