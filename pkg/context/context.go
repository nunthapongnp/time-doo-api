package ctx

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userIDKey   = "userId"
	tenantIDKey = "tenantId"
	roleKey     = "role"
)

func GetUserID(c *gin.Context) (int64, bool) {
	v, ok := c.Get(userIDKey)
	id, isUint := v.(int64)
	return id, ok && isUint
}

func GetTenantID(c *gin.Context) (int64, bool) {
	v, ok := c.Get(tenantIDKey)
	id, isUint := v.(int64)
	return id, ok && isUint
}

func GetRole(c *gin.Context) (string, bool) {
	v, ok := c.Get(roleKey)
	s, isString := v.(string)
	return s, ok && isString
}

func IsAdmin(c *gin.Context) bool {
	role, ok := GetRole(c)
	return ok && (strings.ToLower(role) == "admin" || strings.ToLower(role) == "owner")
}

func IsSuperAdmin(c *gin.Context) bool {
	role, ok := GetRole(c)
	return ok && strings.ToLower(role) == "superadmin"
}
