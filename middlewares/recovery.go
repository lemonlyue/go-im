package middlewares

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
	"go.uber.org/zap"
	"gin-skeleton/pkg/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// get user request info
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				// connection break off
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
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
					logger.Error(c.Request.URL.Path,
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

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "网络错误,请稍后重试",
				})
			}
		}()
		c.Next()
	}
}
