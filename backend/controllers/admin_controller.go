package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AdminController struct {
	service services.AdminService
}

func NewAdminController(service services.AdminService) *AdminController {
	return &AdminController{service: service}
}

func (ctl *AdminController) Create(c *gin.Context) {
	var req dto.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Extract user context
	userID, userName, userRole := extractAdminFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Create(req, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (ctl *AdminController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *AdminController) GetAll(c *gin.Context) {
	limit, offset := parsePagination(c)
	resp, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *AdminController) Update(c *gin.Context) {
	var req dto.UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Extract user context
	userID, userName, userRole := extractAdminFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Update(c.Param("id"), req, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *AdminController) Delete(c *gin.Context) {
	// Extract user context
	userID, userName, userRole := extractAdminFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Delete(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctl *AdminController) Restore(c *gin.Context) {
	// Extract user context
	userID, userName, userRole := extractAdminFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Restore(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Helper function to extract admin from JWT
func extractAdminFromJWT(c *gin.Context) (*string, string, string) {
	claims, _ := c.Get("claims")
	if claims == nil {
		return nil, "Unknown User", "unknown"
	}

	jwtClaims := claims.(jwt.MapClaims)

	adminID, _ := jwtClaims["admin_id"].(string)
	adminNama, _ := jwtClaims["admin_nama"].(string)
	adminRole, _ := jwtClaims["admin_role"].(string)

	return &adminID, adminNama, adminRole
}
