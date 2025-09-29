package middleware

import (
	"fmt"
	"go-api/pkg/logger"
	"go-api/pkg/response"
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

				logger.RequestLogger.Error().Msgf("Recover: %s, Stack: %s", recoverMsg, stack)
				response.Error(ctx, []interface{}{}, "internal server error", http.StatusInternalServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
