package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/services"
)

type UserController struct{}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	userService := services.UserService{}
	users, err := userService.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (uc *UserController) GetByUsername(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	userService := services.UserService{}
	users, err := userService.GetByUsername(username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (uc *UserController) GetByEmail(ctx *gin.Context) {
	email := ctx.GetHeader("email")
	userService := services.UserService{}
	users, err := userService.GetByEmail(email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}
