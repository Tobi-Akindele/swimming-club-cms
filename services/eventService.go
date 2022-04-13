package services

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"strings"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type eventService struct{}

func (es *eventService) CreateEvent(eventDto *models.CreateEvent) (*models.Event, error) {
	event := models.Event{}
	competitionService := GetServiceManagerInstance().GetCompetitionService()
	rawCompetition, _ := competitionService.GetById(eventDto.CompetitionId, false)
	if rawCompetition == nil {
		return nil, errors.New("competition not found")
	}
	competition, _ := rawCompetition.(*models.Competition)
	err := copier.Copy(&event, eventDto)
	if err != nil {
		return nil, err
	}
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	savedEvent, err := eventRepository.SaveEvent(&event)
	if err != nil {
		return nil, err
	}
	competition.Events = append(competition.Events, &mogo.RefField{ID: savedEvent.ID})
	_, _ = competitionService.UpdateCompetition(competition)
	return savedEvent, nil
}

func (es *eventService) GetById(id string, fetchRelationships bool) (interface{}, error) {
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	return eventRepository.FindById(id, fetchRelationships)
}

func (es *eventService) AddParticipant(participant *models.AddParticipant) (*models.Event, error) {
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	rawEvent, _ := eventRepository.FindById(participant.EventId, false)
	if rawEvent == nil {
		return nil, errors.New("event not found")
	}
	event, _ := rawEvent.(*models.Event)
	existingParticipants := utils.ConvertRefFieldSliceToStringMap(event.Participants)
	serviceManager := GetServiceManagerInstance()
	userService := serviceManager.GetUserService()
	rawUser, _ := userService.GetById(participant.ParticipantId, false)
	if rawUser == nil {
		return nil, errors.New("all participants must be registered")
	}
	user, _ := rawUser.(*models.User)
	if !user.Active {
		return nil, errors.New("all participants must be active")
	}
	userTypeService := serviceManager.GetUserTypeService()
	userType, _ := userTypeService.GetById(user.UserType.ID.Hex())
	if utils.SWIMMER != userType.Name {
		return nil, errors.New("all participants must be swimmers")
	}
	if !utils.MapContainsKey(existingParticipants, user.ID.Hex()) {
		event.Participants = append(event.Participants, &mogo.RefField{ID: user.ID})
	}

	return eventRepository.SaveEvent(event)
}

func (es *eventService) GetByName(eventByName *models.EventByName) (*models.Event, error) {
	competitionService := GetServiceManagerInstance().GetCompetitionService()
	rawCompetition, _ := competitionService.GetById(eventByName.CompetitionId, true)
	if rawCompetition == nil {
		return nil, errors.New("competition not found")
	}
	competition, _ := rawCompetition.(*models.CompetitionResult)
	if len(competition.Events) > 0 {
		for _, event := range competition.Events {
			if strings.TrimSpace(eventByName.Name) == event.Name {
				return &event, nil
			}
		}
	}

	return nil, errors.New("event not found")
}

func (es *eventService) RemoveParticipantsFromEvent(removeParticipants *models.RemoveParticipants) (*models.Event, error) {
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	rawEvent, _ := eventRepository.FindById(removeParticipants.EventId, false)
	if rawEvent == nil {
		return nil, errors.New("event not found")
	}
	event, _ := rawEvent.(*models.Event)
	event.Participants = utils.RemoveRefFromRefSlice(event.Participants, removeParticipants.ParticipantIds)
	return eventRepository.SaveEvent(event)
}

func (es *eventService) RecordResults(recordResult *models.RecordResult) (*models.Event, error) {
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	resultRepository := repositories.GetRepositoryManagerInstance().GetResultRepository()
	rawEvent, _ := eventRepository.FindById(recordResult.EventId, false)
	if rawEvent == nil {
		return nil, errors.New("event not found")
	}
	event, _ := rawEvent.(*models.Event)
	participants := utils.ConvertRefFieldSliceToStringMap(event.Participants)
	var results []*models.Result
	for _, resultData := range recordResult.Results {
		if !utils.MapContainsKey(participants, resultData.ParticipantId) {
			return nil, errors.New("only registered participants can have result")
		}
		var result models.Result
		if len(resultData.ResultId) > 0 {
			existingResult, _ := resultRepository.FindById(resultData.ResultId)
			if existingResult != nil {
				result = *existingResult
			}
		}
		result.Time = resultData.Time
		result.FinalPoint = resultData.FinalPoint
		result.Participant = mogo.RefField{ID: bson.ObjectIdHex(resultData.ParticipantId)}
		results = append(results, &result)
	}
	requireUpdate := false
	for _, result := range results {
		result, err := resultRepository.SaveResult(result)
		if err == nil {
			if !utils.MapContainsKey(utils.ConvertRefFieldSliceToStringMap(event.Results), result.ID.Hex()) {
				requireUpdate = true
				event.Results = append(event.Results, &mogo.RefField{ID: result.ID})
			}
		}
	}
	if requireUpdate {
		return eventRepository.SaveEvent(event)
	}
	return event, nil
}
