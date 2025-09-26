package bootstrap

import (
	"go-api/app/router"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	app := gin.New()
	router.RegisterRoutes(app)
	app.Run()
}
