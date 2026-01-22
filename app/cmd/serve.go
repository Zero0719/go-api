package cmd

import (
	"github.com/Zero0719/go-api/bootstrap"
	"github.com/Zero0719/go-api/pkg/config"
	"github.com/Zero0719/go-api/pkg/console"
	"github.com/Zero0719/go-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use: "serve",
	Short: "Start web server",
	Run: runWeb,
	Args: cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// 初始化路由绑定
    bootstrap.SetupRoute(router)

    // 运行服务器
    err := router.Run(":" + config.Get[string]("app.port"))
    if err != nil {
        logger.ErrorString("CMD", "serve", err.Error())
        console.Exit("Unable to start server, error:" + err.Error())
    }
}