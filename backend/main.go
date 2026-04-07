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
	
	//CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

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
	systemAccessLogRepo := repository.NewSystemAccessLogRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)

	// service initialization
	systemAccessLogSvc := services.NewSystemAccessLogService(systemAccessLogRepo)
	authSvc := services.NewAuthService(adminRepo, rumahSakitRepo, systemAccessLogSvc)
	adminSvc := services.NewAdminService(adminRepo, systemAccessLogSvc)
	rumahSakitSvc := services.NewRumahSakitService(rumahSakitRepo, systemAccessLogSvc)
	komponenSvc := services.NewKomponenDarahService(komponenRepo, systemAccessLogSvc)
	permintaanSvc := services.NewPermintaanDarahService(permintaanRepo, statusLogRepo, systemAccessLogSvc)
	detailSvc := services.NewDetailPermintaanDarahService(detailRepo, systemAccessLogSvc)
	statusLogSvc := services.NewStatusLogService(statusLogRepo)
	dashboardSvc := services.NewDashboardService(dashboardRepo)

	// controller initialization
	authCtl := controllers.NewAuthController(&authSvc)
	adminCtl := controllers.NewAdminController(adminSvc)
	rumahSakitCtl := controllers.NewRumahSakitController(rumahSakitSvc)
	komponenCtl := controllers.NewKomponenDarahController(komponenSvc)
	permintaanCtl := controllers.NewPermintaanDarahController(permintaanSvc)
	detailCtl := controllers.NewDetailPermintaanDarahController(detailSvc)
	statusLogCtl := controllers.NewStatusLogController(statusLogSvc)
	systemAccessLogCtl := controllers.NewSystemAccessLogController(systemAccessLogSvc)
	dashboardCtl := controllers.NewDashboardController(dashboardSvc)

	routes.RegisterAPIRoutes(
		r,
		adminCtl,
		rumahSakitCtl,
		komponenCtl,
		permintaanCtl,
		detailCtl,
		statusLogCtl,
		systemAccessLogCtl,
		dashboardCtl,
		authCtl,
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
