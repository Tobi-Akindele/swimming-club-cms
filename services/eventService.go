package services

import (
	"errors"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type eventService struct{}

func (es *eventService) CreateEvent(eventDto *models.CreateEvent) (*models.Event, error) {
	event := models.Event{}
	competitionService := GetServiceManagerInstance().GetCompetitionService()
	competition, _ := competitionService.GetById(eventDto.CompetitionId)
	if competition == nil {
		return nil, errors.New("competition not found")
	}
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

func (es *eventService) GetById(id string) (*models.Event, error) {
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	return eventRepository.FindById(id)
}

func (es *eventService) AddParticipants(participants *models.AddParticipants) (*models.Event, error) {
	eventRepository := repositories.GetRepositoryManagerInstance().GetEventRepository()
	event, _ := eventRepository.FindById(participants.EventId)
	if event == nil {
		return nil, errors.New("event not found")
	}
	existingParticipants := utils.ConvertRefFieldSliceToStringMap(event.Participants)
	userService := GetServiceManagerInstance().GetUserService()
	for idx := range participants.Participants {
		rawUser, _ := userService.GetById(participants.Participants[idx], true)
		if rawUser == nil {
			return nil, errors.New("all participants must be registered on the system")
		}
		user, _ := rawUser.(*models.UserResult)
		if user.UserType.Name != utils.SWIMMER {
			return nil, errors.New("all participants must be swimmers")
		}
		if !utils.MapContainsKey(existingParticipants, user.ID.String()) {
			event.Participants = append(event.Participants, &mogo.RefField{ID: user.ID})
		}
	}
	return eventRepository.SaveEvent(event)
}
