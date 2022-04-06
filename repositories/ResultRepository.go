package repositories

import (
	"github.com/goonode/mogo"
	"swimming-club-cms-be/configs/db"
	"swimming-club-cms-be/models"
)

type resultRepository struct{}

func (rr *resultRepository) SaveResult(result *models.Result) (*models.Result, error) {
	conn := db.GetConnection()
	defer db.CloseConnection(conn)

	resultModel := mogo.NewDoc(result).(*models.Result)
	err := mogo.Save(resultModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return nil, vErr
	}
	return resultModel, err
}
