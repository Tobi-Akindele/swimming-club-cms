package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type userRepository struct{}

func (ur *userRepository) SaveUser(user *models.User) (*models.User, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	userModel := mogo.NewDoc(user).(*models.User)
	err := mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return userModel, err
}

func (ur *userRepository) FindByUsername(username string) (*models.UserResult, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

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
	if len(userType) > 0 {
		userResult.UserType = userType[0]
	}
	if len(role) > 0 {
		userResult.Role = role[0]
	}

	return &userResult, nil
}

func (ur *userRepository) FindByEmail(username string) (*models.UserResult, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

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
	if len(userType) > 0 {
		userResult.UserType = userType[0]
	}
	if len(role) > 0 {
		userResult.Role = role[0]
	}

	return &userResult, nil
}

func (ur *userRepository) FindById(id string) (*models.UserResult, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, userDoc)
	if err != nil {
		return nil, err
	}
	var userResult models.UserResult
	var userType []models.UserType
	var role []models.Role
	err = userDoc.Populate("UserType").Find(bson.M{}).All(&userType)
	err = userDoc.Populate("Role").Find(bson.M{}).All(&role)

	_ = copier.Copy(&userResult, userDoc)
	if len(userType) > 0 {
		userResult.UserType = userType[0]
	}
	if len(role) > 0 {
		userResult.Role = role[0]
	}

	return &userResult, nil
}

func (ur *userRepository) FindAll() ([]*models.UserResult, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	var users []*models.User
	err := userDoc.Find(nil).Q().Sort("-_created").All(&users)
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
		if len(userType) > 0 {
			userResult.UserType = userType[0]
		}
		if len(role) > 0 {
			userResult.Role = role[0]
		}
		result[idx] = &userResult
	}
	return result, nil
}

func (ur *userRepository) FindByActivationCode(activationCode string) (*models.User, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	userDoc := mogo.NewDoc(models.User{}).(*models.User)
	err := userDoc.FindOne(bson.M{"activationcode": activationCode}, userDoc)
	if err != nil {
		return nil, err
	}

	return userDoc, nil
}
