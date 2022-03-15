package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
	"swimming-club-cms-be/utils"
)

type AuthController struct{}

func (authController *AuthController) SignUp(ctx *gin.Context) {

	var signup models.UserDto
	if err := ctx.ShouldBindJSON(&signup); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	_ = validator.SetValidationFunc("datetime", utils.Datetime)
	if errs := validator.Validate(signup); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}
	userService := services.UserService{}
	user, err := userService.CreateUser(&signup)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func (authController *AuthController) Login(ctx *gin.Context) {
	var login models.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	userService := services.UserService{}
	user, err := userService.GetByUsername(login.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{Code: http.StatusNotFound, Message: "User not found"})
		return
	}
	authService := services.AuthService{}
	authUser, err := authService.AuthenticateUser(user, login.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, authUser)
	}
}
