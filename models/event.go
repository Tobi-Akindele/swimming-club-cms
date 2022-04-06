package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Event{})
}

type Event struct {
	mogo.DocumentModel `bson:",inline" coll:"events"`
	Name               string             `json:"name"`
	Participants       mogo.RefFieldSlice `json:"participants" ref:"User"`
	Results            mogo.RefFieldSlice `json:"results" ref:"Result"`
}

type EventResult struct {
	mogo.DocumentModel `bson:",inline"`
	Name               string       `json:"name"`
	Participants       []UserResult `json:"participants"`
	Results            []Result     `json:"results"`
}

type CreateEvent struct {
	Name          string `json:"name" binding:"required"`
	CompetitionId string `json:"competitionId" binding:"required"`
}

type AddParticipant struct {
	EventId     string `json:"eventId" binding:"required" validate:"nonzero"`
	Participant string `json:"participant" binding:"required" validate:"nonzero"`
}

type EventByName struct {
	Name          string `json:"name" binding:"required"`
	CompetitionId string `json:"competitionId" binding:"required"`
}

type RemoveParticipants struct {
	EventId        string   `json:"eventId" binding:"required"`
	ParticipantIds []string `json:"participantIds" binding:"required"`
}
