package controllers

import (
	"backend/dto"
	"backend/utils"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusLogController struct {
	service services.StatusLogService
}

func NewStatusLogController(service services.StatusLogService) *StatusLogController {
	return &StatusLogController{service: service}
}

func (ctl *StatusLogController) Create(c *gin.Context) {
	var req dto.CreateStatusLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}
	resp, err := ctl.service.Create(req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusCreated, "Created successfully", resp)
}

func (ctl *StatusLogController) GetByID(c *gin.Context) {
	resp, err := ctl.service.GetByID(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}

func (ctl *StatusLogController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}
