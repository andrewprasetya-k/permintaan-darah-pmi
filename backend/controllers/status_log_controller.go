package controllers

import (
	"backend/dto"
	"backend/services"
	"backend/utils"
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
		utils.SendValidationError(c, err)
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
	utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", resp)
}

func (ctl *StatusLogController) GetAll(c *gin.Context) {
	limit, offset := utils.ParsePagination(c)
	resp, total, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Data retrieved successfully", resp, total, limit, offset)
}
