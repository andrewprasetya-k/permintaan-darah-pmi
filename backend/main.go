package main

import (
	"backend/config"
	"backend/controllers"
	"backend/repository"
	"backend/routes"
	"backend/services"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	godotenv.Load()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if isOriginAllowed(origin) {
			if origin != "" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				c.Writer.Header().Set("Vary", "Origin")
			}
		} else {
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

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
	statusLogRepo := repository.NewStatusLogRepository(db)
	systemAccessLogRepo := repository.NewSystemAccessLogRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)

	// Initialize WebSocket hub
	wsHub := services.NewHub()
	go wsHub.Run()

	systemAccessLogSvc := services.NewSystemAccessLogService(systemAccessLogRepo, wsHub)
	permintaanSvc := services.NewPermintaanDarahService(permintaanRepo, statusLogRepo, systemAccessLogSvc, wsHub, db)
	statusLogSvc := services.NewStatusLogService(statusLogRepo)
	dashboardSvc := services.NewDashboardService(dashboardRepo)

	// service initialization
	authSvc := services.NewAuthService(adminRepo, rumahSakitRepo, systemAccessLogSvc)
	adminSvc := services.NewAdminService(adminRepo, systemAccessLogSvc)
	rumahSakitSvc := services.NewRumahSakitService(rumahSakitRepo, systemAccessLogSvc)
	komponenSvc := services.NewKomponenDarahService(komponenRepo, systemAccessLogSvc)

	// controller initialization
	authCtl := controllers.NewAuthController(&authSvc)
	adminCtl := controllers.NewAdminController(adminSvc)
	rumahSakitCtl := controllers.NewRumahSakitController(rumahSakitSvc)
	komponenCtl := controllers.NewKomponenDarahController(komponenSvc)
	permintaanCtl := controllers.NewPermintaanDarahController(permintaanSvc)
	statusLogCtl := controllers.NewStatusLogController(statusLogSvc)
	systemAccessLogCtl := controllers.NewSystemAccessLogController(systemAccessLogSvc)
	dashboardCtl := controllers.NewDashboardController(dashboardSvc)

	wsCtl := controllers.NewWebSocketController(wsHub)

	routes.RegisterAPIRoutes(
		r,
		adminCtl,
		rumahSakitCtl,
		komponenCtl,
		permintaanCtl,
		statusLogCtl,
		systemAccessLogCtl,
		dashboardCtl,
		authCtl,
		wsCtl,
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}

func isOriginAllowed(origin string) bool {
	if origin == "" {
		return true
	}

	allowedOrigins := strings.TrimSpace(os.Getenv("CORS_ALLOWED_ORIGIN"))
	if allowedOrigins == "" {
		return gin.Mode() != gin.ReleaseMode
	}

	for _, allowed := range strings.Split(allowedOrigins, ",") {
		allowed = strings.TrimSpace(allowed)
		if allowed == "*" && gin.Mode() != gin.ReleaseMode {
			return true
		}
		if allowed == origin {
			return true
		}
	}

	return false
}
