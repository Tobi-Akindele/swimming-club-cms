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

func setAuthRoutes(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/signup", authController.SignUp)
	router.POST("/login", authController.Login)
}

func setRoleRoutes(router *gin.Engine) {
	roleController := new(controllers.RoleController)

	createRoleRoute := router.Group("/role", middlewares.HasAuthority("CREATE_ROLE"))
	createRoleRoute.POST("/", roleController.CreateRole)
}

func setPermissionRoutes(router *gin.Engine) {
	permissionController := new(controllers.PermissionController)

	getAllPermissionRoute := router.Group("/permissions", middlewares.HasAuthority("GET_ALL_PERMISSIONS"))
	getAllPermissionRoute.GET("/", permissionController.GetAllPermissions)
}

func HandleRequests() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Initialize routes
	setPingRoutes(router)
	setAuthRoutes(router)
	setRoleRoutes(router)
	setPermissionRoutes(router)

	return router
}
