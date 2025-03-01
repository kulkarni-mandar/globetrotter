package middlewares

import (
	"globetrotter/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		logging.Debug(
			c.Request.URL.Path,
			"type", c.Request.Method,
			"latency", latency,
			"responseStatus", c.Writer.Status(),
		)
	}
}
