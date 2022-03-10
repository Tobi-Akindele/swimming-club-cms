package routers

import (
	"github.com/gin-gonic/gin"
	"swimming-club-cms-be/controllers"
)

func setPingRoute(router *gin.Engine) {
	pingController := new(controllers.Ping)
	router.GET("/ping", pingController.PingAPI)
}

func HandleRequests() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Initialize routes
	setPingRoute(router)

	return router
}
