package controllers

import (
	"go-api/app/models"
	"go-api/app/utils"
	"net/http"

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

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *UserController) Create(ctx *gin.Context) {
	var createUserRequest CreateUserRequest
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		utils.ResponseError(ctx, gin.H{}, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	user, err := user.GetByUsername(createUserRequest.Username)
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "添加用户失败", 1)
		return
	}
	
	if user.ID > 0 {
		utils.ResponseError(ctx, gin.H{}, "用户已存在", 1)
		return
	}

	user.Username = createUserRequest.Username
	user.Password = createUserRequest.Password
	err = user.Create()
	if err != nil {
		utils.ResponseError(ctx, gin.H{}, "添加用户失败", 1)
		return
	}
	utils.ResponseSuccess(ctx, gin.H{}, "添加用户成功")
}