package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send(ctx *gin.Context, data interface{}, message string, code int) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"msg": message,
		"code":    code,
	})
}

func Success(ctx *gin.Context, data interface{}, message string) {
	Send(ctx, data, message, 0)
}

func Error(ctx *gin.Context, data interface{}, message string, code int) {
	Send(ctx, data, message, code)
}
