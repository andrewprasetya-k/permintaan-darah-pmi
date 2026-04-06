package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

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

	userID, userName, userRole := extractUserFromContext(c)

	resp, err := ctl.service.Create(req, userID, userName, userRole)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (ctl *DetailPermintaanDarahController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	resp, err := ctl.service.GetByID(id)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var req dto.UpdateDetailPermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, userName, userRole := extractUserFromContext(c)

	resp, err := ctl.service.Update(id, req, userID, userName, userRole)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl *DetailPermintaanDarahController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	userID, userName, userRole := extractUserFromContext(c)

	if err := ctl.service.Delete(id, userID, userName, userRole); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "true"})
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