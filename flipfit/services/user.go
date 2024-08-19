package services

import (
	"errors"
	"goPractice/flipfit/models"
	"goPractice/flipfit/repositories"
)

type UserService struct {
	UserRepo repositories.UserRepo
}

func NewUserService() UserService {
	return UserService{
		UserRepo: repositories.UserRepo{},
	}
}

func (u *UserService) AddUser(name string, cityID int) int {
	return u.UserRepo.AddUser(name, cityID)
}

func (u *UserService) GetUser(userID int) (*models.User, error) {
	user, err := u.UserRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) BookSlot(user *models.User, slot *models.Slot, centerID int) (bool, int, error) {
	if slot.Status == 0 {
		return false, 0, errors.New("slot is not active")
	}
	if slot.WaitlistCount >= slot.WaitingListCapacity {
		return false, 0, errors.New("wait list is full")
	}

	if slot.OccupiedCount < slot.Capacity {
		bookedSlots, found := u.UserRepo.BookedSlots[user.ID]
		if found {
			bookedSlotCountForCenter := bookedSlots[centerID]
			if len(bookedSlotCountForCenter) >= 3 {
				return false, 0, errors.New("can not book more than 3 slots in this center")
			} else {
				bookedSlotCountForCenter = append(bookedSlotCountForCenter, *slot)
				bookedSlots[centerID] = bookedSlotCountForCenter
				u.UserRepo.BookedSlots[user.ID] = bookedSlots
			}
		} else {
			bookedSlotCountForCenter := []models.Slot{*slot}
			bookedSlots[centerID] = bookedSlotCountForCenter
			u.UserRepo.BookedSlots[user.ID] = bookedSlots
		}
		slot.OccupiedCount++
		return true, 0, nil
	} else {
		slot.WaitingList = append(slot.WaitingList, user.ID)
		slot.WaitlistCount++
		return true, slot.WaitlistCount, nil
	}
}
