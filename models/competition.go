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
	Status             int                `json:"status"`
	Events             mogo.RefFieldSlice `json:"events" ref:"Event"`
}

type CreateCompetition struct {
	Name string    `json:"name" binding:"required"`
	Date time.Time `json:"date"`
}
