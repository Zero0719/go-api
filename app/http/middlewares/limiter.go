package middlewares

import (
	"net/http"

	"github.com/Zero0719/go-api/pkg/app"
	"github.com/Zero0719/go-api/pkg/limiter"
	"github.com/Zero0719/go-api/pkg/logger"
	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// LimitIP 全局限流中间件，针对 IP 进行限流
// limit 为格式化字符串，如 "5-S" ，示例:
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
//
func LimitIP(limit string) gin.HandlerFunc {

	if app.IsTesting() {
		limit = "1000000-H"
	}
	
	return func(ctx *gin.Context) {
		key := limiter.GetKeyIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(ctx *gin.Context) {
		ctx.Set("limiter-once", false)
		key := limiter.GetKeyRouteWithIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)

	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "请求过于频繁，请稍后再试",
		})
		return false
	}

	return true
}