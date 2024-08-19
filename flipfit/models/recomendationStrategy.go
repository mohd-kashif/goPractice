package models

type RecommendationStrategy interface {
	Recommend() []Slot
}

type RecommendBasedOnSlotTime struct {
}

func (st RecommendBasedOnSlotTime) Recommend() []Slot {
	return nil
}

type RecommendBasedOnSlotTimeAndWorkoutType struct {
}

func (st RecommendBasedOnSlotTimeAndWorkoutType) Recommend() []Slot {
	return nil
}
