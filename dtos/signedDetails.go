package dtos

import (
	"github.com/dgrijalva/jwt-go"
	"swimming-club-cms-be/models"
)

type SignedDetails struct {
	Username   string
	Email      string
	UserId     string
	Authorized bool
	Roles      []*models.RoleDto
	jwt.StandardClaims
}
