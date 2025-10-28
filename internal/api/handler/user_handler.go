package handler

import (
	"fmt"
	"go-api/internal/db"
	"go-api/internal/repository"
	"go-api/internal/request"
	"go-api/internal/service"
	"go-api/pkg/jwt"
	"go-api/pkg/response"

	"github.com/gin-gonic/gin"
)



func LoginHandler(ctx *gin.Context) {
	request, err := BindAndValidate[request.LoginRequest](ctx)
	if err != nil {
		
	}
	userService := service.NewUserService(repository.NewUserRepository(db.DB))
	user, err := userService.Login(ctx, request.Username, request.Password)
	if err != nil {
		response.Error(ctx, gin.H{}, err.Error(), 1)
		return
	}

	token, refreshToken, err := jwt.GenerateToken(int(user.ID))
	if err != nil {
		response.Error(ctx, gin.H{}, err.Error(), 1)
		return
	}

	response.Success(ctx, gin.H{"token": token, "refreshToken": refreshToken}, "success")
}

func RegisterHandler(ctx *gin.Context) {
	request, err := BindAndValidate[request.RegisterRequest](ctx)

	if err != nil {
		return
	}

	userService := service.NewUserService(repository.NewUserRepository(db.DB))
	user, err := userService.Register(ctx, request.Username, request.Password)
	if err != nil {
		response.Error(ctx, gin.H{}, err.Error(), 1)
		return
	}

	token, refreshToken, err := jwt.GenerateToken(int(user.ID))
	if err != nil {
		response.Error(ctx, gin.H{}, err.Error(), 1)
		return
	}

	response.Success(ctx, gin.H{"token": token, "refreshToken": refreshToken}, "success")
}

type MeResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func MeHandler(ctx *gin.Context) {
	userService := service.NewUserService(repository.NewUserRepository(db.DB))
	user, err := userService.GetByID(ctx, uint(ctx.GetInt("currentUserId")))
	if err != nil {
		response.Error(ctx, gin.H{}, err.Error(), 1)
		return
	}
	response.Success(ctx, MeResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, "success")
}

func RefreshTokenHandler(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Refresh-Token")

	fmt.Println(refreshToken)
	if refreshToken == "" {
		response.Error(ctx, gin.H{}, "Refresh-Token is required", 1)
		return
	}

	token, refreshToken, err := jwt.RefreshToken(refreshToken)
	if err != nil {
		response.Error(ctx, gin.H{}, err.Error(), 1)
		return
	}
	response.Success(ctx, gin.H{"token": token, "refreshToken": refreshToken}, "success")
}
