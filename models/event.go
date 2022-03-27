package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Event{})
}

type Event struct {
	mogo.DocumentModel `bson:",inline" coll:"events"`
	Name               string             `json:"name"`
	Participants       mogo.RefFieldSlice `json:"participants" ref:"User"`
}

type CreateEvent struct {
	Name          string `json:"name" binding:"required"`
	CompetitionId string `json:"competitionId" binding:"required"`
}

type AddParticipants struct {
	EventId      string   `json:"eventId" binding:"required" validate:"nonzero"`
	Participants []string `json:"participants" binding:"required" validate:"min=1"`
}
