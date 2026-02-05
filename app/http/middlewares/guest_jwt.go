package middlewares

import (
	"errors"
	"net/http"

	"github.com/Zero0719/go-api/pkg/jwt"
	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Error(c, errors.New("请使用游客身份访问"), http.StatusForbidden)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}