package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RumahSakitController struct {
	service services.RumahSakitService
}

func NewRumahSakitController(service services.RumahSakitService) *RumahSakitController {
	return &RumahSakitController{service: service}
}

func (ctl *RumahSakitController) Create(c *gin.Context) {
	var req dto.CreateRumahSakitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Extract user context
	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Create(req, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (ctl *RumahSakitController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *RumahSakitController) GetAll(c *gin.Context) {
	limit, offset := parsePagination(c)
	resp, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *RumahSakitController) Update(c *gin.Context) {
	var req dto.UpdateRumahSakitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Extract user context
	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Update(c.Param("id"), req, userID, userName, userRole, &userAgent)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *RumahSakitController) Delete(c *gin.Context) {
	// Extract user context
	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Delete(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctl *RumahSakitController) Restore(c *gin.Context) {
	// Extract user context
	userID, userName, userRole := extractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Restore(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctl *RumahSakitController) GetDistinctRSNama(c *gin.Context) {
	resp, err := ctl.service.GetDistinctRSNama()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
