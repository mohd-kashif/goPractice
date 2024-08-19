package services

import (
	"goPractice/hackathon/constant"
	"goPractice/hackathon/models"
	"goPractice/hackathon/repositories"
)

type Hackathon struct {
	hackathonRepo repositories.HackathonRepo
	userRepo      repositories.UserRepo
	problemRepo   repositories.ProblemRepo
}

func NewHackathon() Hackathon {
	return Hackathon{
		hackathonRepo: repositories.HackathonRepo{},
	}
}

func (h *Hackathon) AddContestPlainScoringContest(name string, scoringStrategy models.ScoringStrategy) models.Contest {
	return h.hackathonRepo.AddContestPlainScoringContest(name)
}

func (h *Hackathon) AddContestTimeScoringContest(name string, scoringStrategy models.ScoringStrategy) models.Contest {
	return h.hackathonRepo.AddContestTimeScoringContest(name)
}

func (h *Hackathon) RegisterUser(name string, department string) models.User {
	return h.userRepo.AddUser(name, department)
}

func (h *Hackathon) RegisterProblem(description string, tag string, level constant.PROBLEM_LEVEL, score float64) models.Problem {
	return h.problemRepo.AddProblem(description, tag, level, score)
}
