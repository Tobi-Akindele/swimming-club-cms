package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-club-cms-be/utils"
)

type Ping struct {
	Message string `json:"message"`
}

func (ping *Ping) PingAPI(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Ping{
		Message: utils.GetEnv(utils.APP_NAME, "SWIMMING-CLUB-CMS") + " UP!",
	})
}
