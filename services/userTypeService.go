package services

import (
	"github.com/jinzhu/copier"
	"strings"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type UserTypeService struct{}

func (uts *UserTypeService) CreateUserType(userTypeDto *models.UserTypeDto) (*models.UserType, error) {
	userType := &models.UserType{}
	err := copier.Copy(userType, userTypeDto)
	if err != nil {
		return nil, err
	}
	userType.Name = strings.ToUpper(userType.Name)
	userType.Updatable = true
	userTypeRepository := repositories.UserTypeRepository{}
	return userTypeRepository.SaveUserType(userType)
}

func (uts *UserTypeService) GetById(id string) (*models.UserType, error) {
	userTypeRepository := repositories.UserTypeRepository{}
	return userTypeRepository.FindById(id)
}

func (uts *UserTypeService) GetAllUserTypes() ([]*models.UserType, error) {
	userTypeRepository := repositories.UserTypeRepository{}
	return userTypeRepository.FindAll()
}
