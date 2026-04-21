package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SystemAccessLogController struct {
	service services.SystemAccessLogService
}

func NewSystemAccessLogController(service services.SystemAccessLogService) *SystemAccessLogController {
	return &SystemAccessLogController{service: service}
}

func (ctl *SystemAccessLogController) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	result, count, err := ctl.service.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Operation successful", result, count, limit, offset)
}

func (ctl *SystemAccessLogController) GetByID(c *gin.Context) {
	logID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid log id")
		return
	}

	result, err := ctl.service.GetByID(logID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ctl *SystemAccessLogController) GetByUserID(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		utils.SendError(c, http.StatusBadRequest, "userId is required")
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	result, count, err := ctl.service.GetByUserID(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Operation successful", result, count, limit, offset)
}

func (ctl *SystemAccessLogController) GetByAction(c *gin.Context) {
	action := c.Param("action")
	if action == "" {
		utils.SendError(c, http.StatusBadRequest, "action is required")
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	result, count, err := ctl.service.GetByAction(action, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Operation successful", result, count, limit, offset)
}

func (ctl *SystemAccessLogController) GetByTargetTable(c *gin.Context) {
	targetTable := c.Param("table")
	if targetTable == "" {
		utils.SendError(c, http.StatusBadRequest, "table is required")
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	result, count, err := ctl.service.GetByTargetTable(targetTable, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Operation successful", result, count, limit, offset)
}

func (ctl *SystemAccessLogController) GetByTargetID(c *gin.Context) {
	targetID := c.Param("targetId")
	if targetID == "" {
		utils.SendError(c, http.StatusBadRequest, "targetId is required")
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	result, count, err := ctl.service.GetByTargetID(targetID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	utils.SendSuccessWithPagination(c, http.StatusOK, "Operation successful", result, count, limit, offset)
}
