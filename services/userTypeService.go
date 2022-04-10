package services

import (
	"github.com/jinzhu/copier"
	"strings"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type userTypeService struct{}

func (uts *userTypeService) CreateUserType(userTypeDto *models.UserTypeDto) (*models.UserType, error) {
	userType := &models.UserType{}
	err := copier.Copy(userType, userTypeDto)
	if err != nil {
		return nil, err
	}
	userType.Name = strings.ToUpper(userType.Name)
	userType.Updatable = true
	userTypeRepository := repositories.GetRepositoryManagerInstance().GetUserTypeRepository()
	return userTypeRepository.SaveUserType(userType)
}

func (uts *userTypeService) GetById(id string) (*models.UserType, error) {
	userTypeRepository := repositories.GetRepositoryManagerInstance().GetUserTypeRepository()
	return userTypeRepository.FindById(id)
}

func (uts *userTypeService) GetAllUserTypes() ([]*models.UserType, error) {
	userTypeRepository := repositories.GetRepositoryManagerInstance().GetUserTypeRepository()
	return userTypeRepository.FindAll()
}

func (uts *userTypeService) GetByName(name string) (*models.UserType, error) {
	userTypeRepository := repositories.GetRepositoryManagerInstance().GetUserTypeRepository()
	return userTypeRepository.FindByName(strings.ToUpper(name))
}
