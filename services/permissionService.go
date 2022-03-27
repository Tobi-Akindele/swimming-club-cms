package services

import (
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type PermissionService struct{}

func (ps *PermissionService) GetAllPermissions() ([]*models.Permission, error) {
	permissionRepository := repositories.PermissionRepository{}
	return permissionRepository.FindAll()
}

func (ps *PermissionService) GetRolePermissions(role models.Role) (map[string]string, error) {
	permissionRepository := repositories.PermissionRepository{}
	permissions := map[string]string{}
	for i := range role.Permissions {
		permission, err := permissionRepository.FindById(role.Permissions[i].ID.Hex())
		if err == nil {
			permissions[permission.Value] = permission.Name
		}
	}

	return permissions, nil
}

func (ps *PermissionService) GetRolesPermissions(roles []models.Role) (map[string]string, error) {
	roleService := RoleService{}
	permissionRepository := repositories.PermissionRepository{}
	permissions := map[string]string{}
	for idx := range roles {
		role, err := roleService.GetById(roles[idx].ID.Hex())
		if err == nil {
			for pIdx := range role.Permissions {
				permission, err := permissionRepository.FindById(role.Permissions[pIdx].ID.Hex())
				if err == nil {
					permissions[permission.Value] = permission.Name
				}
			}
		}
	}
	return permissions, nil
}
