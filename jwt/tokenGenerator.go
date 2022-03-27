package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/utils"
	"time"
)

type TokenGenerator struct{}

func (tg *TokenGenerator) GenerateToken(user *models.UserResult) (string, error) {
	tokenExpiry, _ := strconv.ParseInt(utils.GetEnv(utils.JWT_TOKEN_EXPIRY, ""), 10, 64)
	claims := &dtos.SignedDetails{
		Username:   user.Username,
		Email:      user.Email,
		UserId:     user.ID.Hex(),
		Authorized: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tokenExpiry) * time.Hour).Unix(),
		},
	}
	secretKey := utils.GetEnv(utils.JWT_SECRET_KEY, "")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	return token, err
}
