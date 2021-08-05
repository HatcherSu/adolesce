package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/yudeguang/ratelimit"
	"time"
)

var ErrorLimitExceeded = errors.New("limit exceeded")

type Rate struct {
	DefaultExpiration              time.Duration
	NumberOfAllowedAccesses        int
	EstimatedNumberOfOnlineUserNum []int
}

func IPLimitMiddleware(rule *ratelimit.Rule) func(c *gin.Context) {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		// 限制IP访问次数
		if !rule.AllowVisitByIP4(ip) {
			_ = c.Error(ErrorLimitExceeded)
			c.AbortWithStatus(429)
			return
		}
		c.Next()
	}
}
