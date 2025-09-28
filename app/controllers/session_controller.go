package controllers

import (
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
	token, err := utils.GenerateToken(13800);
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "Login failed", http.StatusInternalServerError)
		return
	}
	utils.ResponseSuccess(ctx, gin.H{
		"token": token,
	}, "success")
}