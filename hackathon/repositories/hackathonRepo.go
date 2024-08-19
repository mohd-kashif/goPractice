package repositories

import "goPractice/hackathon/models"

type HackathonRepo struct {
	Contests map[int]models.Contest
}

func NewContestRepo() HackathonRepo {
	return HackathonRepo{
		Contests: map[int]models.Contest{},
	}
}

func (c *HackathonRepo) AddContestPlainScoringContest(name string) models.Contest {

	contest := models.NewContest(name, &models.PlainScoreStrategy{})
	c.Contests[contest.ID] = contest
	return contest
}

func (c *HackathonRepo) AddContestTimeScoringContest(name string) models.Contest {

	contest := models.NewContest(name, &models.TimeScoreStrategy{})
	c.Contests[contest.ID] = contest
	return contest
}
