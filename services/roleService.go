package services

import (
	"strings"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type RoleService struct{}

func (rs *RoleService) CreateRole(roleDto *models.RoleDto) (*models.Role, error) {
	role := models.Role{}
	role.Name = strings.ToUpper(roleDto.Name)
	role.Assignable = roleDto.Assignable
	role.Updatable = true
	roleRepository := repositories.RoleRepository{}
	return roleRepository.SaveRole(&role)
}

func (rs *RoleService) GetById(id string) (*models.Role, error) {
	roleRepository := repositories.RoleRepository{}
	return roleRepository.FindById(id)
}

func (rs *RoleService) GetUserRoles(roleIds []string) ([]*models.Role, error) {
	var roles []*models.Role
	roleRepository := repositories.RoleRepository{}
	for idx := range roleIds {
		role, err := roleRepository.FindById(roleIds[idx])
		if err == nil {
			roles = append(roles, role)
		}
	}
	return roles, nil
}

func (rs *RoleService) GetAllRoles() ([]*models.Role, error) {
	roleRepository := repositories.RoleRepository{}
	return roleRepository.FindAll()
}
