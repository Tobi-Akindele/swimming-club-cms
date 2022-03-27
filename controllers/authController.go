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
		ctx.JSON(http.StatusCreated, user)
	}
}

func (authController *AuthController) Login(ctx *gin.Context) {
	var login models.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	authService := services.AuthService{}
	authUser, err := authService.AuthenticateUser(&login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, authUser)
	}
}

func (authController *AuthController) SetPassword(ctx *gin.Context) {
	var setPassword models.SetPassword
	if err := ctx.ShouldBindJSON(&setPassword); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(setPassword); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	userService := services.UserService{}
	err := userService.SetPassword(&setPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, dtos.Response{Code: http.StatusOK, Message: "Password set successfully"})
	}
}
