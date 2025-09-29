package api

import (
	"go-api/internal/api/handler"
	"go-api/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMiddlewares(app *gin.Engine) {
	app.Use(middleware.Recover())
	app.Use(middleware.Cors())
	app.Use(middleware.RequestLog())
}

func RegisterRoutes(app *gin.Engine) {
	registerUnauthorizedRoutes(app)
	registerAuthorizedRoutes(app)
}

// 不需要鉴权的接口
func registerUnauthorizedRoutes(app *gin.Engine) {
	app.GET("/", handler.IndexHandler)
	app.POST("/login", handler.LoginHandler)
	app.POST("/register", handler.RegisterHandler)
	app.POST("/refreshToken", handler.RefreshTokenHandler)
}

// 需要鉴权的接口
func registerAuthorizedRoutes(app *gin.Engine) {
	authorized := app.Group("/")
	authorized.Use(middleware.Auth())
	authorized.GET("/me", handler.MeHandler)
}