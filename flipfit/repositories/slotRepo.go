package repositories

import (
	"goPractice/flipfit/constant"
	"goPractice/flipfit/models"
)

type SlotRepo struct {
	Slots map[int]models.Slot
}

func NewSlotRepo() SlotRepo {
	return SlotRepo{
		Slots: make(map[int]models.Slot),
	}
}

func (s *SlotRepo) AddSlot(startHour int, startMinute int, date int, month int, capacity int, workoutType constant.WORKOUT_TYPE, waitingListCapacity int) (int, error) {
	slot, err := models.NewSlot(startHour, startMinute, capacity, workoutType, waitingListCapacity, date, month)
	if err != nil {
		return 0, err
	}
	s.Slots[slot.ID] = *slot
	return slot.ID, nil
}

func (s *SlotRepo) GetSlot(slotID int) models.Slot {
	return s.Slots[slotID]
}
