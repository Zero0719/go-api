package middleware

import (
	"go-api/pkg/jwt"
	"go-api/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		token := strings.Split(authorization, " ")[1]
		if token == "" {
			response.Error(ctx, gin.H{}, "Unauthorized", http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		userId, err := jwt.ParseToken(token)
		if err != nil {
			response.Error(ctx, gin.H{}, "Unauthorized", http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		ctx.Set("currentUserId", userId)
		ctx.Next()
	}
}