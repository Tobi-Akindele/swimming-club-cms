package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
)

type CompetitionController struct{}

func (cc *CompetitionController) CreateCompetition(ctx *gin.Context) {
	var competition models.CreateCompetition
	if err := ctx.ShouldBindJSON(&competition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(competition); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	competitionService := services.CompetitionService{}
	createdCompetition, err := competitionService.CreateCompetition(&competition)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdCompetition)
	}
}

func (cc *CompetitionController) GetCompetitionById(ctx *gin.Context) {
	competitionId := ctx.Param("id")
	competitionService := services.CompetitionService{}
	competition, err := competitionService.GetById(competitionId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, competition)
	}
}
