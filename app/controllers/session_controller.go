package controllers

import (
	"go-api/app/models"
	"go-api/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *SessionController) Login(ctx *gin.Context) {
	var loginRequest LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.ResponseError(ctx, gin.H{}, "Invalid request", http.StatusBadRequest)
		return
	}

	// todo 获取数据库数据进行校验等

	var user models.User
	user, err := user.GetByUsername(loginRequest.Username)
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "Login failed", http.StatusInternalServerError)
		return
	}
	if (user.ID == 0) {
		utils.ResponseError(ctx, gin.H{}, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	if user.Password != utils.MD5(loginRequest.Password + utils.Config.GetString("app.salt")) {
		utils.ResponseError(ctx, gin.H{}, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	token, refreshToken, err := utils.GenerateToken(int(user.ID));
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "Login failed", http.StatusInternalServerError)
		return
	}
	utils.ResponseSuccess(ctx, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	}, "success")
}

func (c *SessionController) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Refresh-Token")
	if refreshToken == "" {
		utils.ResponseError(ctx, gin.H{}, "Refresh token is required", http.StatusBadRequest)
		return
	}
	token, newRefreshToken, err := utils.RefreshToken(refreshToken)
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "Refresh token is invalid", http.StatusUnauthorized)
		return
	}

	utils.ResponseSuccess(ctx, gin.H{
		"token":        token,
		"refreshToken": newRefreshToken,
	}, "success")
}