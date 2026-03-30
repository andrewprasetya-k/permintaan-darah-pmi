package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DetailPermintaanDarahController struct {
	service services.DetailPermintaanDarahService
}

func NewDetailPermintaanDarahController(service services.DetailPermintaanDarahService) *DetailPermintaanDarahController {
	return &DetailPermintaanDarahController{service: service}
}

func (ctl *DetailPermintaanDarahController) Create(c *gin.Context) {
	var req dto.CreateDetailPermintaanDarahRequest
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

func (ctl *DetailPermintaanDarahController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *DetailPermintaanDarahController) GetAll(c *gin.Context) {
	limit, offset := parsePagination(c)
	resp, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *DetailPermintaanDarahController) Update(c *gin.Context) {
	var req dto.UpdateDetailPermintaanDarahRequest
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

func (ctl *DetailPermintaanDarahController) Delete(c *gin.Context) {
	if err := ctl.service.Delete(c.Param("id")); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
