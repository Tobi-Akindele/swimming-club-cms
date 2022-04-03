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
		{
			Name:  "GET ALL ROLES",
			Value: "GET_ALL_ROLES",
		},
		{
			Name:  "GET ROLE BY NAME",
			Value: "GET_ROLE_BY_NAME",
		},
		{
			Name:  "GET ROLE BY ID",
			Value: "GET_ROLE_BY_ID",
		},
		{
			Name:  "ASSIGN PERMISSIONS TO ROLE",
			Value: "ASSIGN_PERMISSIONS_TO_ROLE",
		},
		{
			Name:  "REMOVE ROLE PERMISSIONS",
			Value: "REMOVE_ROLE_PERMISSIONS",
		},
		{
			Name:  "GET ROLE PERMISSIONS",
			Value: "GET_ROLE_PERMISSIONS",
		},
		{
			Name:  "GET ALL COMPETITIONS",
			Value: "GET_ALL_COMPETITIONS",
		},
	}
	return pList
}
