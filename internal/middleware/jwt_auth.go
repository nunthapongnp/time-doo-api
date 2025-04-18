package middleware

import (
	"context"
	"net/http"
	"strings"
	"time-doo-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	userIDKey   = "userId"
	tenantIDKey = "tenantId"
	roleKey     = "role"
)

type ctxKey string

const UserCtxKey ctxKey = "userId"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set(userIDKey, int64(claims.UserID))
		c.Set(tenantIDKey, int64(claims.TenantID))
		c.Set(roleKey, claims.Role)

		ctx := context.WithValue(c.Request.Context(), UserCtxKey, int64(claims.UserID))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
