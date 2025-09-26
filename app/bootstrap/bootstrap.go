package bootstrap

import (
	"go-api/app/middlewares"
	"go-api/app/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	app := gin.New()
	registerGlobalMiddlewares(app)
	router.RegisterRoutes(app)
	app.Run()
}

func registerGlobalMiddlewares(app *gin.Engine)  {
	app.Use(middlewares.Recover())
	app.Use(cors.New(cors.Config{
		// 允许的源，*表示允许所有源
		AllowOrigins:     []string{"*"},
		// 允许的请求方法
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// 允许的请求头
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		// 是否允许带cookie
		AllowCredentials: true,
		// 预检请求的缓存时间
		MaxAge:           12 * time.Hour,
	}))
}
