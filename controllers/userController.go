package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"strings"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
	"swimming-club-cms-be/utils"
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

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.UserUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	_ = validator.SetValidationFunc("datetime", utils.Datetime)
	if errs := validator.Validate(user); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	updatedUser, err := serviceManager.GetUserService().UpdateUser(&user, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, updatedUser)
	}
}

func (uc *UserController) SearchUsersByUserType(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	userType := ctx.GetHeader("userType")
	if len(userType) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: "User type is required"})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	users, err := serviceManager.GetUserService().SearchUsersByUserType(username, strings.ToUpper(userType))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	userService := serviceManager.GetUserService()
	rawUser, err := userService.GetById(userId, true)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		user, _ := rawUser.(*models.UserResult)
		clubService := serviceManager.GetClubService()
		club, _ := clubService.GetByCoachId(user.ID.Hex())
		if club == nil {
			club, _ = clubService.GetByMemberId(user.ID.Hex())
			if club != nil {
				user.Club = *club
			}
		} else {
			user.Club = *club
		}
		parents, _ := userService.GetUsersByChildId(user.ID.Hex())
		user.Parents = parents
		ctx.JSON(http.StatusOK, rawUser)
	}
}

func (uc *UserController) AddChild(ctx *gin.Context) {
	var child models.AddRelation
	if err := ctx.ShouldBindJSON(&child); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(child); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	user, err := serviceManager.GetUserService().AddChild(&child)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, user)
	}
}
