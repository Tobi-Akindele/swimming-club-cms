package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/utils"
)

type PingController struct{}

func (pc *PingController) PingAPI(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dtos.Ping{
		Message: utils.GetEnv(utils.APP_NAME, "SWIMMING-CLUB-CMS") + " UP!",
	})
}
