package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/services"
	"backend/utils"
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
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Create(req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusCreated, "Created successfully", resp)
}

func (ctl *PermintaanDarahController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
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

	limit, offset := utils.ParsePagination(c)
	resp, total, err := ctl.service.GetAll(filters, limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Data retrieved successfully", resp, total, limit, offset)
}

func (ctl *PermintaanDarahController) GetMyRequests(c *gin.Context) {
	userID, _, userRole := utils.ExtractUserFromJWT(c)

	if userRole != "rumah_sakit" {
		utils.SendError(c, http.StatusForbidden, "Only rumah sakit can access this endpoint")
		return
	}

	if userID == nil || *userID == "" {
		utils.SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
		return
	}

	limit, offset := utils.ParsePagination(c)
	resp, total, err := ctl.service.GetByRsID(*userID, limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Data retrieved successfully", resp, total, limit, offset)
}

func (ctl *PermintaanDarahController) Update(c *gin.Context) {
	var req dto.UpdatePermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.Update(c.Param("id"), req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
}

func (ctl *PermintaanDarahController) Delete(c *gin.Context) {
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if userRole != "admin" && userRole != "superadmin" {
		utils.SendError(c, http.StatusForbidden, "Only admin can delete permintaan darah")
		return
	}

	if err := ctl.service.Delete(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", nil)
}

func (ctl *PermintaanDarahController) UpdateStatus(c *gin.Context) {
	var req struct {
		Status string  `json:"status" binding:"required"`
		Reason *string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	resp, err := ctl.service.UpdateStatusWithOwnershipCheck(c.Param("id"), models.PermintaanStatusEnum(req.Status), req.Reason, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Status updated successfully", resp)
}

func (ctl *PermintaanDarahController) UpdateMyRequest(c *gin.Context) {
	var req dto.UpdatePermintaanDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if userRole != "rumah_sakit" {
		utils.SendError(c, http.StatusForbidden, "Only rumah sakit can update their own requests")
		return
	}

	if userID == nil || *userID == "" {
		utils.SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
		return
	}

	resp, err := ctl.service.UpdateMyRequest(c.Param("id"), *userID, req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Request updated successfully", resp)
}