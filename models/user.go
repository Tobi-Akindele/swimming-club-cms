package models

import (
	"github.com/goonode/mogo"
	"time"
)

func init() {
	mogo.ModelRegistry.Register(User{})
}

type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Username           string        `json:"username" idx:"{username}, unique"`
	Image              string        `json:"image"`
	Email              string        `json:"email" idx:"{email}, unique"`
	FirstName          string        `json:"firstName"`
	LastName           string        `json:"lastName"`
	MiddleName         string        `json:"middleName"`
	Password           string        `json:"password"`
	DateOfBirth        time.Time     `json:"dateOfBirth"`
	UserType           mogo.RefField `json:"userType" ref:"UserType"`
	Admin              bool          `json:"admin"`
	Gender             string        `json:"gender"`
	Updatable          bool          `json:"updatable"`
	PhoneNumber        string        `json:"phoneNumber"`
	Address            string        `json:"address"`
	Role               mogo.RefField `json:"role" ref:"Role"`
	Active             bool          `json:"active"`
	ActivationCode     string        `json:"activationCode"`
	Club               mogo.RefField `json:"club" ref:"Club"`
}

//goland:noinspection ALL
type UserDto struct {
	Username    string `json:"username" binding:"required" validate:"min=3, max=40, regexp=^[a-zA-Z0-9._]*$"`
	Image       string `json:"image" binding:"required"`
	Email       string `json:"email" binding:"required" validate:"regexp=^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`
	FirstName   string `json:"firstName" binding:"required" validate:"min=2, max=40"`
	LastName    string `json:"lastName" binding:"required" validate:"min=2, max=40"`
	MiddleName  string `json:"middleName"`
	DateOfBirth string `json:"dateOfBirth" binding:"required" validate:"datetime"`
	UserTypeId  string `json:"userTypeId" binding:"required" validate:"nonzero"`
	Admin       bool   `json:"admin"`
	Gender      string `json:"gender" binding:"required" validate:"nonzero"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	RoleId      string `json:"roleId" binding:"required" validate:"nonzero"`
}

type UserUpdate struct {
	Image       string `json:"image" binding:"required"`
	FirstName   string `json:"firstName" binding:"required" validate:"min=2, max=40"`
	LastName    string `json:"lastName" binding:"required" validate:"min=2, max=40"`
	MiddleName  string `json:"middleName"`
	DateOfBirth string `json:"dateOfBirth" binding:"required" validate:"datetime" copier:"-"`
	UserTypeId  string `json:"userTypeId" binding:"required" validate:"nonzero"`
	Gender      string `json:"gender" binding:"required" validate:"nonzero"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	RoleId      string `json:"roleId" binding:"required" validate:"nonzero"`
}

type UserResult struct {
	mogo.DocumentModel `bson:",inline"`
	Username           string    `json:"username"`
	Image              string    `json:"image"`
	Email              string    `json:"email"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	MiddleName         string    `json:"middleName"`
	Password           string    `json:"password"`
	DateOfBirth        time.Time `json:"dateOfBirth"`
	UserType           UserType  `json:"userType"`
	Admin              bool      `json:"admin"`
	Gender             string    `json:"gender"`
	Updatable          bool      `json:"updatable"`
	PhoneNumber        string    `json:"phoneNumber"`
	Address            string    `json:"address"`
	Role               Role      `json:"role"`
	Active             bool      `json:"active"`
	ActivationCode     string    `json:"activationCode"`
	Club               Club      `json:"club"`
	Time               string    `json:"time"`
	FinalPoint         int       `json:"finalPoint"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SetPassword struct {
	ActivationCode  string `json:"activationCode" binding:"required" validate:"nonzero"`
	Password        string `json:"password" binding:"required" validate:"nonzero, min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" validate:"nonzero, min=8"`
}
