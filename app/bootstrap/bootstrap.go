package bootstrap

import (
	"context"
	"go-api/app/middlewares"
	"go-api/app/models"
	"go-api/app/router"
	"go-api/app/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	// 初始化日志
	utils.InitLogger()

	// 初始化配置
	utils.InitConfig()

	// 初始化数据库
	models.InitDB()

	// 初始化redis
	utils.InitRedis()

	app := gin.New()
	registerGlobalMiddlewares(app)
	router.RegisterRoutes(app)
	utils.Logger.Info().Msg("Starting server...")


	srv := &http.Server{
		Addr:    ":" + utils.Config.GetString("app.port"),
		Handler: app.Handler(),
	}


	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Error().Msgf("Server error: %v", err)
		}
	}()

	utils.Logger.Info().Msg("Server started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Logger.Info().Msg("Server shutting down")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger.Error().Msgf("Server shutdown error: %v", err)
	}
	utils.Logger.Info().Msg("Server shutdown")
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
