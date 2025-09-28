package controllers

import (
	"go-api/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
}

func (c *SessionController) Login(ctx *gin.Context) {
	token, err := utils.GenerateToken(13800);
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "Login failed", http.StatusInternalServerError)
		return
	}
	utils.ResponseSuccess(ctx, gin.H{
		"token": token,
	}, "success")
}