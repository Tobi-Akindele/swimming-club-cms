package services

import "sync"

var lock = &sync.Mutex{}

type serviceManager struct {
	authService        *authService
	clubService        *clubService
	competitionService *competitionService
	eventService       *eventService
	permissionService  *permissionService
	roleService        *roleService
	userService        *userService
	userTypeService    *userTypeService
}

var instance *serviceManager

func GetServiceManagerInstance() *serviceManager {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &serviceManager{}
		}
	}
	return instance
}

func (sm *serviceManager) GetAuthService() *authService {
	if sm.authService == nil {
		sm.authService = &authService{}
	}
	return sm.authService
}

func (sm *serviceManager) GetClubService() *clubService {
	if sm.clubService == nil {
		sm.clubService = &clubService{}
	}
	return sm.clubService
}

func (sm *serviceManager) GetCompetitionService() *competitionService {
	if sm.competitionService == nil {
		sm.competitionService = &competitionService{}
	}
	return sm.competitionService
}

func (sm *serviceManager) GetEventService() *eventService {
	if sm.eventService == nil {
		sm.eventService = &eventService{}
	}
	return sm.eventService
}

func (sm *serviceManager) GetPermissionService() *permissionService {
	if sm.permissionService == nil {
		sm.permissionService = &permissionService{}
	}
	return sm.permissionService
}

func (sm *serviceManager) GetRoleService() *roleService {
	if sm.roleService == nil {
		sm.roleService = &roleService{}
	}
	return sm.roleService
}

func (sm *serviceManager) GetUserService() *userService {
	if sm.userService == nil {
		sm.userService = &userService{}
	}
	return sm.userService
}

func (sm *serviceManager) GetUserTypeService() *userTypeService {
	if sm.userTypeService == nil {
		sm.userTypeService = &userTypeService{}
	}
	return sm.userTypeService
}
