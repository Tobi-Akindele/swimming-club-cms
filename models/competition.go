package models

import (
	"github.com/goonode/mogo"
	"time"
)

func init() {
	mogo.ModelRegistry.Register(Competition{})
}

type Competition struct {
	mogo.DocumentModel `bson:",inline" coll:"competitions"`
	Name               string             `json:"name" idx:"{name}, unique"`
	Date               time.Time          `json:"date"`
	Events             mogo.RefFieldSlice `json:"events" ref:"Event"`
}

type CreateCompetition struct {
	Name string `json:"name" binding:"required"`
	Date string `json:"date" binding:"required"`
}

type CompetitionResult struct {
	mogo.DocumentModel `bson:",inline"`
	Name               string    `json:"name"`
	Date               time.Time `json:"date"`
	Events             []Event   `json:"events"`
}

type DeleteCompetition struct {
	CompetitionIds []string `json:"competitionIds" binding:"required"`
}

type RemoveEvents struct {
	CompetitionId string   `json:"competitionId" binding:"required"`
	EventIds      []string `json:"eventIds" binding:"required"`
}
