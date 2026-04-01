package main

import (
	"backend/config"
	"backend/controllers"
	"backend/repository"
	"backend/routes"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	//load env variables
	godotenv.Load()

	// db connection
	db, err := config.ConnectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// repository initialization
	adminRepo := repository.NewAdminRepository(db)
	rumahSakitRepo := repository.NewRumahSakitRepository(db)
	komponenRepo := repository.NewKomponenDarahRepository(db)
	permintaanRepo := repository.NewPermintaanDarahRepository(db)
	detailRepo := repository.NewDetailPermintaanDarahRepository(db)
	statusLogRepo := repository.NewStatusLogRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)

	// service initialization
	adminSvc := services.NewAdminService(adminRepo)
	rumahSakitSvc := services.NewRumahSakitService(rumahSakitRepo)
	komponenSvc := services.NewKomponenDarahService(komponenRepo)
	permintaanSvc := services.NewPermintaanDarahService(permintaanRepo)
	detailSvc := services.NewDetailPermintaanDarahService(detailRepo)
	statusLogSvc := services.NewStatusLogService(statusLogRepo)
	dashboardSvc := services.NewDashboardService(dashboardRepo)

	// controller initialization
	adminCtl := controllers.NewAdminController(adminSvc)
	rumahSakitCtl := controllers.NewRumahSakitController(rumahSakitSvc)
	komponenCtl := controllers.NewKomponenDarahController(komponenSvc)
	permintaanCtl := controllers.NewPermintaanDarahController(permintaanSvc)
	detailCtl := controllers.NewDetailPermintaanDarahController(detailSvc)
	statusLogCtl := controllers.NewStatusLogController(statusLogSvc)
	dashboardCtl := controllers.NewDashboardController(dashboardSvc)

	routes.RegisterAPIRoutes(
		r,
		adminCtl,
		rumahSakitCtl,
		komponenCtl,
		permintaanCtl,
		detailCtl,
		statusLogCtl,
		dashboardCtl,
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
