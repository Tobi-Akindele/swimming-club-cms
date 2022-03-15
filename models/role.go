package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Role{})
}

type Role struct {
	mogo.DocumentModel `bson:",inline" coll:"roles"`
	Name               string             `json:"name" idx:"{name}, unique" binding:"required"`
	Updatable          bool               `json:"updatable" binding:"required"`
	Permissions        mogo.RefFieldSlice `json:"permissions" ref:"Permission"`
}

type RoleDto struct {
	Name      string `json:"name" binding:"required" validate:"min=3 max=30 regexp=^[a-zA-Z0-9]*$"`
	Updatable bool   `json:"updatable" binding:"required"`
}
