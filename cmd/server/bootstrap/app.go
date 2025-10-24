package bootstrap

import (
	"context"
	"go-api/internal/api"
	"go-api/internal/config"
	"go-api/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	initConfig()
	initLogger()
	initDB()
	initRedis()
	startServer()
}

func startServer() {
	app := gin.New()
	// 注册中间件	
	appConfig := config.Get().App
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(appConfig.Port),
		Handler: app.Handler(),
	}
	logger.CommonLogger.Info().Msg("Starting server...")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.CommonLogger.Error().Msgf("Server error: %v", err)
		}
	}()


	api.RegisterMiddlewares(app)
	api.RegisterRoutes(app)

	logger.CommonLogger.Info().Msg("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.CommonLogger.Info().Msg("Server shutting down")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.CommonLogger.Error().Msgf("Server shutdown error: %v", err)
	}
	logger.CommonLogger.Info().Msg("Server shutdown")
}