package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
)

type RoleController struct{}

func (rc *RoleController) CreateRole(ctx *gin.Context) {
	var role models.RoleDto
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(role); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	serviceManager := services.GetServiceManagerInstance()
	createdRole, err := serviceManager.GetRoleService().CreateRole(&role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdRole)
	}
}

func (rc *RoleController) GetAllRoles(ctx *gin.Context) {
	serviceManager := services.GetServiceManagerInstance()
	roles, err := serviceManager.GetRoleService().GetAllRoles()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, roles)
	}
}

func (rc *RoleController) GetRoleByName(ctx *gin.Context) {
	name := ctx.GetHeader("name")
	serviceManager := services.GetServiceManagerInstance()
	role, err := serviceManager.GetRoleService().GetByName(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, role)
	}
}

func (rc *RoleController) GetRoleById(ctx *gin.Context) {
	roleId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	role, err := serviceManager.GetRoleService().GetById(roleId, true)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, role)
	}
}

func (rc *RoleController) AssignPermissionsToRole(ctx *gin.Context) {
	var permissionIds models.AssignPermissions
	if err := ctx.ShouldBindJSON(&permissionIds); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	if errs := validator.Validate(permissionIds); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	serviceManager := services.GetServiceManagerInstance()
	role, err := serviceManager.GetRoleService().AssignPermissionsToRole(&permissionIds)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, role)
	}
}

func (rc *RoleController) RemovePermissionsFromRole(ctx *gin.Context) {
	var removePermissions models.RemovePermissions
	if err := ctx.ShouldBindJSON(&removePermissions); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	log.Println(removePermissions)
	if errs := validator.Validate(removePermissions); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: errs.Error()})
		return
	}

	serviceManager := services.GetServiceManagerInstance()
	role, err := serviceManager.GetRoleService().RemovePermissionsFromRole(&removePermissions)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, role)
	}
}

func (rc *RoleController) GetRolePermissions(ctx *gin.Context) {
	roleId := ctx.Param("id")
	serviceManager := services.GetServiceManagerInstance()
	rolePermissions, err := serviceManager.GetRoleService().GetRolePermissions(roleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{Code: http.StatusBadRequest, Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, rolePermissions)
	}
}
