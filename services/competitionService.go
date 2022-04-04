package services

import (
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

func (cs *competitionService) GetById(id string) (*models.Competition, error) {
	competitionRepository := repositories.GetRepositoryManagerInstance().GetCompetitionRepository()
	return competitionRepository.FindById(id)
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
		competition, _ := competitionRepository.FindById(competitionId)
		if competition != nil {
			deletions = append(deletions, competition)
		}
	}
	return competitionRepository.DeleteCompetitions(deletions)
}
