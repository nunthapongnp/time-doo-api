package middleware

import (
	"time"
	ctx "time-doo-api/pkg/context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		userID, _ := ctx.GetUserID(c)
		tenantID, _ := ctx.GetTenantID(c)

		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.Int64("userId", userID),
			zap.Int64("tenantId", tenantID),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.String("clientIp", c.ClientIP()),
			zap.Duration("latency", latency),
		)
	}
}
