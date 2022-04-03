package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/jwt"
	"swimming-club-cms-be/models"
)

type authService struct{}

func (as *authService) AuthenticateUser(login *models.Login) (*dtos.AuthResponse, error) {
	userService := GetServiceManagerInstance().GetUserService()
	user, err := userService.GetByUsername(login.Username)
	if err != nil {
		return nil, err
	}
	if !user.Active {
		return nil, errors.New("kindly activate your account")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		log.Println(err)
		return nil, errors.New("username or password is invalid")
	}
	//permissionsService := permissionService{}
	//permissions, err := permissionsService.GetRolePermissions(user.Role)
	//if err != nil {
	//	return nil, err
	//}
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
		Role:        user.Role,
		//Permissions: permissions,
	}, nil
}
