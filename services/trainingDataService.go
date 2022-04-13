package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type trainingDataService struct{}

func (tds *trainingDataService) CreateTrainingData(ctd *models.CreateTrainingData, ctx *gin.Context) (*models.TrainingData, error) {
	trainingData := models.TrainingData{}
	clubService := GetServiceManagerInstance().GetClubService()
	rawClub, _ := clubService.GetById(ctd.ClubId, false)
	if rawClub == nil {
		return nil, errors.New("club not found")
	}
	club, _ := rawClub.(*models.Club)
	if requester, isAdmin := utils.IsAdmin(ctx); !isAdmin {
		if requester.ID.Hex() != club.Coach.ID.Hex() {
			return nil, errors.New("action not allowed")
		}
	}
	trainingData.Name = ctd.Name
	trainingData.ClubId = mogo.RefField{ID: club.ID}
	trainingDataRepository := repositories.GetRepositoryManagerInstance().GetTrainingDataRepository()
	return trainingDataRepository.SaveTrainingData(&trainingData)
}

func (tds *trainingDataService) GetByClubId(trainingDataId string) (interface{}, error) {
	trainingDataRepository := repositories.GetRepositoryManagerInstance().GetTrainingDataRepository()
	return trainingDataRepository.FindByClubId(trainingDataId)
}

func (tds *trainingDataService) GetById(trainingDataId string, fetchRelationship bool) (interface{}, error) {
	trainingDataRepository := repositories.GetRepositoryManagerInstance().GetTrainingDataRepository()
	return trainingDataRepository.FindById(trainingDataId, fetchRelationship)
}

func (tds *trainingDataService) AddParticipants(participants *models.AddTDParticipants, ctx *gin.Context) (*models.TrainingData, error) {
	trainingDataRepository := repositories.GetRepositoryManagerInstance().GetTrainingDataRepository()
	rawTD, _ := trainingDataRepository.FindById(participants.TrainingDataId, false)
	if rawTD == nil {
		return nil, errors.New("training data not found")
	}
	td, _ := rawTD.(*models.TrainingData)
	existingParticipants := utils.ConvertRefFieldSliceToStringMap(td.Participants)
	serviceManager := GetServiceManagerInstance()
	userService := serviceManager.GetUserService()
	userTypeService := serviceManager.GetUserTypeService()
	clubService := serviceManager.GetClubService()
	rawClub, _ := clubService.GetById(td.ClubId.ID.Hex(), false)
	if rawClub != nil {
		if requester, isAdmin := utils.IsAdmin(ctx); !isAdmin {
			club, _ := rawClub.(*models.Club)
			if requester.ID.Hex() != club.Coach.ID.Hex() {
				return nil, errors.New("action not allowed")
			}
		}
	}
	for _, participantId := range participants.ParticipantIds {
		rawUser, _ := userService.GetById(participantId, false)
		if rawUser == nil {
			return nil, errors.New("all participants must be registered")
		}
		user, _ := rawUser.(*models.User)
		userType, _ := userTypeService.GetById(user.UserType.ID.Hex())
		if utils.SWIMMER != userType.Name {
			return nil, errors.New("all participants must be swimmers")
		}
		if !utils.MapContainsKey(existingParticipants, user.ID.Hex()) {
			td.Participants = append(td.Participants, &mogo.RefField{ID: user.ID})
		}
	}

	return trainingDataRepository.SaveTrainingData(td)
}

func (tds *trainingDataService) RecordTDResults(recordResults *models.RecordTDResult, ctx *gin.Context) (*models.TrainingData, error) {
	trainingDataRepository := repositories.GetRepositoryManagerInstance().GetTrainingDataRepository()
	resultRepository := repositories.GetRepositoryManagerInstance().GetResultRepository()
	rawTD, _ := trainingDataRepository.FindById(recordResults.TrainingDataId, false)
	if rawTD == nil {
		return nil, errors.New("training data not found")
	}
	td, _ := rawTD.(*models.TrainingData)
	serviceManager := GetServiceManagerInstance()
	clubService := serviceManager.GetClubService()
	rawClub, _ := clubService.GetById(td.ClubId.ID.Hex(), false)
	if rawClub != nil {
		if requester, isAdmin := utils.IsAdmin(ctx); !isAdmin {
			club, _ := rawClub.(*models.Club)
			if requester.ID.Hex() != club.Coach.ID.Hex() {
				return nil, errors.New("action not allowed")
			}
		}
	}
	participants := utils.ConvertRefFieldSliceToStringMap(td.Participants)
	var results []*models.Result
	for _, resultData := range recordResults.Results {
		if !utils.MapContainsKey(participants, resultData.ParticipantId) {
			return nil, errors.New("only registered participants can have result")
		}
		var result models.Result
		if len(resultData.ResultId) > 0 {
			exResult, _ := resultRepository.FindById(resultData.ResultId)
			if exResult != nil {
				result = *exResult
			}
		}
		result.Time = resultData.Time
		result.FinalPoint = resultData.FinalPoint
		result.Participant = mogo.RefField{ID: bson.ObjectIdHex(resultData.ParticipantId)}
		results = append(results, &result)
	}
	requireUpdate := false
	for _, result := range results {
		result, err := resultRepository.SaveResult(result)
		if err == nil {
			if !utils.MapContainsKey(utils.ConvertRefFieldSliceToStringMap(td.Results), result.ID.Hex()) {
				requireUpdate = true
				td.Results = append(td.Results, &mogo.RefField{ID: result.ID})
			}
		}
	}
	if requireUpdate {
		return trainingDataRepository.SaveTrainingData(td)
	}
	return td, nil
}
