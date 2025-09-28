package middlewares

import (
	"go-api/app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		token := strings.Split(authorization, " ")[1]
		if token == "" {
			utils.ResponseError(ctx, gin.H{}, "Unauthorized", http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		userId, err := utils.ParseToken(token)
		if err != nil {
			utils.ResponseError(ctx, gin.H{}, "Unauthorized", http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}