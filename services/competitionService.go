package services

import (
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
)

type CompetitionService struct{}

func (cs *CompetitionService) CreateCompetition(competitionDto *models.CreateCompetition) (*models.Competition, error) {
	competition := models.Competition{}
	err := copier.Copy(&competition, competitionDto)
	if err != nil {
		return nil, err
	}
	competition.Status = 0
	competitionRepository := repositories.CompetitionRepository{}
	return competitionRepository.SaveCompetition(&competition)
}

func (cs *CompetitionService) GetById(id string) (*models.Competition, error) {
	competitionRepository := repositories.CompetitionRepository{}
	return competitionRepository.FindById(id)
}

func (cs *CompetitionService) UpdateCompetition(competition *models.Competition) (*models.Competition, error) {
	competitionRepository := repositories.CompetitionRepository{}
	return competitionRepository.SaveCompetition(competition)
}
