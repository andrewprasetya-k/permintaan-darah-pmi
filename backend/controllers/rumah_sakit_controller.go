package controllers

import (
	"backend/dto"
	"backend/services"
	"backend/utils"
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
		utils.SendValidationError(c, err)
		return
	}

	// Extract user context
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Create(req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusCreated, "Created successfully", resp)
}

func (ctl *RumahSakitController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
}

func (ctl *RumahSakitController) GetMe(c *gin.Context) {
	userID, _, userRole := utils.ExtractUserFromJWT(c)

	if userRole != "rumah_sakit" {
		utils.SendError(c, http.StatusForbidden, "Only rumah sakit can access this endpoint")
		return
	}

	if userID == nil || *userID == "" {
		utils.SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
		return
	}

	resp, err := ctl.service.GetByID(*userID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
}

func (ctl *RumahSakitController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, total, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Data retrieved successfully", resp, total, limit, offset)
}

func (ctl *RumahSakitController) Update(c *gin.Context) {
	var req dto.UpdateRumahSakitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendValidationError(c, err)
		return
	}

	// Extract user context
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Update(c.Param("id"), req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data updated successfully", resp)
}

func (ctl *RumahSakitController) UpdateMe(c *gin.Context) {
	var req dto.UpdateRumahSakitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendValidationError(c, err)
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if userRole != "rumah_sakit" {
		utils.SendError(c, http.StatusForbidden, "Only rumah sakit can access this endpoint")
		return
	}

	if userID == nil || *userID == "" {
		utils.SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
		return
	}

	resp, err := ctl.service.Update(*userID, req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data updated successfully", resp)
}

func (ctl *RumahSakitController) Delete(c *gin.Context) {
	// Extract user context
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Delete(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", nil)
}

func (ctl *RumahSakitController) Restore(c *gin.Context) {
	// Extract user context
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Restore(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", nil)
}

func (ctl *RumahSakitController) GetDistinctRSNama(c *gin.Context) {
	resp, err := ctl.service.GetDistinctRSNama()
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
}
