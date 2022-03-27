package dtos

import (
	"github.com/globalsign/mgo/bson"
	"swimming-club-cms-be/models"
)

type AuthResponse struct {
	AccessToken string            `json:"accessToken"`
	TokenType   string            `json:"tokenType"`
	Id          bson.ObjectId     `json:"_id"`
	Username    string            `json:"username"`
	Email       string            `json:"email"`
	Admin       bool              `json:"admin"`
	UserType    models.UserType   `json:"userType"`
	Role        models.Role       `json:"role"`
	Permissions map[string]string `json:"permissions"`
}
