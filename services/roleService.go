package services

import (
	"errors"
	"github.com/goonode/mogo"
	"strings"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type roleService struct{}

func (rs *roleService) CreateRole(roleDto *models.RoleDto) (*models.Role, error) {
	role := models.Role{}
	role.Name = strings.ToUpper(roleDto.Name)
	role.Assignable = true
	role.Updatable = true
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	return roleRepository.SaveRole(&role)
}

func (rs *roleService) GetById(id string, fetchRelationships bool) (interface{}, error) {
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	return roleRepository.FindById(id, fetchRelationships)
}

func (rs *roleService) GetAllRoles() ([]*models.Role, error) {
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	return roleRepository.FindAll()
}

func (rs *roleService) GetByName(name string) (*models.Role, error) {
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	return roleRepository.FindByName(strings.ToUpper(name))
}

func (rs *roleService) AssignPermissionsToRole(assignPermissions *models.AssignPermissions) (*models.Role, error) {
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	rawRole, _ := roleRepository.FindById(assignPermissions.RoleId, false)
	if rawRole == nil {
		return nil, errors.New("role not found")
	}
	role, _ := rawRole.(*models.Role)
	existingPermissions := utils.ConvertRefFieldSliceToStringMap(role.Permissions)
	permissionService := GetServiceManagerInstance().GetPermissionService()
	for i := range assignPermissions.PermissionIds {
		permission, _ := permissionService.GetById(assignPermissions.PermissionIds[i])
		if permission == nil {
			return nil, errors.New("unable to validate all permissions")
		}
		if !utils.MapContainsKey(existingPermissions, permission.ID.String()) {
			role.Permissions = append(role.Permissions, &mogo.RefField{ID: permission.ID})
		}
	}
	return roleRepository.SaveRole(role)
}

func (rs *roleService) RemovePermissionsFromRole(removePermissions *models.RemovePermissions) (*models.Role, error) {
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	rawRole, _ := roleRepository.FindById(removePermissions.RoleId, false)
	if rawRole == nil {
		return nil, errors.New("role not found")
	}
	role, _ := rawRole.(*models.Role)
	role.Permissions = utils.RemoveRefFromRefSlice(role.Permissions, removePermissions.PermissionIds)
	return roleRepository.SaveRole(role)
}

func (rs *roleService) GetRolePermissions(roleId string) (*[]models.Permission, error) {
	roleRepository := repositories.GetRepositoryManagerInstance().GetRoleRepository()
	rawRole, _ := roleRepository.FindById(roleId, true)
	role, _ := rawRole.(*models.RoleResult)
	return &role.Permissions, nil
}
