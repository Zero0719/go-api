package router

import (
	"go-api/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	controller := &controllers.IndexController{}

	app.GET("/", controller.Index)
}
