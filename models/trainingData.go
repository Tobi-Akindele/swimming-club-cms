package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(TrainingData{})
}

type TrainingData struct {
	mogo.DocumentModel `bson:",inline" coll:"trainingData"`
	Name               string             `json:"name"`
	ClubId             mogo.RefField      `json:"clubId" ref:"Club"`
	Participants       mogo.RefFieldSlice `json:"participants" ref:"User"`
	Results            mogo.RefFieldSlice `json:"results" ref:"Result"`
}

type TrainingDataResult struct {
	mogo.DocumentModel `bson:",inline"`
	Name               string       `json:"name"`
	Participants       []UserResult `json:"participants"`
	Results            []Result     `json:"results"`
}

type CreateTrainingData struct {
	Name   string `json:"name" binding:"required" validate:"nonzero"`
	ClubId string `json:"clubId" binding:"required" validate:"nonzero"`
}

type AddTDParticipants struct {
	TrainingDataId string   `json:"trainingDataId" binding:"required" validate:"nonzero"`
	ParticipantIds []string `json:"participantIds" binding:"required" validate:"min=1"`
}
