package services

import (
	"errors"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
	"time"
)

type competitionService struct{}

func (cs *competitionService) CreateCompetition(competitionDto *models.CreateCompetition) (*models.Competition, error) {
	competition := models.Competition{}
	err := copier.Copy(&competition, competitionDto)
	if err != nil {
		return nil, err
	}
	parseDate, err := time.Parse(utils.DATE_FORMAT, competitionDto.Date)
	competition.Date = parseDate
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.SaveCompetition(&competition)
}

func (cs *competitionService) GetById(id string, fetchRelationship bool) (interface{}, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.FindById(id, fetchRelationship)
}

func (cs *competitionService) UpdateCompetition(competition *models.Competition) (*models.Competition, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.SaveCompetition(competition)
}

func (cs *competitionService) GetAllCompetitions() ([]*models.CompetitionResult, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.FindAll()
}

func (cs *competitionService) GetByName(name string) (*models.Competition, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.FindByName(name)
}

func (cs *competitionService) DeleteCompetitions(deleteCompetition *models.DeleteCompetition) []error {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	var deletions []*models.Competition
	for _, competitionId := range deleteCompetition.CompetitionIds {
		rawCompetition, _ := competitionRepository.FindById(competitionId, false)
		if rawCompetition != nil {
			competition, _ := rawCompetition.(*models.Competition)
			deletions = append(deletions, competition)
		}
	}
	return competitionRepository.DeleteCompetitions(deletions)
}

func (cs *competitionService) RemoveEventFromCompetition(removeEvents *models.RemoveEvents) (*models.Competition, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	rawCompetition, _ := competitionRepository.FindById(removeEvents.CompetitionId, false)
	if rawCompetition == nil {
		return nil, errors.New("competition not found")
	}
	competition, _ := rawCompetition.(*models.Competition)
	competition.Events = utils.RemoveRefFromRefSlice(competition.Events, removeEvents.EventIds)
	return competitionRepository.SaveCompetition(competition)
}

func (cs *competitionService) GetTotalCompetitions() (*int, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.FindAllCompetitionsCount()
}

func (cs *competitionService) GetOpenCompetitionsCount() (*int, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.FindAllOpenCompetitionsCount()
}
