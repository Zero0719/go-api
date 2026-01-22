package middlewares

import (
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/Zero0719/go-api/pkg/logger"
	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, true)

				var brokenPipe bool

				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					logger.Error(
						c.Request.URL.Path, 
						zap.Time("time", time.Now()), 
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				logger.Error("recovery from panic",
					zap.Time("time", time.Now()), 
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.Stack("stacktrace"),
				)
				
				response.Abort500(c)
			}
		}()
		c.Next()
	}
}