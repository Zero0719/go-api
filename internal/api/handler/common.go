package handler

import (
	"go-api/pkg/response"
	"go-api/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindAndValidate 通用的参数绑定和验证函数
func BindAndValidate[T any](ctx *gin.Context) (*T, error) {
	var request T

	// 参数绑定
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Error(ctx, gin.H{}, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	// 参数验证
	if err := validator.ValidateStructFirstError(request); err != nil {
		response.Error(ctx, gin.H{}, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &request, nil
}
