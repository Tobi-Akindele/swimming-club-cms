package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type UserRepository struct{}

func (ur *UserRepository) SaveUser(user *models.User) (*models.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userModel := mogo.NewDoc(user).(*models.User)
	err := mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return userModel, err
}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"username": username}, userDoc)
	if err != nil {
		return nil, err
	}
	return userDoc, nil
}

func (ur *UserRepository) FindById(id string) (*models.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, userDoc)
	if err != nil {
		return nil, err
	}
	return userDoc, nil
}
