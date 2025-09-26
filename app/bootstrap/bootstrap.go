package bootstrap

import (
	"go-api/app/middlewares"
	"go-api/app/router"
	"go-api/app/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	// 初始化日志
	utils.InitLogger()

	app := gin.New()
	registerGlobalMiddlewares(app)
	router.RegisterRoutes(app)
	utils.Logger.Info().Msg("Starting server...")
	app.Run()
}

func registerGlobalMiddlewares(app *gin.Engine) {
	app.Use(middlewares.Recover())
	app.Use(middlewares.RequestLog()) // 添加请求日志中间件
	app.Use(cors.New(cors.Config{
		// 允许的源，*表示允许所有源
		AllowOrigins: []string{"*"},
		// 允许的请求方法
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// 允许的请求头
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		// 是否允许带cookie
		AllowCredentials: true,
		// 预检请求的缓存时间
		MaxAge: 12 * time.Hour,
	}))
}
