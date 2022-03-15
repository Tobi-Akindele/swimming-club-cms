package dbInits

import (
	"github.com/goonode/mogo"
	"log"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

var permissionRepository repositories.PermissionRepository
var roleRepository repositories.RoleRepository
var userRepository repositories.UserRepository

func InitializeDB() {
	loadPermissions()
	loadSuperAdminRole()
	loadSuperUser()
}

func loadPermissions() {
	log.Println("[+] Loading permissions... [+]")
	_, err := permissionRepository.SavePermissions(permissions())
	if err != nil {
		panic(err)
	}
	log.Println("[+] Permissions loaded successfully [+]")
}

func loadSuperAdminRole() {
	log.Println("[+] Loading super admin role [+]")
	superRole, _ := roleRepository.FindByName(utils.SUPER_ADMIN)
	if superRole == nil {
		superRole = superAdminRole()
	}
	permissions, err := permissionRepository.FindAll()
	if err != nil {
		panic(err)
	}
	permissions = newPermissions(superRole.Permissions, permissions)
	for p := range permissions {
		superRole.Permissions = append(superRole.Permissions, &mogo.RefField{ID: permissions[p].ID})
	}
	_, _ = roleRepository.SaveRole(superRole)
	log.Println("[+] Super admin role loaded successfully [+]")
}

func loadSuperUser() {
	log.Println("[+] Loading super user [+]")
	superRole, _ := roleRepository.FindByName(utils.SUPER_ADMIN)
	if superRole == nil {
		panic("super role not found")
	}
	su := superUser()
	su.Roles = append(su.Roles, &mogo.RefField{ID: superRole.ID})
	_, _ = userRepository.SaveUser(su)
	log.Println("[+] Super user loaded successfully [+]")
}

func newPermissions(oldPermissions mogo.RefFieldSlice, allPermissions []*models.Permission) []*models.Permission {
	var newPermissionsSlice []*models.Permission
	oldPermissionsStringSlice := utils.ConvertRefFieldSliceToStringMap(oldPermissions)
	for p := range allPermissions {
		_, ok := oldPermissionsStringSlice[allPermissions[p].ID.String()]
		if !ok {
			newPermissionsSlice = append(newPermissionsSlice, allPermissions[p])
		}
	}
	return newPermissionsSlice
}
