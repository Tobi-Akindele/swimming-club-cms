package dbInits

import "swimming-club-cms-be/models"

func permissions() []*models.Permission {
	pList := []*models.Permission{
		{
			Name:  "CREATE ROLE",
			Value: "CREATE_ROLE",
		},
		{
			Name:  "GET ALL PERMISSIONS",
			Value: "GET_ALL_PERMISSIONS",
		},
	}
	return pList
}
