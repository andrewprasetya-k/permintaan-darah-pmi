package controllers

import (
	"backend/dto"
	"backend/utils"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KomponenDarahController struct {
	service services.KomponenDarahService
}

func NewKomponenDarahController(service services.KomponenDarahService) *KomponenDarahController {
	return &KomponenDarahController{service: service}
}

func (ctl *KomponenDarahController) Create(c *gin.Context) {
	var req dto.CreateKomponenDarahRequest
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

func (ctl *KomponenDarahController) GetByID(c *gin.Context) {
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

func (ctl *KomponenDarahController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *KomponenDarahController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}
	var req dto.UpdateKomponenDarahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")
	resp, err := ctl.service.Update(id, req, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *KomponenDarahController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")
	if err := ctl.service.Delete(id, userID, userName, userRole, &userAgent); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (ctl *KomponenDarahController) Activate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")
	resp, err := ctl.service.ActivateKomponenDarah(id, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *KomponenDarahController) Deactivate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	userID, userName, userRole := utils.ExtractUserFromJWT(c)
	userAgent := c.GetHeader("User-Agent")
	resp, err := ctl.service.DeactivateKomponenDarah(id, userID, userName, userRole, &userAgent)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}
