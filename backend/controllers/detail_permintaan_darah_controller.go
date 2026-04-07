package controllers

import (
	"backend/dto"
	"backend/services"
	"backend/utils"
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
		utils.SendValidationError(c, err)
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)

	// If rumah sakit, validate that they can only create detail for their own permintaan
	if userRole == "rumah_sakit" {
		if userID == nil || *userID == "" {
			utils.SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
			return
		}
		// This validation will be done in service with permintaan lookup
	}

	resp, err := ctl.service.CreateWithOwnershipCheck(req, userID, userName, userRole)
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
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
}

func (ctl *DetailPermintaanDarahController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, total, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Data retrieved successfully", resp, limit, offset, total)
}

func (ctl *DetailPermintaanDarahController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}
	var req dto.UpdateDetailPermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendValidationError(c, err)
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)

	resp, err := ctl.service.UpdateWithOwnershipCheck(id, req, userID, userName, userRole)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Updated successfully", resp)
}

func (ctl *DetailPermintaanDarahController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)

	if err := ctl.service.DeleteWithOwnershipCheck(id, userID, userName, userRole); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "true"})
}