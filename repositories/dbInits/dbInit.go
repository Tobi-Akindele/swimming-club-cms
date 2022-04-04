package dbInits

import (
	"github.com/goonode/mogo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

var repositoryManager = repositories.GetRepositoryManagerInstance()
var permissionRepository = repositoryManager.GetPermissionRepository()
var roleRepository = repositoryManager.GetRoleRepository()
var userRepository = repositoryManager.GetUserRepository()
var userTypeRepository = repositoryManager.GetUserTypeRepository()

func InitializeDB() {
	loadPermissions()
	loadDefaultUserTypes()
	loadSuperAdminRole()
	loadSuperUser()
}

func loadPermissions() {
	log.Println("[+] Loading permissions... [+]")
	_, _ = permissionRepository.SavePermissions(permissions())
	log.Println("[+] Permissions loaded successfully [+]")
}

func loadDefaultUserTypes() {
	log.Println("[+] Loading default user types... [+]")
	_, _ = userTypeRepository.SaveUserTypes(defaultUserTypes())
	log.Println("[+] Default user types loaded successfully [+]")
}

func loadSuperAdminRole() {
	log.Println("[+] Loading super admin role [+]")
	superRole, _ := roleRepository.FindByName(utils.SUPER_ADMIN_ROLE)
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
	superRole, err := roleRepository.FindByName(utils.SUPER_ADMIN_ROLE)
	if err != nil {
		panic(err)
	}
	superUserType, err := userTypeRepository.FindByName(utils.SUPER_USER)
	if err != nil {
		panic(err)
	}
	su := superUser()
	hash, err := bcrypt.GenerateFromPassword([]byte(su.Password), bcrypt.MinCost)
	su.Password = string(hash)
	su.Role = mogo.RefField{ID: superRole.ID}
	su.UserType = mogo.RefField{ID: superUserType.ID}
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
