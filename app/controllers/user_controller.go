package controllers

import (
	"go-api/app/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) Me(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	utils.ResponseSuccess(ctx, gin.H{
		"userId": userId,
	}, "success")
}