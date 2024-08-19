package models

var contestID = 0

type Contest struct {
	ID              int
	Name            string
	Problems        map[int]Problem
	Users           map[int]User
	ScoringStrategy ScoringStrategy
}

func NewContest(name string, scoringStrategy ScoringStrategy) Contest {
	ID := getContestID()
	return Contest{
		ID:              ID,
		Name:            name,
		Problems:        map[int]Problem{},
		Users:           map[int]User{},
		ScoringStrategy: scoringStrategy,
	}
}

func getContestID() int {
	contestID++
	return contestID
}

type ScoringStrategy interface {
	CalculateScore(Scoring) float64
}

type PlainScoreStrategy struct {
}

func (ps *PlainScoreStrategy) CalculateScore(scoring Scoring) float64 {
	return scoring.Problem.Score
}

type TimeScoreStrategy struct {
}

func (ps *TimeScoreStrategy) CalculateScore(scoring Scoring) float64 {
	return scoring.Problem.Score / float64(scoring.TimeTaken)
}
