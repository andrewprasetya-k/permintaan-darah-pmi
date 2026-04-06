package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func parsePagination(c *gin.Context) (int, int) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

func handleError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
}

func extractUserFromJWT(c *gin.Context) (*string, string, string) {
	claims, _ := c.Get("claims")
	if claims == nil {
		return nil, "Unknown User", "unknown"
	}

	jwtClaims := claims.(jwt.MapClaims)

	userID, _ := jwtClaims["user_id"].(string)
	userName, _ := jwtClaims["username"].(string)
	userRole, _ := jwtClaims["role"].(string)

	return &userID, userName, userRole
}
