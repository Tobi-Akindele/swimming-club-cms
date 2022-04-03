package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type competitionRepository struct{}

func (cr *competitionRepository) FindById(id string) (*models.Competition, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	competitionDoc := mogo.NewDoc(models.Competition{}).(*models.Competition)
	err := competitionDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, competitionDoc)
	if err != nil {
		return nil, err
	}
	return competitionDoc, nil
}

func (cr *competitionRepository) SaveCompetition(competition *models.Competition) (*models.Competition, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	competitionModel := mogo.NewDoc(competition).(*models.Competition)
	err := mogo.Save(competitionModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return competitionModel, err
}

func (cr *competitionRepository) FindAll() ([]*models.CompetitionResult, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	competitionDoc := mogo.NewDoc(models.Competition{}).(*models.Competition)
	var competitions []*models.Competition
	err := competitionDoc.Find(nil).Q().Sort("-_created").All(&competitions)
	if err != nil {
		return nil, err
	}
	result := make([]*models.CompetitionResult, len(competitions))
	for i, competition := range competitions {
		var competitionResult models.CompetitionResult
		var events []models.Event
		u := mogo.NewDoc(competition).(*models.Competition)
		_ = u.Populate("Events").Find(bson.M{}).All(&events)
		_ = copier.Copy(&competitionResult, competition)
		if len(events) > 0 {
			competitionResult.Events = events
		}
		result[i] = &competitionResult
	}
	return result, nil
}
