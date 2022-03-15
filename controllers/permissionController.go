package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/services"
)

type PermissionController struct{}

func (pc *PermissionController) GetAllPermissions(ctx *gin.Context) {
	permissionService := services.PermissionService{}
	permissions, err := permissionService.GetAllPermissions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, permissions)
	}
}
