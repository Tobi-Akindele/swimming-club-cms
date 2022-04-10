package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Result{})
}

type Result struct {
	mogo.DocumentModel `bson:",inline" coll:"results"`
	Participant        mogo.RefField `json:"participant" ref:"User"`
	Time               string        `json:"time"`
	FinalPoint         int           `json:"finalPoint"`
}
type RecordResult struct {
	EventId string       `json:"eventId" binding:"required"`
	Results []ResultData `json:"results" binding:"required"`
}
type ResultData struct {
	ParticipantId string `json:"participantId" binding:"required"`
	ResultId      string `json:"resultId"`
	Time          string `json:"time" binding:"required"`
	FinalPoint    int    `json:"finalPoint" binding:"required"`
}
