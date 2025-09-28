package router

import (
	"go-api/app/controllers"
	"go-api/app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	indexController := &controllers.IndexController{}
	sessionController := &controllers.SessionController{}
	userController := &controllers.UserController{}

	authorized := app.Group("/")
	authorized.Use(middlewares.Auth())

	notAuthorized := app.Group("/")

	app.GET("/", indexController.Index)

	// 不需要登陆
	notAuthorized.POST("/login", sessionController.Login)

	// 需要登陆
	authorized.GET("/me", userController.Me)
}
