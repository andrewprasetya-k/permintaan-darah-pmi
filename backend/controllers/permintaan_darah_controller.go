package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PermintaanDarahController struct {
	service services.PermintaanDarahService
}

func NewPermintaanDarahController(service services.PermintaanDarahService) *PermintaanDarahController {
	return &PermintaanDarahController{service: service}
}

func (ctl *PermintaanDarahController) Create(c *gin.Context) {
	var req dto.CreatePermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Create(req, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (ctl *PermintaanDarahController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *PermintaanDarahController) GetAll(c *gin.Context) {
	status := c.Query("status")
	rsID := c.Query("rsID")
	golDarah := c.Query("golDarah")
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	var startDate, endDate *time.Time
	if startDateStr != "" {
		if t, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &t
		}
	}
	if endDateStr != "" {
		if t, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &t
		}
	}

	filters := &dto.PermintaanDarahFilters{
		Status:    &status,
		RsID:      &rsID,
		GolDarah:  &golDarah,
		StartDate: startDate,
		EndDate:   endDate,
	}

	limit, offset := parsePagination(c)
	resp, err := ctl.service.GetAll(filters, limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *PermintaanDarahController) GetByRsID(c *gin.Context) {
	limit, offset := parsePagination(c)
	resp, err := ctl.service.GetByRsID(c.Param("rsID"), limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *PermintaanDarahController) Update(c *gin.Context) {
	var req dto.UpdatePermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Update(c.Param("id"), req, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *PermintaanDarahController) Delete(c *gin.Context) {
	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Delete(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctl *PermintaanDarahController) Restore(c *gin.Context) {
	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Restore(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctl *PermintaanDarahController) UpdateStatus(c *gin.Context) {
	var req struct {
		Status string  `json:"status" binding:"required"`
		Reason *string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.UpdateStatus(c.Param("id"), models.PermintaanStatusEnum(req.Status), req.Reason, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func extractUserFromJWT(c *gin.Context) (*string, string, string) {
	claims, _ := c.Get("claims")
	if claims == nil {
		return nil, "Unknown User", "unknown"
	}

	jwtClaims := claims.(jwt.MapClaims)

	userID, _ := jwtClaims["admin_id"].(string)
	userName, _ := jwtClaims["admin_nama"].(string)
	userRole, _ := jwtClaims["admin_role"].(string)

	return &userID, userName, userRole
}