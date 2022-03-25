package services

import (
	"errors"
	"github.com/goonode/mogo"
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
	userTypeService := UserTypeService{}
	userType, err := userTypeService.GetById(userDto.UserTypeId)
	if err != nil {
		return nil, err
	} else if !userType.Assignable {
		return nil, errors.New("user type " + userType.Name + " is not assignable")
	}
	roleService := RoleService{}
	roles, err := roleService.GetUserRoles(userDto.Roles)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	if err := copier.Copy(&user, &userDto); err != nil {
		log.Println(err)
	}
	user.DateOfBirth = parsedDOB
	user.UserType = mogo.RefField{ID: userType.ID}
	for r := range roles {
		if roles[r].Assignable {
			user.Roles = append(user.Roles, &mogo.RefField{ID: roles[r].ID})
		}
	}
	if len(user.Roles) == 0 {
		return nil, errors.New("roles does not exist or not assignable")
	}
	user.Updatable = true
	userRepository := repositories.UserRepository{}
	return userRepository.SaveUser(user)
}

func (us *UserService) GetByUsername(username string) (*models.UserResult, error) {
	userRepository := repositories.UserRepository{}
	return userRepository.FindByUsername(username)
}

func (us *UserService) GetById(id string) (*models.UserResult, error) {
	userRepository := repositories.UserRepository{}
	return userRepository.FindById(id)
}

func (us *UserService) GetAllUsers() ([]*models.UserResult, error) {
	userRepository := repositories.UserRepository{}
	return userRepository.FindAll()
}

func (us *UserService) GetByEmail(email string) (*models.UserResult, error) {
	userRepository := repositories.UserRepository{}
	return userRepository.FindByEmail(email)
}
