package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	resp, err := ctl.service.Create(req)
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
	limit, offset := parsePagination(c)
	
	filters := &dto.PermintaanDarahFilters{}
	if status := c.Query("status"); status != "" {
		filters.Status = &status
	}
	if rsID := c.Query("id"); rsID != "" {
		filters.RsID = &rsID
	}
	if gender := c.Query("gender"); gender != "" {
		filters.Gender = &gender
	}
	if golDarah := c.Query("gol_darah"); golDarah != "" {
		filters.GolDarah = &golDarah
	}
	if startDate := c.Query("start_date"); startDate != "" {
		if t, err := time.Parse("2006-01-02", startDate); err == nil {
			filters.StartDate = &t
		}
	}
	if endDate := c.Query("end_date"); endDate != "" {
		if t, err := time.Parse("2006-01-02", endDate); err == nil {
			filters.EndDate = &t
		}
	}
	
	resp, err := ctl.service.GetAll(filters, limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *PermintaanDarahController) GetByRsID(c *gin.Context) {
	limit, offset := parsePagination(c)
	resp, err := ctl.service.GetByRsID(c.Param("rsId"), limit, offset)
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
	resp, err := ctl.service.Update(c.Param("id"), req)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *PermintaanDarahController) Delete(c *gin.Context) {
	if err := ctl.service.Delete(c.Param("id")); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (ctl *PermintaanDarahController) UpdateStatus(c *gin.Context) {
	var req dto.UpdatePermintaanStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err := ctl.service.UpdateStatus(c.Param("id"), req.Status, req.Reason)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
