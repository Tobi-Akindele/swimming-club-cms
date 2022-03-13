package services

import (
	"errors"
	"github.com/jinzhu/copier"
	"log"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type AuthService struct{}

func (authService *AuthService) CreateUser(userDto *models.UserDto) (*models.User, error) {
	if userDto.Password != userDto.ConfirmPassword {
		return nil, errors.New("password does not match")
	}
	user := &models.User{}
	if err := copier.Copy(&user, &userDto); err != nil {
		log.Println(err)
	}
	userRepository := repositories.UserRepository{}
	return userRepository.SaveUser(user)
}
