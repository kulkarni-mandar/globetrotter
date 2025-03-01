package middlewares

import (
	"errors"
	"globetrotter/pkg/logging"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware is a custom panic recovery middleware
func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				logging.Error(
					errors.New("panic recovery"),
					"stackTrace", debug.Stack(),
				)

				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": "Something went wrong while processing",
				})

				c.Abort()
			}
		}()

		c.Next()
	}
}
