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

type EventController struct{}

func (ec *EventController) CreateEvent(ctx *gin.Context) {
	var event models.CreateEvent
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(event); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	serviceManager := services.GetServiceManagerInstance()
	createdEvent, err := serviceManager.GetEventService().CreateEvent(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdEvent)
	}
}

func (ec *EventController) GetEventById(ctx *gin.Context) {
	eventId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	rawEvent, err := serviceManager.GetEventService().GetById(eventId, true)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		userService := services.GetServiceManagerInstance().GetUserService()
		event, _ := rawEvent.(*models.EventResult)
		participants, _ := userService.GetByIds(utils.ExtractUserIdsFromUserStructs(event.Participants))
		if event.Results != nil && participants != nil {
			for _, result := range event.Results {
				for i := range participants {
					if participants[i].ID.Hex() == result.Participant.ID.Hex() {
						participants[i].Time = result.Time
						participants[i].FinalPoint = result.FinalPoint
						break
					}
				}
			}
		}
		event.Participants = participants
		ctx.JSON(http.StatusOK, event)
	}
}

func (ec *EventController) AddParticipants(ctx *gin.Context) {
	var participant models.AddParticipant
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
	event, err := serviceManager.GetEventService().AddParticipant(&participant)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, event)
	}
}

func (ec *EventController) GetEventByName(ctx *gin.Context) {
	var eventByName models.EventByName
	if err := ctx.ShouldBindJSON(&eventByName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	if errs := validator.Validate(eventByName); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
			Code: http.StatusBadRequest, Message: errs.Error(),
		})
		return
	}
	eventService := services.GetServiceManagerInstance().GetEventService()
	event, err := eventService.GetByName(&eventByName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, event)
	}
}

func (ec *EventController) RemoveParticipantsFromEvent(ctx *gin.Context) {
	var removeParticipants models.RemoveParticipants
	if err := ctx.ShouldBindJSON(&removeParticipants); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(removeParticipants); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	serviceManager := services.GetServiceManagerInstance()
	event, err := serviceManager.GetEventService().RemoveParticipantsFromEvent(&removeParticipants)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, event)
	}
}

func (ec *EventController) RecordResults(ctx *gin.Context) {
	var recordResult models.RecordResult
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
	event, err := serviceManager.GetEventService().RecordResults(&recordResult)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, event)
	}
}
