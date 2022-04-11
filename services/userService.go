package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"log"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
	"time"
)

type userService struct{}

func (us *userService) CreateUser(userDto *models.UserDto) (*models.User, error) {
	parsedDOB, err := time.Parse(utils.DATE_FORMAT, userDto.DateOfBirth)
	if err != nil {
		return nil, errors.New("date format does not match yyyy-MM-dd")
	}
	userTypeService := GetServiceManagerInstance().GetUserTypeService()
	userType, err := userTypeService.GetById(userDto.UserTypeId)
	if err != nil {
		return nil, errors.New("unable to validate user type")
	} else if !userType.Assignable {
		return nil, errors.New("user type " + userType.Name + " is not assignable")
	}
	roleService := GetServiceManagerInstance().GetRoleService()
	rawRole, err := roleService.GetById(userDto.RoleId, false)
	role, _ := rawRole.(*models.Role)
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
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	user, err = userRepository.SaveUser(user)
	if err != nil {
		return nil, err
	}
	utils.SendMail(utils.ACCOUNT_ACTIVATION, user.Email, utils.ComposeAccountActivationEmail(user.FirstName, user.ActivationCode))
	return user, nil
}

func (us *userService) GetByUsername(username string) (*models.UserResult, error) {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	return userRepository.FindByUsername(username)
}

func (us *userService) GetById(id string, fetchRelationShips bool) (interface{}, error) {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	return userRepository.FindById(id, fetchRelationShips)
}

func (us *userService) GetByIds(ids []string) ([]models.UserResult, error) {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	return userRepository.FindByIds(ids)
}

func (us *userService) GetAllUsers() ([]*models.UserResult, error) {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	return userRepository.FindAll()
}

func (us *userService) GetByEmail(email string) (*models.UserResult, error) {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	return userRepository.FindByEmail(email)
}

func (us *userService) SetPassword(setPassword *models.SetPassword) error {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	user, err := userRepository.FindByActivationCode(setPassword.ActivationCode)
	if err != nil {
		return errors.New("unable to verify user")
	}
	if setPassword.Password != setPassword.ConfirmPassword {
		return errors.New("password does not match")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(setPassword.Password), bcrypt.MinCost)
	user.Password = string(hash)
	user.Active = true
	user.ActivationCode = ""
	_, err = userRepository.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) UpdateUser(userUpdate *models.UserUpdate, userId string) (*models.User, error) {
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	rawUser, _ := userRepository.FindById(userId, false)
	if rawUser == nil {
		return nil, errors.New("user not found")
	}
	user, _ := rawUser.(*models.User)
	if !user.Updatable {
		return nil, errors.New("user cannot be modified")
	}
	parsedDOB, err := time.Parse(utils.DATE_FORMAT, userUpdate.DateOfBirth)
	if err != nil {
		return nil, errors.New("date format does not match yyyy-MM-dd")
	}
	if len(userUpdate.UserTypeId) > 0 {
		userTypeService := GetServiceManagerInstance().GetUserTypeService()
		userType, err := userTypeService.GetById(userUpdate.UserTypeId)
		if err != nil {
			return nil, errors.New("unable to validate user type")
		} else if !userType.Assignable {
			return nil, errors.New("user type " + userType.Name + " is not assignable")
		}
		user.UserType = mogo.RefField{ID: userType.ID}
	}
	if len(userUpdate.RoleId) > 0 {
		roleService := GetServiceManagerInstance().GetRoleService()
		rawRole, err := roleService.GetById(userUpdate.RoleId, false)
		role, _ := rawRole.(*models.Role)
		if err != nil {
			return nil, errors.New("unable to validate role")
		} else if !role.Assignable {
			return nil, errors.New("role" + role.Name + " is not assignable")
		}
		user.Role = mogo.RefField{ID: role.ID}
	}

	if err := copier.Copy(&user, &userUpdate); err != nil {
		log.Println(err)
	}
	user.DateOfBirth = parsedDOB

	return userRepository.SaveUser(user)
}

func (us *userService) UpdateUsers(users []*models.User) ([]*models.User, []error) {
	var errs []error
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	for _, user := range users {
		_, err := userRepository.SaveUser(user)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errs
	}
	return users, nil
}

func (us *userService) SearchUsersByUserType(username string, userTypeName string) ([]*models.User, error) {
	userTypeService := GetServiceManagerInstance().GetUserTypeService()
	var userType *models.UserType
	if userTypeName == utils.COACH {
		userType, _ = userTypeService.GetByName(utils.COACH)
	} else if userTypeName == utils.SWIMMER {
		userType, _ = userTypeService.GetByName(utils.SWIMMER)
	}
	userRepository := repositories.GetRepositoryManagerInstance().GetUserRepository()
	return userRepository.FindAllUsersByUsernameAndUserType(username, userType.ID.Hex())
}
