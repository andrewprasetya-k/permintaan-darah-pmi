package dto

import (
	"backend/models"
	"time"
)

type CreatePermintaanDarahRequest struct {
	PDID            string                      `json:"permintaanDarahId" binding:"required"`
	PDRsID          string                      `json:"rumahSakitId" binding:"required"`
	PDPsnID         string                      `json:"pasienId" binding:"required"`
	PDStatus        models.PermintaanStatusEnum `json:"status" binding:"required,oneof=diterima diproses bisa_diambil selesai dibatalkan"`
	PDCancelReason  *string                     `json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time                   `json:"tanggalPermintaan" binding:"required"`
}

type UpdatePermintaanDarahRequest struct {
	PDRsID          string                      `json:"rumahSakitId" binding:"required"`
	PDPsnID         string                      `json:"pasienId" binding:"required"`
	PDStatus        models.PermintaanStatusEnum `json:"status" binding:"required,oneof=diterima diproses bisa_diambil selesai dibatalkan"`
	PDCancelReason  *string                     `json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time                   `json:"tanggalPermintaan" binding:"required"`
}

type PermintaanDarahResponse struct {
	PDID            string                      `json:"permintaanDarahId"`
	PDRsID          string                      `json:"rumahSakitId"`
	PDPsnID         string                      `json:"pasienId"`
	PDStatus        models.PermintaanStatusEnum `json:"status"`
	PDCancelReason  *string                     `json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time                   `json:"tanggalPermintaan"`
	CreatedAt       time.Time                   `json:"createdAt"`
	UpdatedAt       time.Time                   `json:"updatedAt"`
	DeletedAt       *time.Time                  `json:"deletedAt,omitempty"`
}
