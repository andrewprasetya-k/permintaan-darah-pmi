package middleware

import (
	"github.com/gin-gonic/gin"
)

func SuperAdminOnly() gin.HandlerFunc{
	return func (c *gin.Context) { 
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "super_admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Super Admins only"})
			return
		}
	}
}

func AdminOnly() gin.HandlerFunc{
	return func (c *gin.Context) { 
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins only"})
			return
		}
	}
}

func RumahSakitOnly() gin.HandlerFunc{
	return func (c *gin.Context) { 
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "rumah_sakit" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Rumah Sakit only"})
			return
		}
	}
}

func AdminOrRumahSakit() gin.HandlerFunc{
	return func (c *gin.Context) { 
		userRole, exists := c.Get("userRole")
		if !exists || (userRole != "admin" && userRole != "rumah_sakit") {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins or Rumah Sakit only"})
			return
		}
	}
}