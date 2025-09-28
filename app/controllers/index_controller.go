package controllers

import (
	"go-api/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (c *IndexController) Index(ctx *gin.Context) {
	utils.ResponseSuccess(ctx, gin.H{
		"message": "Hello go-api",
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}, "success")
}
