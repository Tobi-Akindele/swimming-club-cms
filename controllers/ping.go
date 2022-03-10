package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-club-cms-be/utils"
)

type Ping struct {
	Status string `json:"status"`
}

func (ping *Ping) PingAPI(ctx *gin.Context) {

	p := Ping{
		Status: utils.GetEnv(utils.APP_NAME, "SWIMMING-CLUB-CMS") + " is running",
	}
	ctx.JSON(http.StatusOK, p)
}
