package constant

type WORKOUT_TYPE string

const (
	Weights  WORKOUT_TYPE = "Weights"
	Cardio   WORKOUT_TYPE = "Cardio"
	Yoga     WORKOUT_TYPE = "Yoga"
	Swimming WORKOUT_TYPE = "Swimming"
)

var ValidWorkouts = map[string]bool{
	string(Weights):  true,
	string(Cardio):   true,
	string(Yoga):     true,
	string(Swimming): true,
}
