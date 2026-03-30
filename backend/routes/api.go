package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(
	r *gin.Engine,
	adminCtl *controllers.AdminController,
	rumahSakitCtl *controllers.RumahSakitController,
	pasienCtl *controllers.PasienController,
	komponenCtl *controllers.KomponenDarahController,
	permintaanCtl *controllers.PermintaanDarahController,
	detailCtl *controllers.DetailPermintaanDarahController,
	statusLogCtl *controllers.StatusLogController,
) {
	api := r.Group("/api")

	admins := api.Group("/admins")
	admins.POST("", adminCtl.Create)
	admins.GET("", adminCtl.GetAll)
	admins.GET("/:id", adminCtl.GetByID)
	admins.PUT("/:id", adminCtl.Update)
	admins.DELETE("/:id", adminCtl.Delete)

	rumahSakit := api.Group("/rumah-sakit")
	rumahSakit.POST("", rumahSakitCtl.Create)
	rumahSakit.GET("", rumahSakitCtl.GetAll)
	rumahSakit.GET("/:id", rumahSakitCtl.GetByID)
	rumahSakit.PUT("/:id", rumahSakitCtl.Update)
	rumahSakit.DELETE("/:id", rumahSakitCtl.Delete)

	pasien := api.Group("/pasien")
	pasien.POST("", pasienCtl.Create)
	pasien.GET("", pasienCtl.GetAll)
	pasien.GET("/:id", pasienCtl.GetByID)
	pasien.PUT("/:id", pasienCtl.Update)
	pasien.DELETE("/:id", pasienCtl.Delete)

	komponen := api.Group("/komponen-darah")
	komponen.POST("", komponenCtl.Create)
	komponen.GET("", komponenCtl.GetAll)
	komponen.GET("/:id", komponenCtl.GetByID)
	komponen.PUT("/:id", komponenCtl.Update)
	komponen.DELETE("/:id", komponenCtl.Delete)

	permintaan := api.Group("/permintaan-darah")
	permintaan.POST("", permintaanCtl.Create)
	permintaan.GET("", permintaanCtl.GetAll)
	permintaan.GET("/:id", permintaanCtl.GetByID)
	permintaan.PUT("/:id", permintaanCtl.Update)
	permintaan.DELETE("/:id", permintaanCtl.Delete)

	detail := api.Group("/detail-permintaan-darah")
	detail.POST("", detailCtl.Create)
	detail.GET("", detailCtl.GetAll)
	detail.GET("/:id", detailCtl.GetByID)
	detail.PUT("/:id", detailCtl.Update)
	detail.DELETE("/:id", detailCtl.Delete)

	statusLogs := api.Group("/status-logs")
	statusLogs.POST("", statusLogCtl.Create)
	statusLogs.GET("", statusLogCtl.GetAll)
	statusLogs.GET("/:id", statusLogCtl.GetByID)
	statusLogs.PUT("/:id", statusLogCtl.Update)
	statusLogs.DELETE("/:id", statusLogCtl.Delete)
}
