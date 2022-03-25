package models

import "github.com/goonode/mogo"

func init() {
	mogo.ModelRegistry.Register(Permission{})
}

type Permission struct {
	mogo.DocumentModel `bson:",inline" coll:"permissions"`
	Name               string `json:"name" idx:"{name}, unique"`
	Value              string `json:"value" idx:"{value}, unique"`
}

//goland:noinspection ALL
type PermissionDto struct {
	Name string `json:"name" idx:"{name}, unique" binding:"required" validate:"min=3 max=30 regexp=^[A-Z](\s?[A-Z0-9])*$"`
}
