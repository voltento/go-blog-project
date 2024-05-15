package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		slog.Info("Request handled",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"ip", c.ClientIP(),
			"user-agent", c.Request.UserAgent(),
			"status", c.Writer.Status(),
			"duration", duration,
		)
	}
}
