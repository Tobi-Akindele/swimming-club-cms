package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type ClubRepository struct{}

func (cr *ClubRepository) SaveClub(club *models.Club) (*models.Club, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	clubModel := mogo.NewDoc(club).(*models.Club)
	err := mogo.Save(clubModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return clubModel, err
}

func (cr *ClubRepository) FindById(id string) (*models.Club, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, clubDoc)
	if err != nil {
		return nil, err
	}
	return clubDoc, nil
}

func (cr *ClubRepository) FindByMemberId(memberId string) (*models.Club, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"members._id": bson.ObjectIdHex(memberId)}, clubDoc)
	if err != nil {
		return nil, err
	}
	return clubDoc, nil
}

func (cr *ClubRepository) FindAll() ([]*models.Club, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	var results []*models.Club
	err := clubDoc.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
