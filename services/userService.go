package services

import (
	"errors"
	"github.com/jinzhu/copier"
	"log"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
	"time"
)

type UserService struct{}

func (us *UserService) CreateUser(userDto *models.UserDto) (*models.User, error) {
	if userDto.Password != userDto.ConfirmPassword {
		return nil, errors.New("password does not match")
	}
	parsedDOB, err := time.Parse(utils.DOB_DATE_FORMAT, userDto.DateOfBirth)
	if err != nil {
		return nil, errors.New("date format does not match yyyy-MM-dd")
	}

	user := &models.User{}
	if err := copier.Copy(&user, &userDto); err != nil {
		log.Println(err)
	}
	user.DateOfBirth = parsedDOB
	userRepository := repositories.UserRepository{}
	return userRepository.SaveUser(user)
}

func (us *UserService) GetByUsername(username string) (*models.User, error) {
	userRepository := repositories.UserRepository{}
	return userRepository.FindByUsername(username)
}

func (us *UserService) GetById(id string) (*models.User, error) {
	userRepository := repositories.UserRepository{}
	return userRepository.FindById(id)
}
