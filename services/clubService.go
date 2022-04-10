package services

import (
	"errors"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type clubService struct{}

func (cb *clubService) CreateClub(clubDto *models.ClubDto) (*models.Club, error) {
	club := models.Club{}
	serviceManager := GetServiceManagerInstance()
	userService := serviceManager.GetUserService()
	rawUser, _ := userService.GetById(clubDto.CoachId, true)
	if rawUser == nil {
		return nil, errors.New("unable to validate coach")
	}
	user, _ := rawUser.(*models.UserResult)
	if &user.Club != nil {
		return nil, errors.New("coach is not available")
	}
	if user.UserType.Name != utils.COACH {
		return nil, errors.New("user must be a coach")
	}
	err := copier.Copy(&club, clubDto)
	if err != nil {
		return nil, err
	}
	club.Coach = mogo.RefField{ID: user.ID}
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	createdClub, err := clubRepository.SaveClub(&club)
	if err != nil {
		return nil, err
	}
	var userToUpdateSlice []*models.User
	rawUser, _ = userService.GetById(clubDto.CoachId, false)
	userToUpdate, _ := rawUser.(*models.User)
	userToUpdate.Club = mogo.RefField{ID: createdClub.ID}
	userToUpdateSlice = append(userToUpdateSlice, userToUpdate)
	_, _ = GetServiceManagerInstance().GetUserService().UpdateUsers(userToUpdateSlice)

	return createdClub, nil
}

func (cb *clubService) GetById(id string, fetchRelationShips bool) (interface{}, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindById(id, fetchRelationShips)
}

func (cb *clubService) AddMember(newMember *models.AddMember) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	rawClub, _ := clubRepository.FindById(newMember.ClubId, false)
	if rawClub == nil {
		return nil, errors.New("club not found")
	}
	club, _ := rawClub.(*models.Club)
	userService := GetServiceManagerInstance().GetUserService()
	var updateUserSlice []*models.User
	rawUser, _ := userService.GetById(newMember.MemberId, true)
	if rawUser == nil {
		return nil, errors.New("all members must be registered on the system")
	}
	user, _ := rawUser.(*models.UserResult)
	if user.UserType.Name != utils.SWIMMER {
		return nil, errors.New("all members must be swimmers")
	}
	if len(user.Club.ID.Hex()) > 0 {
		return nil, errors.New("user already belong to a club")
	}
	rawUser, _ = userService.GetById(newMember.MemberId, false)
	if userToUpdate, ok := rawUser.(*models.User); ok {
		userToUpdate.Club = mogo.RefField{ID: club.ID}
		updateUserSlice = append(updateUserSlice, userToUpdate)
	}
	club.Members = append(club.Members, &mogo.RefField{ID: user.ID})

	club, err := clubRepository.SaveClub(club)
	if err != nil {
		return nil, err
	}
	_, _ = userService.UpdateUsers(updateUserSlice)
	return club, nil
}

func (cb *clubService) GetAllClubs() ([]*models.ClubResult, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindAll()
}

func (cb *clubService) GetByName(name string) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindByName(name)
}
