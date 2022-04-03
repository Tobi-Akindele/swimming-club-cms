package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/services"
)

type UserController struct{}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	serviceManager := services.GetServiceManagerInstance()
	users, err := serviceManager.GetUserService().GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (uc *UserController) GetByUsername(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	serviceManager := services.GetServiceManagerInstance()
	users, err := serviceManager.GetUserService().GetByUsername(username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (uc *UserController) GetByEmail(ctx *gin.Context) {
	email := ctx.GetHeader("email")
	serviceManager := services.GetServiceManagerInstance()
	users, err := serviceManager.GetUserService().GetByEmail(email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}
