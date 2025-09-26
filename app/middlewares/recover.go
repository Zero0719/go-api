package middlewares

import (
	"fmt"
	"go-api/app/utils"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				recoverMsg := fmt.Sprintf("%v", r)

				utils.Logger.Error().Msgf("Recover: %s, Stack: %s", recoverMsg, stack)
				utils.ResponseError(ctx, []interface{}{}, "internal server error", http.StatusInternalServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
