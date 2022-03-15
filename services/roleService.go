package services

import (
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type RoleService struct{}

func (rs *RoleService) CreateRole(roleDto *models.RoleDto) (*models.Role, error) {
	role := models.Role{}
	role.Name = roleDto.Name
	roleRepository := repositories.RoleRepository{}
	return roleRepository.SaveRole(&role)
}

func (rs *RoleService) GetById(id string) (*models.Role, error) {
	roleRepository := repositories.RoleRepository{}
	return roleRepository.FindById(id)
}

func (rs *RoleService) GetUserRoles(roleIds []string) ([]*models.RoleDto, error) {
	var roles []*models.RoleDto
	roleRepository := repositories.RoleRepository{}
	for idx := range roleIds {
		role, err := roleRepository.FindById(roleIds[idx])
		if err == nil {
			roleDto := &models.RoleDto{
				Name:      role.Name,
				Updatable: role.Updatable,
			}
			roles = append(roles, roleDto)
		}
	}
	return roles, nil
}
