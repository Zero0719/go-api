package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, data interface{}, message string, code int) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": message,
		"code":    code,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}, message string) {
	Response(ctx, data, message, 0)
}

func ResponseError(ctx *gin.Context, data interface{}, message string, code int) {
	Response(ctx, data, message, code)
}
