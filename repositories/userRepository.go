package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
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

func (ur *UserRepository) FindByUsername(username string) (*models.UserResult, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"username": username}, userDoc)
	if err != nil {
		return nil, err
	}
	var userResult models.UserResult
	var userType []models.UserType
	var role []models.Role
	_ = userDoc.Populate("UserType").Find(bson.M{}).All(&userType)
	_ = userDoc.Populate("Role").Find(bson.M{}).All(&role)

	_ = copier.Copy(&userResult, userDoc)
	userResult.UserType = userType[0]
	userResult.Role = role[0]

	return &userResult, nil
}

func (ur *UserRepository) FindByEmail(username string) (*models.UserResult, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"username": username}, userDoc)
	if err != nil {
		return nil, err
	}
	var userResult models.UserResult
	var userType []models.UserType
	var role []models.Role
	_ = userDoc.Populate("UserType").Find(bson.M{}).All(&userType)
	_ = userDoc.Populate("Role").Find(bson.M{}).All(&role)

	_ = copier.Copy(&userResult, userDoc)
	userResult.UserType = userType[0]
	userResult.Role = role[0]

	return &userResult, nil
}

func (ur *UserRepository) FindById(id string) (*models.UserResult, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, userDoc)
	if err != nil {
		return nil, err
	}
	var userResult models.UserResult
	var userType []models.UserType
	var role []models.Role
	_ = userDoc.Populate("UserType").Find(bson.M{}).All(&userType)
	_ = userDoc.Populate("Role").Find(bson.M{}).All(&role)

	_ = copier.Copy(&userResult, userDoc)
	userResult.UserType = userType[0]
	userResult.Role = role[0]

	return &userResult, nil
}

func (ur *UserRepository) FindAll() ([]*models.UserResult, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	var users []*models.User
	err := userDoc.Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	result := make([]*models.UserResult, len(users))
	for idx := range users {
		var userResult models.UserResult
		var userType []models.UserType
		var role []models.Role
		u := mogo.NewDoc(users[idx]).(*models.User)
		_ = u.Populate("UserType").Find(bson.M{}).All(&userType)
		_ = u.Populate("Role").Find(bson.M{}).All(&role)
		_ = copier.Copy(&userResult, users[idx])
		userResult.UserType = userType[0]
		userResult.Role = role[0]
		result[idx] = &userResult
	}
	return result, nil
}

func (ur *UserRepository) FindByActivationCode(activationCode string) (*models.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"activationcode": activationCode}, userDoc)
	if err != nil {
		return nil, err
	}

	return userDoc, nil
}
