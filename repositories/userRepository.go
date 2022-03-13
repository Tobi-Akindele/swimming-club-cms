package repositories

import (
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
