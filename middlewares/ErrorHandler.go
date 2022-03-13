package middlewares

import "github.com/gin-gonic/gin"

func ErrorHandler(statusCode int, ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		ctx.JSON(statusCode, gin.H{
			"errors": ctx.Errors,
		})
	}
}
