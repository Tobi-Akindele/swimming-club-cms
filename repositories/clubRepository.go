package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
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

func (cr *clubRepository) FindById(id string, fetchRelationShips bool) (interface{}, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, clubDoc)
	if err != nil {
		return nil, err
	}
	if fetchRelationShips {
		var clubResult models.ClubResult
		var coach []models.User
		var members []models.User
		_ = clubDoc.Populate("Coach").Find(bson.M{}).All(&coach)
		_ = clubDoc.Populate("Members").Find(bson.M{}).All(&members)
		_ = copier.Copy(&clubResult, clubDoc)
		if len(coach) > 0 {
			clubResult.Coach = coach[0]
		}
		clubResult.Members = members

		return &clubResult, nil
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

func (cr *clubRepository) FindAll() ([]*models.ClubResult, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	var clubs []*models.Club
	err := clubDoc.Find(nil).Q().Sort("-_created").All(&clubs)
	if err != nil {
		return nil, err
	}
	result := make([]*models.ClubResult, len(clubs))
	for i, club := range clubs {
		var clubResult models.ClubResult
		var coach []models.User
		var members []models.User
		u := mogo.NewDoc(club).(*models.Club)
		_ = u.Populate("Coach").Find(bson.M{}).All(&coach)
		_ = u.Populate("Members").Find(bson.M{}).All(&members)
		_ = copier.Copy(&clubResult, club)
		if len(coach) > 0 {
			clubResult.Coach = coach[0]
		}
		clubResult.Members = members
		result[i] = &clubResult
	}
	return result, nil
}

func (cr *clubRepository) FindByName(name string) (*models.Club, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"name": name}, clubDoc)
	if err != nil {
		return nil, err
	}
	return clubDoc, nil
}

func (cr *clubRepository) FindByCoachId(coachId string) (*models.Club, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	clubDoc := mogo.NewDoc(models.Club{}).(*models.Club)
	err := clubDoc.FindOne(bson.M{"coach._id": bson.ObjectIdHex(coachId)}, clubDoc)
	if err != nil {
		return nil, err
	}
	return clubDoc, nil
}
