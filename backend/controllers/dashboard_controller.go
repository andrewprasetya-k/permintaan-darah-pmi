package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"
	"strings"

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
	userID, _, userRole := utils.ExtractUserFromJWT(c)

	if userRole == "rumah_sakit" {
		if userID == nil || *userID == "" {
			utils.SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
			return
		}
		rumahSakitID = *userID
	}

	if strings.EqualFold(rumahSakitID, "all") {
		rumahSakitID = ""
	}

	resp, err := ctl.service.StatusSummary(rumahSakitID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.SendSuccess(c, http.StatusOK, "Operation successful", resp)
}
