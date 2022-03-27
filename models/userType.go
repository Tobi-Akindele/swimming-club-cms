package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(UserType{})
}

type UserType struct {
	mogo.DocumentModel `bson:",inline" coll:"usertypes"`
	Name               string `json:"name" idx:"{name}, unique"`
	Updatable          bool   `json:"updatable"`
	Assignable         bool   `json:"assignable"`
}

type UserTypeDto struct {
	Name       string `json:"name" idx:"{name}, unique" binding:"required" validate:"min=3 max=30 regexp=^[A-Z](\s?[A-Z0-9])*$"`
	Updatable  bool   `json:"updatable"`
	Assignable bool   `json:"assignable" binding:"required"`
}
