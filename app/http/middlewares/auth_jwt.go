package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Zero0719/go-api/app/models/user"
	"github.com/Zero0719/go-api/pkg/jwt"
	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)

		if err != nil {
			response.Error(c, errors.New(fmt.Sprintf("请先登录: %v", err.Error())), http.StatusUnauthorized)
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Error(c, errors.New("找不到对应用户，请重新登录"), http.StatusUnauthorized)
			return
		}

		c.Set("current_user_id", userModel.ID)
		c.Set("current_user", userModel)
		c.Next()
	}
}