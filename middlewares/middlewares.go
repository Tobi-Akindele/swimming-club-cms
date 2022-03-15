package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/services"
	"swimming-club-cms-be/utils"
	"time"
)

func HasAuthority(permission string) gin.HandlerFunc {
	return Authentication(permission)
}

func Authentication(permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get(utils.AUTHORIZATION)
		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
				Code:    http.StatusBadRequest,
				Message: "authorization header is required",
			})
			return
		}
		authHeaderSlice := strings.Split(authHeader, " ")
		if len(authHeaderSlice) < 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
				Code:    http.StatusBadRequest,
				Message: "invalid token",
			})
			return
		}
		tokenString := strings.TrimSpace(authHeaderSlice[1])
		token, err := jwt.ParseWithClaims(
			tokenString,
			&dtos.SignedDetails{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(utils.GetEnv(utils.JWT_SECRET_KEY, "")), nil
			})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}
		if claims, ok := token.Claims.(*dtos.SignedDetails); ok && token.Valid {
			if claims.ExpiresAt < time.Now().Local().Unix() {
				ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
					Code:    http.StatusNotFound,
					Message: "token is expired",
				})
				return
			}
			userService := services.UserService{}
			user, err := userService.GetById(claims.UserId)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
					Code:    http.StatusNotFound,
					Message: "user not found",
				})
				return
			}
			permissionService := services.PermissionService{}
			permissions, _ := permissionService.GetRolesPermissions(utils.ConvertRefFieldSliceToStringSlice(user.Roles))
			if !utils.HasPermission(permissions, permission) {
				ctx.AbortWithStatusJSON(http.StatusForbidden, dtos.Response{
					Code:    http.StatusForbidden,
					Message: "access denied",
				})
				return
			}
			ctx.Set(utils.USER, user)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
				Code:    http.StatusBadRequest,
				Message: "invalid token",
			})
			return
		}
	}
}
