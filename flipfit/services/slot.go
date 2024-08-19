package services

import (
	"errors"
	"goPractice/flipfit/models"
	"goPractice/flipfit/repositories"
)

type SlotService struct {
	SlotRepo repositories.SlotRepo
}

func NewSlotService() SlotService {
	return SlotService{
		SlotRepo: repositories.SlotRepo{},
	}
}

func (s *SlotService) GetSlot(slotID int) (*models.Slot, error) {
	slot, found := s.SlotRepo.Slots[slotID]
	if !found {
		return nil, errors.New("slot not found")
	}
	return &slot, nil
}
