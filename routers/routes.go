package routers

import (
	"github.com/gin-gonic/gin"
	"swimming-club-cms-be/controllers"
)

func setPingRoute(router *gin.Engine) {
	pingController := new(controllers.Ping)
	router.GET("/ping", pingController.PingAPI)
}

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/signup", authController.SignUp)
}

func HandleRequests() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Initialize routes
	setPingRoute(router)
	setAuthRoute(router)

	return router
}
