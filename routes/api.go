package routes

import (
	"net/http"

	"github.com/Zero0719/go-api/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
    v1 := router.Group("/v1")
    {
        authGroup := v1.Group("/auth")
        {
            suc := new(auth.SignupController)
            authGroup.POST("/signup/email/exist", suc.IsEmailExist)
        }


        v1.GET("/", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "Hello": "World!",
            })
        })
    }
}