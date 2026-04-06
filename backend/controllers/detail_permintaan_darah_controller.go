package controllers

import (
	"backend/dto"
	"backend/utils"
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
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)

	resp, err := ctl.service.Create(req, userID, userName, userRole)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusCreated, "Created successfully", resp)
}

func (ctl *DetailPermintaanDarahController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := ctl.service.GetByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *DetailPermintaanDarahController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *DetailPermintaanDarahController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}
	var req dto.UpdateDetailPermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)

	resp, err := ctl.service.Update(id, req, userID, userName, userRole)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *DetailPermintaanDarahController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)

	if err := ctl.service.Delete(id, userID, userName, userRole); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "true"})
}