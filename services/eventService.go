package services

import (
	"errors"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type EventService struct{}

func (es *EventService) CreateEvent(eventDto *models.CreateEvent) (*models.Event, error) {
	event := models.Event{}
	competitionService := CompetitionService{}
	competition, _ := competitionService.GetById(eventDto.CompetitionId)
	if competition == nil {
		return nil, errors.New("competition not found")
	}
	err := copier.Copy(&event, eventDto)
	if err != nil {
		return nil, err
	}
	eventRepository := repositories.EventRepository{}
	savedEvent, err := eventRepository.SaveEvent(&event)
	if err != nil {
		return nil, err
	}
	competition.Events = append(competition.Events, &mogo.RefField{ID: savedEvent.ID})
	_, _ = competitionService.UpdateCompetition(competition)
	return savedEvent, nil
}

func (es *EventService) GetById(id string) (*models.Event, error) {
	eventRepository := repositories.EventRepository{}
	return eventRepository.FindById(id)
}

func (es *EventService) AddParticipants(participants *models.AddParticipants) (*models.Event, error) {
	eventRepository := repositories.EventRepository{}
	event, _ := eventRepository.FindById(participants.EventId)
	if event == nil {
		return nil, errors.New("event not found")
	}
	existingParticipants := utils.ConvertRefFieldSliceToStringMap(event.Participants)
	userService := UserService{}
	for idx := range participants.Participants {
		user, _ := userService.GetById(participants.Participants[idx])
		if user == nil {
			return nil, errors.New("all participants must be registered on the system")
		}
		if user.UserType.Name != utils.SWIMMER {
			return nil, errors.New("all participants must be swimmers")
		}
		if !utils.MapContainsKey(existingParticipants, user.ID.String()) {
			event.Participants = append(event.Participants, &mogo.RefField{ID: user.ID})
		}
	}
	return eventRepository.SaveEvent(event)
}
