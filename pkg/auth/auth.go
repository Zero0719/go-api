package auth

import (
	"errors"
	"github.com/Zero0719/go-api/app/models/user"

	"github.com/Zero0719/go-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("获取当前用户失败"))
		return user.User{}
	}
	return userModel
}

func CurrentUserID(c *gin.Context) string {
	return c.GetString("current_user_id")
}