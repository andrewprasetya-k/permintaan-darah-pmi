package controllers

import (
	"backend/dto"
	"backend/utils"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service services.AdminService
}

func NewAdminController(service services.AdminService) *AdminController {
	return &AdminController{service: service}
}

func (ctl *AdminController) Create(c *gin.Context) {
	var req dto.CreateAdminRequest
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
	utils.SendSuccess(c, http.StatusCreated, "Admin created successfully", resp)
}

func (ctl *AdminController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Admin retrieved successfully", resp)
}

func (ctl *AdminController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, total, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Admins retrieved successfully", resp, total, limit, offset)
}

func (ctl *AdminController) Update(c *gin.Context) {
	var req dto.UpdateAdminRequest
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
	utils.SendSuccess(c, http.StatusOK, "Admin updated successfully", resp)
}

func (ctl *AdminController) Delete(c *gin.Context) {
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Delete(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Admin deleted successfully", nil)
}

func (ctl *AdminController) Restore(c *gin.Context) {
	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")

	if err := ctl.service.Restore(c.Param("id"), userID, userName, userRole, &userAgent); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Admin restored successfully", nil)
}