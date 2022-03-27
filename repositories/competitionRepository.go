package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type CompetitionRepository struct{}

func (cr *CompetitionRepository) FindById(id string) (*models.Competition, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	competitionDoc := mogo.NewDoc(models.Competition{}).(*models.Competition)
	err := competitionDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, competitionDoc)
	if err != nil {
		return nil, err
	}
	return competitionDoc, nil
}

func (cr *CompetitionRepository) SaveCompetition(competition *models.Competition) (*models.Competition, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	competitionModel := mogo.NewDoc(competition).(*models.Competition)
	err := mogo.Save(competitionModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return competitionModel, err
}
