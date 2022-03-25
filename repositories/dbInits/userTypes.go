package dbInits

import (
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/utils"
)

func defaultUserTypes() []*models.UserType {
	return []*models.UserType{
		{
			Name:       "SUPER USER",
			Updatable:  false,
			Assignable: false,
		},
		{
			Name:       utils.ADMIN,
			Updatable:  false,
			Assignable: true,
		},
		{
			Name:       utils.COACH,
			Updatable:  false,
			Assignable: true,
		},
		{
			Name:       utils.PARENT,
			Updatable:  false,
			Assignable: true,
		},
		{
			Name:       utils.SWIMMER,
			Updatable:  false,
			Assignable: true,
		},
	}
}
