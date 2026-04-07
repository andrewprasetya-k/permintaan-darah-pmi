package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(
	r *gin.Engine,
	adminController *controllers.AdminController,
	rumahSakitController *controllers.RumahSakitController,
	komponenController *controllers.KomponenDarahController,
	permintaanController *controllers.PermintaanDarahController,
	detailController *controllers.DetailPermintaanDarahController,
	statusLogController *controllers.StatusLogController,
	systemAccessLogController *controllers.SystemAccessLogController,
	dashboardController *controllers.DashboardController,
	authController *controllers.AuthController,
) {
	api := r.Group("/api")

	// public routes (no JWT required)
	auth := api.Group("/auth")
	auth.POST("/login/admin", authController.AdminLogin)
	auth.POST("/login/rumah-sakit", authController.RumahSakitLogin)

	// apply JWT middleware to all protected routes
	protected := api.Group("")
	protected.Use(utils.JWTMiddleware())

	// superadmin only
	admins := protected.Group("/admin")
	admins.Use(middleware.SuperAdminOnly())
	admins.POST("", adminController.Create)
	admins.GET("", adminController.GetAll)
	admins.GET("/:id", adminController.GetByID)
	admins.PUT("/:id", adminController.Update)
	admins.DELETE("/:id", adminController.Delete)
	admins.PUT("/restore/:id", adminController.Restore)

	// admin only
	rumahSakit := protected.Group("/rumah-sakit")
	rumahSakit.Use(middleware.AdminOnly())
	rumahSakit.POST("", rumahSakitController.Create)
	rumahSakit.GET("", rumahSakitController.GetAll)
	rumahSakit.GET("/:id", rumahSakitController.GetByID)
	rumahSakit.PUT("/:id", rumahSakitController.Update)
	rumahSakit.DELETE("/:id", rumahSakitController.Delete)
	rumahSakit.PUT("/restore/:id", rumahSakitController.Restore)

	komponen := protected.Group("/komponen-darah")
	komponen.Use(middleware.AdminOnly())
	komponen.POST("", komponenController.Create)
	komponen.GET("", komponenController.GetAll)
	komponen.GET("/:id", komponenController.GetByID)
	komponen.PUT("/:id", komponenController.Update)
	komponen.DELETE("/:id", komponenController.Delete)
	komponen.PUT("/activate/:id", komponenController.Activate)
	komponen.PUT("/deactivate/:id", komponenController.Deactivate)

	// admin only
	statusLogs := protected.Group("/status-logs")
	statusLogs.Use(middleware.AdminOnly())
	statusLogs.GET("", statusLogController.GetAll)

	systemLogs := protected.Group("/system-logs")
	systemLogs.Use(middleware.AdminOnly())
	systemLogs.GET("", systemAccessLogController.GetAll)
	systemLogs.GET("/:id", systemAccessLogController.GetByID)
	systemLogs.GET("/user/:userId", systemAccessLogController.GetByUserID)
	systemLogs.GET("/action/:action", systemAccessLogController.GetByAction)
	systemLogs.GET("/table/:table", systemAccessLogController.GetByTargetTable)
	systemLogs.GET("/target/:targetId", systemAccessLogController.GetByTargetID)

	// admin/rumah sakit
	permintaan := protected.Group("/permintaan-darah")
	permintaan.Use(middleware.AdminOrRumahSakit())
	permintaan.POST("", permintaanController.Create)
	permintaan.GET("", permintaanController.GetAll)
	permintaan.GET("/:id", permintaanController.GetByID)
	permintaan.PUT("/:id", permintaanController.Update)
	permintaan.DELETE("/:id", permintaanController.Delete)
	permintaan.PUT("/restore/:id", permintaanController.Restore)
	permintaan.PUT("/update/:id", permintaanController.UpdateStatus)

	detail := protected.Group("/detail-permintaan-darah")
	detail.Use(middleware.RumahSakitOnly())
	detail.POST("", detailController.Create)
	detail.GET("", detailController.GetAll)
	detail.GET("/:id", detailController.GetByID)
	detail.PUT("/:id", detailController.Update)
	detail.DELETE("/:id", detailController.Delete)

	filter := protected.Group("/filter")
	filter.Use(middleware.AdminOrRumahSakit())
	filter.GET("/rumah-sakit/", rumahSakitController.GetDistinctRSNama)

	dashboard := protected.Group("/dashboard")
	dashboard.Use(middleware.AdminOrRumahSakit())
	dashboard.GET("/status-summary/:rumahSakitID", dashboardController.StatusSummary)
}
