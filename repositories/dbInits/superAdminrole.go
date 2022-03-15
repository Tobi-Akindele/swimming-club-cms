package dbInits

import (
	"swimming-club-cms-be/models"
)

func superAdminRole() *models.Role {
	return &models.Role{
		Name:      "SUPER ADMIN",
		Updatable: false,
	}
}
