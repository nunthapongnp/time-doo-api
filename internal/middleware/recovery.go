package middleware

import (
	"errors"
	"net/http"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryWithLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic recovered", zap.Any("error", r))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
		}()
		c.Next()
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				response.Error(c, errToError(r), http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func errToError(r interface{}) error {
	switch x := r.(type) {
	case string:
		return errors.New(x)
	case error:
		return x
	default:
		return errors.New("unknown error")
	}
}
