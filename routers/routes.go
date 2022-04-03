package routers

import (
	"github.com/gin-gonic/gin"
	"swimming-club-cms-be/controllers"
	"swimming-club-cms-be/middlewares"
)

func setPingRoutes(router *gin.Engine) {
	pingController := new(controllers.PingController)
	router.GET("/ping", pingController.PingAPI)
}

func setUnAuthenticatedRoutes(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/signup", authController.SignUp)
	router.POST("/login", authController.Login)
	router.POST("/set/password", authController.SetPassword)
}

func setRoleRoutes(router *gin.Engine) {
	roleController := new(controllers.RoleController)

	createRoleRoute := router.Group("/role", middlewares.HasAuthority("CREATE_ROLE"))
	createRoleRoute.POST("", roleController.CreateRole)

	getAllRoleRoute := router.Group("/roles", middlewares.HasAuthority("GET_ALL_ROLES"))
	getAllRoleRoute.GET("", roleController.GetAllRoles)

	getRoleByNameRoute := router.Group("/role/name", middlewares.HasAuthority("GET_ROLE_BY_NAME"))
	getRoleByNameRoute.GET("", roleController.GetRoleByName)

	getRoleByIdRoute := router.Group("/role/:id", middlewares.HasAuthority("GET_ROLE_BY_ID"))
	getRoleByIdRoute.GET("", roleController.GetRoleById)

	getRolePermissionsRoute := router.Group("/role/:id/permissions", middlewares.HasAuthority("GET_ROLE_PERMISSIONS"))
	getRolePermissionsRoute.GET("", roleController.GetRolePermissions)

	assignPermissionsToRoleRoute := router.Group("/assign/permissions", middlewares.HasAuthority("ASSIGN_PERMISSIONS_TO_ROLE"))
	assignPermissionsToRoleRoute.POST("", roleController.AssignPermissionsToRole)

	removePermissionsFromRoleRoute := router.Group("/remove/permissions", middlewares.HasAuthority("REMOVE_ROLE_PERMISSIONS"))
	removePermissionsFromRoleRoute.POST("", roleController.RemovePermissionsFromRole)
}

func setPermissionRoutes(router *gin.Engine) {
	permissionController := new(controllers.PermissionController)

	getAllPermissionRoute := router.Group("/permissions", middlewares.HasAuthority("GET_ALL_PERMISSIONS"))
	getAllPermissionRoute.GET("", permissionController.GetAllPermissions)
}

func setUserTypeRoutes(router *gin.Engine) {
	userTypeController := new(controllers.UserTypeController)

	createUserTypeRoute := router.Group("/usertype", middlewares.HasAuthority("CREATE_USER_TYPE"))
	createUserTypeRoute.POST("", userTypeController.CreateUserType)

	getAllUserTypeRoute := router.Group("/user-types", middlewares.HasAuthority("GET_ALL_USER_TYPES"))
	getAllUserTypeRoute.GET("", userTypeController.GetAllUserTypes)
}

func setClubRoutes(router *gin.Engine) {
	clubController := new(controllers.ClubController)

	createClubRoute := router.Group("/club", middlewares.HasAuthority("CREATE_CLUB"))
	createClubRoute.POST("", clubController.CreateClub)

	addMemberToClubRoute := router.Group("/club/add/member", middlewares.HasAuthority("ADD_MEMBER_TO_CLUB"))
	addMemberToClubRoute.POST("", clubController.AddMembers)

	getClubByIdRoute := router.Group("/club/:id", middlewares.HasAuthority("GET_CLUB_BY_ID"))
	getClubByIdRoute.GET("", clubController.GetClubById)

	getAllClubsRoute := router.Group("/clubs", middlewares.HasAuthority("GET_ALL_CLUBS"))
	getAllClubsRoute.GET("", clubController.GetAllClubs)
}

func setCompetitionRoutes(router *gin.Engine) {
	competitionController := new(controllers.CompetitionController)

	createCompetitionRoute := router.Group("/competition", middlewares.HasAuthority("CREATE_COMPETITION"))
	createCompetitionRoute.POST("", competitionController.CreateCompetition)

	getCompetitionByIdRoute := router.Group("/competition/:id", middlewares.HasAuthority("GET_COMPETITION_BY_ID"))
	getCompetitionByIdRoute.GET("", competitionController.GetCompetitionById)
}

func setEventRoutes(router *gin.Engine) {
	eventController := new(controllers.EventController)

	createEventRoute := router.Group("/event", middlewares.HasAuthority("CREATE_EVENT"))
	createEventRoute.POST("", eventController.CreateEvent)

	getEventByIdRoute := router.Group("/event/:id", middlewares.HasAuthority("GET_EVENT_BY_ID"))
	getEventByIdRoute.GET("", eventController.GetEventById)

	addParticipantsToClubRoute := router.Group("/event/add/participants", middlewares.HasAuthority("ADD_PARTICIPANTS_TO_EVENT"))
	addParticipantsToClubRoute.POST("", eventController.AddParticipants)
}

func setUserRoutes(router *gin.Engine) {
	userController := new(controllers.UserController)

	getAllUsersRoute := router.Group("/users", middlewares.HasAuthority("GET_ALL_USERS"))
	getAllUsersRoute.GET("", userController.GetAllUsers)

	getByUsernameRoute := router.Group("/user/username", middlewares.HasAuthority("GET_USER_BY_USERNAME"))
	getByUsernameRoute.GET("", userController.GetByUsername)

	getByEmailRoute := router.Group("/user/email", middlewares.HasAuthority("GET_USER_BY_EMAIL"))
	getByEmailRoute.GET("", userController.GetByEmail)
}

func HandleRequests() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	// Initialize routes
	setPingRoutes(router)
	setUnAuthenticatedRoutes(router)
	setRoleRoutes(router)
	setPermissionRoutes(router)
	setUserTypeRoutes(router)
	setClubRoutes(router)
	setCompetitionRoutes(router)
	setEventRoutes(router)
	setUserRoutes(router)

	return router
}
