package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
)

type ClubController struct{}

func (cc *ClubController) CreateClub(ctx *gin.Context) {
	var club models.ClubDto
	if err := ctx.ShouldBindJSON(&club); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(club); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	createdClub, err := serviceManager.GetClubService().CreateClub(&club)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdClub)
	}
}

func (cc *ClubController) AddMembers(ctx *gin.Context) {
	var newMember models.AddMember
	if err := ctx.ShouldBindJSON(&newMember); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(newMember); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	club, err := serviceManager.GetClubService().AddMember(&newMember)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, club)
	}
}

func (cc *ClubController) GetClubById(ctx *gin.Context) {
	clubId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	club, err := serviceManager.GetClubService().GetById(clubId, true)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, club)
	}
}

func (cc *ClubController) GetAllClubs(ctx *gin.Context) {
	serviceManager := services.GetServiceManagerInstance()
	clubs, err := serviceManager.GetClubService().GetAllClubs()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, clubs)
	}
}

func (cc *ClubController) GetClubByName(ctx *gin.Context) {
	name := ctx.GetHeader("name")
	serviceManager := services.GetServiceManagerInstance()
	club, err := serviceManager.GetClubService().GetByName(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, club)
	}
}
