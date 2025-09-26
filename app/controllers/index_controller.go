package controllers

import (
	"go-api/app/utils"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	
}

func (c *IndexController) Index(ctx *gin.Context) {
	utils.ResponseSuccess(ctx, []interface{}{}, "success")
}