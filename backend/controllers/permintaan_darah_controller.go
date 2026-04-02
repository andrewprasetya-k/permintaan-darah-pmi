package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

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
	resp, err := ctl.service.GetAll(limit, offset)
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
