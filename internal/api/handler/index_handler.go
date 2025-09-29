package handler

import (
	"go-api/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"message": "Hello go-api", 
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}, "success")
}