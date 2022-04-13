package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
)

type TrainingDataController struct{}

func (tdc *TrainingDataController) CreateTrainingData(ctx *gin.Context) {
	var createTD models.CreateTrainingData
	if err := ctx.ShouldBindJSON(&createTD); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(createTD); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	createdTD, err := serviceManager.GetTrainingDataService().CreateTrainingData(&createTD, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdTD)
	}
}

func (tdc *TrainingDataController) GeTrainingDataById(ctx *gin.Context) {
	trainingDataId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	rawTD, err := serviceManager.GetTrainingDataService().GetById(trainingDataId, true)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		td, _ := rawTD.(*models.TrainingDataResult)
		if td.Participants != nil && td.Results != nil {
			for _, result := range td.Results {
				for i := range td.Participants {
					if td.Participants[i].ID.Hex() == result.Participant.ID.Hex() {
						td.Participants[i].Time = result.Time
						td.Participants[i].FinalPoint = result.FinalPoint
						td.Participants[i].ResultId = result.ID.Hex()
						break
					}
				}
			}
		}
		ctx.JSON(http.StatusOK, rawTD)
	}
}

func (tdc *TrainingDataController) AddParticipants(ctx *gin.Context) {
	var participant models.AddTDParticipants
	if err := ctx.ShouldBindJSON(&participant); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(participant); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	trainingData, err := serviceManager.GetTrainingDataService().AddParticipants(&participant, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, trainingData)
	}
}

func (tdc *TrainingDataController) RecordTDResults(ctx *gin.Context) {
	var recordResult models.RecordTDResult
	if err := ctx.ShouldBindJSON(&recordResult); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(recordResult); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	trainingData, err := serviceManager.GetTrainingDataService().RecordTDResults(&recordResult, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, trainingData)
	}
}
