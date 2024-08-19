package models

import (
	"errors"
	"goPractice/flipfit/constant"
)

var slotID = 0

type Slot struct {
	ID                  int
	StartTime           int
	EndTime             int
	Date                int
	Month               int
	Status              int
	Capacity            int
	WaitingListCapacity int
	OccupiedCount       int
	WaitlistCount       int
	WorkoutType         constant.WORKOUT_TYPE
	WaitingList         []int
}

func NewSlot(startHour int, startMinute int, capacity int, workoutType constant.WORKOUT_TYPE, waitingListCapacity int, date int, month int) (*Slot, error) {
	_, err := validateDate(date, month, startHour, startMinute)
	if err != nil {
		return nil, err
	}
	id := getNextSlotID()
	startTime := startHour*60 + startMinute
	endTime := (startHour+1)*60 + startMinute
	return &Slot{
		ID:                  id,
		StartTime:           startTime,
		EndTime:             endTime,
		Date:                date,
		Month:               month,
		Status:              1,
		Capacity:            capacity,
		WaitingListCapacity: waitingListCapacity,
		OccupiedCount:       0,
		WaitlistCount:       0,
		WorkoutType:         workoutType,
		WaitingList:         []int{},
	}, nil
}

func getNextSlotID() int {
	slotID++
	return slotID
}

func validateDate(date int, month int, hour int, minute int) (bool, error) {
	if month < 1 || month > 12 {
		return false, errors.New("invalid month: must be between 1 and 12")
	}

	if date < 1 || date > 31 {
		return false, errors.New("invalid date: must be between 1 and 31")
	}

	if hour < 0 || hour > 23 {
		return false, errors.New("invalid hour: must be between 0 and 23")
	}

	if minute < 0 || minute > 59 {
		return false, errors.New("invalid minute: must be between 0 and 59")
	}

	daysInMonth := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	if date > daysInMonth[month] {
		return false, errors.New("invalid date for the given month")
	}

	return true, nil
}

func (s *Slot) Enqueue(memberID int) bool {
	if len(s.WaitingList) < s.WaitingListCapacity {
		s.WaitingList = append(s.WaitingList, memberID)
		return true
	}
	return false
}

func (s *Slot) Dequeue() (int, bool) {
	if len(s.WaitingList) == 0 {
		return 0, false
	}
	memberID := s.WaitingList[0]
	s.WaitingList = s.WaitingList[1:]
	return memberID, true
}
