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
	serviceManager := services.GetServiceManagerInstance()
	createdCompetition, err := serviceManager.GetCompetitionService().CreateCompetition(&competition)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdCompetition)
	}
}

func (cc *CompetitionController) GetCompetitionById(ctx *gin.Context) {
	competitionId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	competition, err := serviceManager.GetCompetitionService().GetById(competitionId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, competition)
	}
}

func (cc *CompetitionController) GetAllCompetitions(ctx *gin.Context) {
	serviceManager := services.GetServiceManagerInstance()
	competitions, err := serviceManager.GetCompetitionService().GetAllCompetitions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, competitions)
	}
}
