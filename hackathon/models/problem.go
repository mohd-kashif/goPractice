package models

import "goPractice/hackathon/constant"

var problemID = 0

type Problem struct {
	ID          int
	Description string
	Tag         string
	Level       constant.PROBLEM_LEVEL
	AverageTime float64
	SolvedCount int
	Score       float64
}

func NewProblem(description string, tag string, level constant.PROBLEM_LEVEL, score float64) Problem {
	problemID := getNextProblemID()
	return Problem{
		ID:          problemID,
		Description: description,
		Tag:         tag,
		Level:       level,
		AverageTime: 0,
		SolvedCount: 0,
		Score:       score,
	}
}

func getNextProblemID() int {
	problemID++
	return problemID
}
