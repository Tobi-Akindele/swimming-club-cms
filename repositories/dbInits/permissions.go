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
		{
			Name:  "CREATE USER TYPE",
			Value: "CREATE_USER_TYPE",
		},
		{
			Name:  "GET ALL USER TYPES",
			Value: "GET_ALL_USER_TYPES",
		},
		{
			Name:  "CREATE CLUB",
			Value: "CREATE_CLUB",
		},
		{
			Name:  "ADD MEMBER TO CLUB",
			Value: "ADD_MEMBER_TO_CLUB",
		},
		{
			Name:  "GET CLUB BY ID",
			Value: "GET_CLUB_BY_ID",
		},
		{
			Name:  "GET ALL CLUBS",
			Value: "GET_ALL_CLUBS",
		},
		{
			Name:  "CREATE COMPETITION",
			Value: "CREATE_COMPETITION",
		},
		{
			Name:  "GET COMPETITION BY ID",
			Value: "GET_COMPETITION_BY_ID",
		},
		{
			Name:  "CREATE EVENT",
			Value: "CREATE_EVENT",
		},
		{
			Name:  "GET EVENT BY ID",
			Value: "GET_EVENT_BY_ID",
		},
		{
			Name:  "ADD PARTICIPANTS TO EVENT",
			Value: "ADD_PARTICIPANTS_TO_EVENT",
		},
		{
			Name:  "GET ALL USERS",
			Value: "GET_ALL_USERS",
		},
		{
			Name:  "GET USER BY USERNAME",
			Value: "GET_USER_BY_USERNAME",
		},
		{
			Name:  "GET USER BY EMAIL",
			Value: "GET_USER_BY_EMAIL",
		},
	}
	return pList
}
