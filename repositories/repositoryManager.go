package repositories

import "sync"

var lock = &sync.Mutex{}

type repositoryManager struct {
	clubRepository        *clubRepository
	competitionRepository *competitionRepository
	eventRepository       *eventRepository
	permissionRepository  *permissionRepository
	roleRepository        *roleRepository
	userRepository        *userRepository
	userTypeRepository    *userTypeRepository
}

var instance *repositoryManager

func GetRepositoryManagerInstance() *repositoryManager {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &repositoryManager{}
		}
	}
	return instance
}

func (rm *repositoryManager) GetClubRepository() *clubRepository {
	if rm.clubRepository == nil {
		rm.clubRepository = &clubRepository{}
	}
	return rm.clubRepository
}

func (rm *repositoryManager) GetCompetitionRepository() *competitionRepository {
	if rm.competitionRepository == nil {
		rm.competitionRepository = &competitionRepository{}
	}
	return rm.competitionRepository
}

func (rm *repositoryManager) GetEventRepository() *eventRepository {
	if rm.eventRepository == nil {
		rm.eventRepository = &eventRepository{}
	}
	return rm.eventRepository
}

func (rm *repositoryManager) GetPermissionRepository() *permissionRepository {
	if rm.permissionRepository == nil {
		rm.permissionRepository = &permissionRepository{}
	}
	return rm.permissionRepository
}

func (rm *repositoryManager) GetRoleRepository() *roleRepository {
	if rm.roleRepository == nil {
		rm.roleRepository = &roleRepository{}
	}
	return rm.roleRepository
}

func (rm *repositoryManager) GetUserRepository() *userRepository {
	if rm.userRepository == nil {
		rm.userRepository = &userRepository{}
	}
	return rm.userRepository
}

func (rm *repositoryManager) GetUserTypeRepository() *userTypeRepository {
	if rm.userTypeRepository == nil {
		rm.userTypeRepository = &userTypeRepository{}
	}
	return rm.userTypeRepository
}
