package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Club{})
}

type Club struct {
	mogo.DocumentModel `bson:",inline" coll:"clubs"`
	Name               string             `json:"name" idx:"{name}, unique"`
	Coach              mogo.RefField      `json:"coach" ref:"User"`
	Members            mogo.RefFieldSlice `json:"members" ref:"User"`
}

type ClubDto struct {
	Name    string `json:"name" binding:"required"`
	CoachId string `json:"coachId" binding:"required"`
}

type AddMember struct {
	ClubId     string   `json:"clubId" binding:"required" validate:"nonzero"`
	NewMembers []string `json:"newMembers" binding:"required" validate:"min=1"`
}
