package middlewares

import (
	"fmt"

	"github.com/Zero0719/go-api/app/models/user"
	"github.com/Zero0719/go-api/pkg/jwt"
	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)

		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请先登录: %v", err.Error()))
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，请重新登录")
			return
		}

		c.Set("current_user_id", userModel.ID)
		c.Set("current_user", userModel)
		c.Next()
	}
}