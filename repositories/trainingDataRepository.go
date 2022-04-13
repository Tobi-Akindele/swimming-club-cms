package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type trainingDataRepository struct{}

func (tdr *trainingDataRepository) SaveTrainingData(trainingData *models.TrainingData) (*models.TrainingData, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	trainingDataModel := mogo.NewDoc(trainingData).(*models.TrainingData)
	err := mogo.Save(trainingDataModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return trainingDataModel, err
}

func (tdr *trainingDataRepository) FindByClubId(clubId string) ([]*models.TrainingData, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	tdDoc := mogo.NewDoc(models.TrainingData{}).(*models.TrainingData)
	var trainingData []*models.TrainingData
	err := tdDoc.Find(bson.M{"clubid._id": bson.ObjectIdHex(clubId)}).Q().Sort("_created").All(&trainingData)
	if err != nil {
		return nil, err
	}
	return trainingData, nil
}

func (tdr *trainingDataRepository) FindById(id string, fetchRelationships bool) (interface{}, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	tdDoc := mogo.NewDoc(models.TrainingData{}).(*models.TrainingData)
	err := tdDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, tdDoc)
	if err != nil {
		return nil, err
	}
	if fetchRelationships {
		var trainingDataResult models.TrainingDataResult
		var participants []models.UserResult
		var results []models.Result
		_ = tdDoc.Populate("Participants").Find(bson.M{}).All(&participants)
		_ = tdDoc.Populate("Results").Find(bson.M{}).All(&results)
		_ = copier.Copy(&trainingDataResult, tdDoc)
		trainingDataResult.Participants = participants
		trainingDataResult.Results = results

		return &trainingDataResult, nil
	}
	return tdDoc, nil
}
