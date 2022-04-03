package dtos

import (
	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Username   string
	Email      string
	UserId     string
	RoleId     string
	Authorized bool
	jwt.StandardClaims
}
