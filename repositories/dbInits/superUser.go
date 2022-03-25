package dbInits

import (
	"github.com/goonode/mogo"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/utils"
	"time"
)

func superUser() *models.User {
	return &models.User{
		Username:    "SUPERUSER",
		Email:       "superuser@swimmingclub.com",
		FirstName:   "SUPER ADMIN",
		LastName:    "SUPER ADMIN",
		MiddleName:  "",
		Password:    utils.GetEnv(utils.SUPER_USER_PASS_KEY, ""),
		DateOfBirth: time.Now(),
		UserType:    mogo.RefField{},
		Admin:       true,
		Updatable:   false,
		PhoneNumber: models.Phone{},
		Address:     models.Address{},
		Roles:       mogo.RefFieldSlice{},
	}
}
