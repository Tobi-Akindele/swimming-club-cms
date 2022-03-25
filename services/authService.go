package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/jwt"
	"swimming-club-cms-be/models"
)

type AuthService struct{}

func (as *AuthService) AuthenticateUser(user *models.UserResult, password string) (*dtos.AuthResponse, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username or password is invalid")
	}
	permissionsService := PermissionService{}
	permissions, err := permissionsService.GetRolesPermissions(user.Roles)
	if err != nil {
		return nil, err
	}
	jwtTokenGenerator := jwt.TokenGenerator{}
	accessToken, err := jwtTokenGenerator.GenerateToken(user)
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
		UserType:    user.UserType,
		Roles:       user.Roles,
		Permissions: permissions,
	}, nil
}
