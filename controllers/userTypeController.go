package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
)

type UserTypeController struct{}

func (utc *UserTypeController) CreateUserType(ctx *gin.Context) {
	var userType models.UserTypeDto
	if err := ctx.ShouldBindJSON(&userType); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(userType); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	userTypeService := services.UserTypeService{}
	createdUserType, err := userTypeService.CreateUserType(&userType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdUserType)
	}
}

func (utc *UserTypeController) GetAllUserTypes(ctx *gin.Context) {
	userTypeService := services.UserTypeService{}
	userTypes, err := userTypeService.GetAllUserTypes()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userTypes)
	}
}
