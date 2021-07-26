package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// timeout context
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func(ctx context.Context) {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			cancel()
		}(ctx)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
