package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goonode/mogo"
	"github.com/jinzhu/copier"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/repositories"
	"swimming-club-cms-be/utils"
)

type clubService struct{}

func (cb *clubService) CreateClub(clubDto *models.ClubDto) (*models.Club, error) {
	club := models.Club{}
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	serviceManager := GetServiceManagerInstance()
	userService := serviceManager.GetUserService()
	clubByCoachId, _ := clubRepository.FindByCoachId(clubDto.CoachId)
	if clubByCoachId != nil {
		return nil, errors.New("coach is not available")
	}
	rawUser, _ := userService.GetById(clubDto.CoachId, true)
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
	return clubRepository.SaveClub(&club)
}

func (cb *clubService) GetById(id string, fetchRelationShips bool) (interface{}, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindById(id, fetchRelationShips)
}

func (cb *clubService) AddMember(newMember *models.AddMember, ctx *gin.Context) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	rawClub, _ := clubRepository.FindById(newMember.ClubId, false)
	if rawClub == nil {
		return nil, errors.New("club not found")
	}
	club, _ := rawClub.(*models.Club)
	if requester, isAdmin := utils.IsAdmin(ctx); !isAdmin {
		if requester.ID.Hex() != club.ID.Hex() {
			return nil, errors.New("action not allowed")
		}
	}
	clubByMemberId, _ := clubRepository.FindByMemberId(newMember.MemberId)
	if clubByMemberId != nil {
		return nil, errors.New("user already belong to a club")
	}
	userService := GetServiceManagerInstance().GetUserService()
	rawUser, _ := userService.GetById(newMember.MemberId, true)
	if rawUser == nil {
		return nil, errors.New("all members must be registered on the system")
	}
	user, _ := rawUser.(*models.UserResult)
	if user.UserType.Name != utils.SWIMMER {
		return nil, errors.New("all members must be swimmers")
	}
	club.Members = append(club.Members, &mogo.RefField{ID: user.ID})

	return clubRepository.SaveClub(club)
}

func (cb *clubService) GetAllClubs() ([]*models.ClubResult, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindAll()
}

func (cb *clubService) GetByName(name string) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindByName(name)
}

func (cb *clubService) RemoveMembers(removeMembers *models.RemoveMembers, ctx *gin.Context) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	rawClub, _ := clubRepository.FindById(removeMembers.ClubId, false)
	if rawClub == nil {
		return nil, errors.New("club not found")
	}

	club, _ := rawClub.(*models.Club)
	if requester, isAdmin := utils.IsAdmin(ctx); !isAdmin {
		if requester.ID.Hex() != club.ID.Hex() {
			return nil, errors.New("action not allowed")
		}
	}
	club.Members = utils.RemoveRefFromRefSlice(club.Members, removeMembers.MemberIds)

	return clubRepository.SaveClub(club)
}

func (cb *clubService) UpdateClub(clubUpdate *models.ClubUpdate, ctx *gin.Context) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	rawClub, _ := clubRepository.FindById(clubUpdate.ClubId, false)
	if rawClub == nil {
		return nil, errors.New("club not found")
	}
	club, _ := rawClub.(*models.Club)
	if requester, isAdmin := utils.IsAdmin(ctx); !isAdmin {
		if requester.ID.Hex() != club.ID.Hex() {
			return nil, errors.New("action not allowed")
		}
	}
	userService := GetServiceManagerInstance().GetUserService()
	rawUser, _ := userService.GetById(clubUpdate.CoachId, false)
	if rawUser == nil {
		return nil, errors.New("user not found")
	}
	user, _ := rawUser.(*models.User)
	club.Name = clubUpdate.Name
	club.Coach = mogo.RefField{ID: user.ID}

	return clubRepository.SaveClub(club)
}

func (cb *clubService) GetByMemberId(memberId string) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindByMemberId(memberId)
}

func (cb *clubService) GetByCoachId(coachId string) (*models.Club, error) {
	clubRepository := repositories.GetRepositoryManagerInstance().GetClubRepository()
	return clubRepository.FindByCoachId(coachId)
}
