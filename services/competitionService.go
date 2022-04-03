package services

import (
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type competitionService struct{}

func (cs *competitionService) CreateCompetition(competitionDto *models.CreateCompetition) (*models.Competition, error) {
	competition := models.Competition{}
	err := copier.Copy(&competition, competitionDto)
	if err != nil {
		return nil, err
	}
	competition.Status = 0
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
