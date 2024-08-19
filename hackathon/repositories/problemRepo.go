package repositories

import (
	"goPractice/hackathon/constant"
	"goPractice/hackathon/models"
)

type ProblemRepo struct {
	Problems map[int]models.Problem
}

func NewProblemRepo() ProblemRepo {
	return ProblemRepo{
		Problems: map[int]models.Problem{},
	}
}

func (p *ProblemRepo) AddProblem(description string, tag string, level constant.PROBLEM_LEVEL, score float64) models.Problem {
	problem := models.NewProblem(description, tag, level, score)
	p.Problems[problem.ID] = problem
	return problem
}
