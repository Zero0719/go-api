package middlewares

import (
	"go-api/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.ResponseError(ctx, []interface{}{}, "internal server error", http.StatusInternalServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
