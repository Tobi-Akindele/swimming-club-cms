package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"swimming-club-cms-be/dtos"
	"swimming-club-cms-be/models"
	"swimming-club-cms-be/services"
	"swimming-club-cms-be/utils"
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
		if !strings.HasPrefix(authHeader, utils.BEARER) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
				Code:    http.StatusBadRequest,
				Message: "bearer token is required",
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
			serviceManager := services.GetServiceManagerInstance()
			rawUser, _ := serviceManager.GetUserService().GetById(claims.UserId, false)
			if rawUser == nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, dtos.Response{
					Code:    http.StatusNotFound,
					Message: "User not found",
				})
				return
			}
			rawRole, _ := serviceManager.GetRoleService().GetById(claims.RoleId, false)
			if rawRole == nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.Response{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized access",
				})
				return
			}
			permissionObj, _ := serviceManager.GetPermissionService().GetByValue(permission)
			role, _ := rawRole.(*models.Role)
			if !utils.MapContainsKey1(utils.ConvertRefFieldSliceToMap(role.Permissions), permissionObj.ID.Hex()) {
				ctx.AbortWithStatusJSON(http.StatusForbidden, dtos.Response{
					Code:    http.StatusForbidden,
					Message: "Forbidden access",
				})
				return
			}
			if user, parsed := rawUser.(*models.User); !parsed {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.Response{
					Code:    http.StatusBadRequest,
					Message: "An occurred parsing user data",
				})
				return
			} else {
				ctx.Set(utils.USER, user)
			}
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, *")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
