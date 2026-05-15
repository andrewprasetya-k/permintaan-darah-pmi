package middleware

import (
	"github.com/gin-gonic/gin"
)

func SuperAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "superadmin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Hanya Super Admin"})
			return
		}
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || (userRole != "admin" && userRole != "superadmin") {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Hanya Admin"})
			return
		}
		c.Next()
	}
}

func RumahSakitOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "rumah_sakit" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Hanya Rumah Sakit"})
			return
		}
		c.Next()
	}
}

func AdminOrRumahSakit() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || (userRole != "admin" && userRole != "superadmin" && userRole != "rumah_sakit") {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Hanya Admin atau Rumah Sakit"})
			return
		}
		c.Next()
	}
}