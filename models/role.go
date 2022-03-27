package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Role{})
}

type Role struct {
	mogo.DocumentModel `bson:",inline" coll:"roles"`
	Name               string             `json:"name" idx:"{name}, unique"`
	Updatable          bool               `json:"updatable"`
	Assignable         bool               `json:"assignable"`
	Permissions        mogo.RefFieldSlice `json:"permissions" ref:"Permission"`
}

//goland:noinspection ALL
type RoleDto struct {
	Name       string `json:"name" binding:"required" validate:"min=3 max=30 regexp=^[A-Z](\s?[A-Z0-9])*$"`
	Updatable  bool   `json:"updatable"`
	Assignable bool   `json:"assignable" binding:"required"`
}
