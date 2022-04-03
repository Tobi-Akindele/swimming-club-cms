package services

import (
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type permissionService struct{}

func (ps *permissionService) GetAllPermissions() ([]*models.Permission, error) {
	permissionRepository := repositories.GetRepositoryManagerInstance().GetPermissionRepository()
	return permissionRepository.FindAll()
}

func (ps *permissionService) GetRolePermissions(role models.Role) (map[string]string, error) {
	permissionRepository := repositories.GetRepositoryManagerInstance().GetPermissionRepository()
	permissions := map[string]string{}
	for i := range role.Permissions {
		permission, err := permissionRepository.FindById(role.Permissions[i].ID.Hex())
		if err == nil {
			permissions[permission.Value] = permission.Name
		}
	}

	return permissions, nil
}

func (ps *permissionService) GetById(id string) (*models.Permission, error) {
	permissionRepository := repositories.GetRepositoryManagerInstance().GetPermissionRepository()
	return permissionRepository.FindById(id)
}
