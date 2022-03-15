package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Permission{})
}

type Permission struct {
	mogo.DocumentModel `bson:",inline" coll:"permissions"`
	Name               string `json:"name" idx:"{name}, unique" binding:"required"`
	Value              string `json:"value" idx:"{value}, unique" binding:"required"`
}
