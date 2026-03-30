package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(
	r *gin.Engine,
	adminController *controllers.AdminController,
	rumahSakitController *controllers.RumahSakitController,
	pasienController *controllers.PasienController,
	komponenController *controllers.KomponenDarahController,
	permintaanController *controllers.PermintaanDarahController,
	detailController *controllers.DetailPermintaanDarahController,
	statusLogController *controllers.StatusLogController,
) {
	api := r.Group("/api")

	admins := api.Group("/admin")
	admins.POST("", adminController.Create)
	admins.GET("", adminController.GetAll)
	admins.GET("/:id", adminController.GetByID)
	admins.PUT("/:id", adminController.Update)
	admins.DELETE("/:id", adminController.Delete)

	rumahSakit := api.Group("/rumah-sakit")
	rumahSakit.POST("", rumahSakitController.Create)
	rumahSakit.GET("", rumahSakitController.GetAll)
	rumahSakit.GET("/:id", rumahSakitController.GetByID)
	rumahSakit.PUT("/:id", rumahSakitController.Update)
	rumahSakit.DELETE("/:id", rumahSakitController.Delete)

	pasien := api.Group("/pasien")
	pasien.POST("", pasienController.Create)
	pasien.GET("", pasienController.GetAll)
	pasien.GET("/:id", pasienController.GetByID)
	pasien.PUT("/:id", pasienController.Update)
	pasien.DELETE("/:id", pasienController.Delete)

	komponen := api.Group("/komponen-darah")
	komponen.POST("", komponenController.Create)
	komponen.GET("", komponenController.GetAll)
	komponen.GET("/:id", komponenController.GetByID)
	komponen.PUT("/:id", komponenController.Update)
	komponen.DELETE("/:id", komponenController.Delete)

	permintaan := api.Group("/permintaan-darah")
	permintaan.POST("", permintaanController.Create)
	permintaan.GET("", permintaanController.GetAll)
	permintaan.GET("/:id", permintaanController.GetByID)
	permintaan.PUT("/:id", permintaanController.Update)
	permintaan.DELETE("/:id", permintaanController.Delete)

	detail := api.Group("/detail-permintaan-darah")
	detail.POST("", detailController.Create)
	detail.GET("", detailController.GetAll)
	detail.GET("/:id", detailController.GetByID)
	detail.PUT("/:id", detailController.Update)
	detail.DELETE("/:id", detailController.Delete)

	statusLogs := api.Group("/status-logs")
	statusLogs.POST("", statusLogController.Create)
	statusLogs.GET("", statusLogController.GetAll)
	statusLogs.GET("/:id", statusLogController.GetByID)
	statusLogs.PUT("/:id", statusLogController.Update)
	statusLogs.DELETE("/:id", statusLogController.Delete)
}
