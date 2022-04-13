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

type CompetitionController struct{}

func (cc *CompetitionController) CreateCompetition(ctx *gin.Context) {
	var competition models.CreateCompetition
	if err := ctx.ShouldBindJSON(&competition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	_ = validator.SetValidationFunc("datetime", utils.Datetime)
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
	competition, err := serviceManager.GetCompetitionService().GetById(competitionId, false)
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

func (cc *CompetitionController) GetByName(ctx *gin.Context) {
	name := ctx.GetHeader("name")
	serviceManager := services.GetServiceManagerInstance()
	competition, err := serviceManager.GetCompetitionService().GetByName(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, competition)
	}
}

func (cc *CompetitionController) DeleteCompetitions(ctx *gin.Context) {
	var competitions models.DeleteCompetition
	if err := ctx.ShouldBindJSON(&competitions); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(competitions); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	serviceManager := services.GetServiceManagerInstance()
	err := serviceManager.GetCompetitionService().DeleteCompetitions(&competitions)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: "An error occurred"})
	} else {
		ctx.JSON(http.StatusOK, dtos.Response{
			Code:    http.StatusOK,
			Message: "Deletion successful",
		})
	}
}

func (cc *CompetitionController) RemoveEventFromCompetition(ctx *gin.Context) {
	var removeEvents models.RemoveEvents
	if err := ctx.ShouldBindJSON(&removeEvents); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(removeEvents); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	serviceManager := services.GetServiceManagerInstance()
	competition, err := serviceManager.GetCompetitionService().RemoveEventFromCompetition(&removeEvents)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, competition)
	}
}

func (cc *CompetitionController) GetTotalCompetitions(ctx *gin.Context) {
	serviceManager := services.GetServiceManagerInstance()
	total, err := serviceManager.GetCompetitionService().GetTotalCompetitions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, dtos.Response{Code: http.StatusBadRequest, Message: "", Count: *total})
	}
}

func (cc *CompetitionController) GetOpenCompetitionsCount(ctx *gin.Context) {
	serviceManager := services.GetServiceManagerInstance()
	total, err := serviceManager.GetCompetitionService().GetOpenCompetitionsCount()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, dtos.Response{Code: http.StatusBadRequest, Message: "", Count: *total})
	}
}
