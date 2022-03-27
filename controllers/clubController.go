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
	clubService := services.ClubService{}
	createdClub, err := clubService.CreateClub(&club)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdClub)
	}
}

func (cc *ClubController) AddMembers(ctx *gin.Context) {
	var newMembers models.AddMember
	if err := ctx.ShouldBindJSON(&newMembers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(newMembers); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	clubService := services.ClubService{}
	club, err := clubService.AddMembers(&newMembers)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, club)
	}
}

func (cc *ClubController) GetClubById(ctx *gin.Context) {
	clubId := ctx.Param("id")
	clubService := services.ClubService{}
	club, err := clubService.GetById(clubId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, club)
	}
}

func (cc *ClubController) GetAllClubs(ctx *gin.Context) {
	clubService := services.ClubService{}
	clubs, err := clubService.GetAllClubs()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, clubs)
	}
}
