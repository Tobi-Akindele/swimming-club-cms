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
	rawUser, _ := serviceManager.GetUserService().GetById(clubDto.CoachId, true)
	if rawUser == nil {
		return nil, errors.New("unable to validate coach")
	}
	user, _ := rawUser.(*models.UserResult)
	if user.UserType.Name != utils.COACH {
		return nil, errors.New("user must be a coach")
	}
	err := copier.Copy(&club, clubDto)
	if err != nil {
		return nil, err
	}
	club.Coach = mogo.RefField{ID: user.ID}
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.SaveClub(&club)
}

func (cb *clubService) GetById(id string) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindById(id)
}

func (cb *clubService) AddMembers(newMembers *models.AddMember) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	club, err := clubRepository.FindById(newMembers.ClubId)
	if err != nil {
		return nil, err
	} else if club == nil {
		return nil, errors.New("club not found")
	}
	userService := GetServiceManagerInstance().GetUserService()
	for idx := range newMembers.NewMembers {
		rawUser, _ := userService.GetById(newMembers.NewMembers[idx], true)
		if rawUser == nil {
			return nil, errors.New("all members must be registered on the system")
		}
		user, _ := rawUser.(*models.UserResult)
		if user.UserType.Name != utils.SWIMMER {
			return nil, errors.New("all members must be swimmers")
		}
		userHasClub, _ := clubRepository.FindByMemberId(user.ID.Hex())
		if userHasClub != nil {
			return nil, errors.New("user already belong to a club")
		}
		club.Members = append(club.Members, &mogo.RefField{ID: user.ID})
	}
	return clubRepository.SaveClub(club)
}

func (cb *clubService) GetAllClubs() ([]*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindAll()
}
