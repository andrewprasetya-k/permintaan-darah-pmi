package utils

import (
	"backend/dto"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// ParsePagination extracts limit and offset from query parameters
func ParsePagination(c *gin.Context) (int, int) {
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

// HandleError handles error response with appropriate HTTP status code
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// Check if it's an AppError
	if appErr := GetAppError(err); appErr != nil && appErr.Code != http.StatusInternalServerError {
		c.JSON(appErr.Code, dto.ErrorResponse(appErr.Message, appErr.Details...))
		return
	}

	// Check if it's a GORM error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse("Data not found"))
		return
	}

	// Default internal server error
	c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Internal server error", err.Error()))
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, dto.SuccessResponse(message, data))
}

// SendSuccessWithPagination sends a success response with pagination info
func SendSuccessWithPagination(c *gin.Context, statusCode int, message string, data interface{}, total int, limit int, offset int) {
	c.JSON(statusCode, dto.SuccessResponseWithPagination(message, data, total, limit, offset))
}

// SendError sends an error response
func SendError(c *gin.Context, statusCode int, message string, errors ...string) {
	c.JSON(statusCode, dto.ErrorResponse(message, errors...))
}

// ExtractUserFromJWT extracts user info from JWT claims
func ExtractUserFromJWT(c *gin.Context) (*string, string, string) {
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
