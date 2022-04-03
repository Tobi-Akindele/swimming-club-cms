package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type clubRepository struct{}

func (cr *clubRepository) SaveClub(club *models.Club) (*models.Club, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubModel := mogo.NewDoc(club).(*models.Club)
	err := mogo.Save(clubModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return clubModel, err
}

func (cr *clubRepository) FindById(id string) (*models.Club, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, clubDoc)
	if err != nil {
		return nil, err
	}
	return clubDoc, nil
}

func (cr *clubRepository) FindByMemberId(memberId string) (*models.Club, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"members._id": bson.ObjectIdHex(memberId)}, clubDoc)
	if err != nil {
		return nil, err
	}
	return clubDoc, nil
}

func (cr *clubRepository) FindAll() ([]*models.Club, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	var results []*models.Club
	err := clubDoc.Find(nil).Q().Sort("-_created").All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
