package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	service services.DashboardService
}

func NewDashboardController(service services.DashboardService) *DashboardController {
	return &DashboardController{service: service}
}

func (ctl *DashboardController) StatusSummary(c *gin.Context) {
	rumahSakitID := c.Param("rumahSakitID")
	resp, err := ctl.service.StatusSummary(rumahSakitID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}