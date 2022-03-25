package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type UserTypeRepository struct{}

func (utr *UserTypeRepository) SaveUserType(userType *models.UserType) (*models.UserType, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userTypeModel := mogo.NewDoc(userType).(*models.UserType)
	err := mogo.Save(userTypeModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return userTypeModel, err
}

func (utr *UserTypeRepository) SaveUserTypes(userTypes []*models.UserType) ([]*models.UserType, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	for u := range userTypes {
		userTypeModel := mogo.NewDoc(userTypes[u]).(*models.UserType)
		err := mogo.Save(userTypeModel)
		if _, ok := err.(*mogo.ValidationError); ok {
			continue
		}
		userTypes[u] = userTypeModel
	}
	return userTypes, nil
}

func (utr *UserTypeRepository) FindByName(name string) (*models.UserType, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userTypeDoc := mogo.NewDoc(models.UserType{}).(*models.UserType)
	err := userTypeDoc.FindOne(bson.M{"name": name}, userTypeDoc)
	if err != nil {
		return nil, err
	}
	return userTypeDoc, nil
}

func (utr *UserTypeRepository) FindById(id string) (*models.UserType, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userTypeDoc := mogo.NewDoc(models.UserType{}).(*models.UserType)
	err := userTypeDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, userTypeDoc)
	if err != nil {
		return nil, err
	}
	return userTypeDoc, nil
}

func (utr *UserTypeRepository) FindAll() ([]*models.UserType, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userTypeDoc := mogo.NewDoc(models.UserType{}).(*models.UserType)
	var results []*models.UserType
	err := userTypeDoc.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
