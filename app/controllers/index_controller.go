package controllers

import (
	"go-api/app/utils"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (c *IndexController) Index(ctx *gin.Context) {
	utils.Logger.Info().Msg("IndexController Index")
	utils.ResponseSuccess(ctx, gin.H{
		"message":   "Hello World",
		"timestamp": "2025-09-26",
	}, "success")
}
