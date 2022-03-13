package models

import (
	"github.com/goonode/mogo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func init() {
	mogo.ModelRegistry.Register(User{})
}

type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Username           string    `json:"username" idx:"{username}, unique" binding:"required"`
	Email              string    `json:"email" idx:"{email}, unique" binding:"required"`
	FirstName          string    `json:"firstName" binding:"required"`
	LastName           string    `json:"lastName" binding:"required"`
	MiddleName         string    `json:"middleName"`
	Password           string    `json:"password" binding:"required"`
	DateOfBirth        time.Time `json:"dateOfBirth" binding:"required"`
	UserType           string    `json:"userType" binding:"required"`
	Admin              bool      `json:"admin" binding:"required"`
	PhoneNumber        Phone     `json:"phoneNumber"`
	Address            Address   `json:"address"`
}
type Phone struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}
type Address struct {
	PostCode    string `json:"postCode"`
	HouseNumber string `json:"houseNumber"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

//goland:noinspection ALL
type UserDto struct {
	Username        string  `json:"username" binding:"required" validate:"min=3, max=40, regexp=^[a-zA-Z0-9]*$"`
	Email           string  `json:"email" binding:"required" validate:"regexp=^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`
	FirstName       string  `json:"firstName" binding:"required" validate:"min=3, max=40"`
	LastName        string  `json:"lastName" binding:"required" validate:"min=3, max=40"`
	MiddleName      string  `json:"middleName"`
	Password        string  `json:"password" binding:"required" validate:"min=8"`
	ConfirmPassword string  `json:"confirmPassword" binding:"required" validate:"min=8"`
	DateOfBirth     string  `json:"dateOfBirth" binding:"required" validate:"datetime"`
	UserType        string  `json:"userType" binding:"required" validate:"nonzero"`
	Admin           bool    `json:"admin" binding:"required"`
	PhoneNumber     Phone   `json:"phoneNumber"`
	Address         Address `json:"address"`
}

func (u *User) BeforeSave() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return err
}

func (u *User) AfterSave() error {
	u.Password = ""
	return nil
}
