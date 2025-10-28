package handler

import (
	"go-api/pkg/response"
	"go-api/pkg/validator"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BindAndValidate 通用的参数绑定和验证函数
func BindAndValidate[T any](ctx *gin.Context) (*T, error) {
	var request T

	// 参数绑定
	if err := ctx.ShouldBind(&request); err != nil {
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

// Map 通用切片转换函数，将 []T 转换为 []R
func Map[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// FormatTime 格式化时间为字符串，统一的格式
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
