package services

import (
	"errors"
	"github.com/google/uuid"
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
	parsedDOB, err := time.Parse(utils.DOB_DATE_FORMAT, userDto.DateOfBirth)
	if err != nil {
		return nil, errors.New("date format does not match yyyy-MM-dd")
	}
	userTypeService := UserTypeService{}
	userType, err := userTypeService.GetById(userDto.UserTypeId)
	if err != nil {
		return nil, errors.New("unable to validate user type")
	} else if !userType.Assignable {
		return nil, errors.New("user type " + userType.Name + " is not assignable")
	}
	roleService := RoleService{}
	role, err := roleService.GetById(userDto.RoleId)
	if err != nil {
		return nil, errors.New("unable to validate role")
	} else if !role.Assignable {
		return nil, errors.New("role " + role.Name + " is not assignable")
	}

	user := &models.User{}
	if err := copier.Copy(&user, &userDto); err != nil {
		log.Println(err)
	}
	user.DateOfBirth = parsedDOB
	user.UserType = mogo.RefField{ID: userType.ID}
	user.Role = mogo.RefField{ID: role.ID}
	user.Updatable = true
	user.Active = false
	user.ActivationCode = uuid.New().String()
	userRepository := repositories.UserRepository{}
	user, err = userRepository.SaveUser(user)
	if err != nil {
		return nil, err
	}
	utils.SendMail(utils.ACCOUNT_ACTIVATION, user.Email, utils.ComposeAccountActivationEmail(user.FirstName, user.ActivationCode))
	return user, nil
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

func (us *UserService) SetPassword(setPassword *models.SetPassword) error {
	userRepository := repositories.UserRepository{}
	user, err := userRepository.FindByActivationCode(setPassword.ActivationCode)
	if err != nil {
		return errors.New("unable to verify user")
	}
	if setPassword.Password != setPassword.ConfirmPassword {
		return errors.New("password does not match")
	}
	user.Password = setPassword.Password
	user.Active = true
	user.ActivationCode = ""
	_, err = userRepository.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}
