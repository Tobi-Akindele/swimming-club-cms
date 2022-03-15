package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/jwt"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/utils"
)

type AuthService struct{}

func (as *AuthService) AuthenticateUser(user *models.User, password string) (*dtos.AuthResponse, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username or password is invalid")
	}
	jwtTokenGenerator := jwt.TokenGenerator{}
	roleService := RoleService{}
	roles, err := roleService.GetUserRoles(utils.ConvertRefFieldSliceToStringSlice(user.Roles))
	if err != nil {
		return nil, err
	}
	accessToken, err := jwtTokenGenerator.GenerateToken(user, roles)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Admin:       user.Admin,
		Roles:       roles,
	}, nil
}
