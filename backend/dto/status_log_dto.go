package dto

import (
	"backend/models"
	"time"
)

type CreateStatusLogRequest struct {
	LogPdID       string                      `json:"permintaanDarahId" binding:"required"`
	LogAdminID    *string                     `json:"adminId,omitempty"`
	LogAdminNama  string                      `json:"adminNama" binding:"required"`
	LogStatusFrom *models.PermintaanStatusEnum `json:"statusFrom,omitempty"`
	LogStatusTo   models.PermintaanStatusEnum `json:"statusTo" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	LogNotes      *string                     `json:"notes,omitempty"`
}

type UpdateStatusLogRequest struct {
	LogPdID       string                      `json:"permintaanDarahId" binding:"required"`
	LogAdminID    *string                     `json:"adminId,omitempty"`
	LogAdminNama  string                      `json:"adminNama" binding:"required"`
	LogStatusFrom *models.PermintaanStatusEnum `json:"statusFrom,omitempty"`
	LogStatusTo   models.PermintaanStatusEnum `json:"statusTo" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	LogNotes      *string                     `json:"notes,omitempty"`
}

type StatusLogResponse struct {
	LogID         int64                       `json:"logId"`
	LogPdID       string                      `json:"permintaanDarahId"`
	LogAdminID    *string                     `json:"adminId,omitempty"`
	LogAdminNama  string                      `json:"adminNama"`
	LogStatusFrom *models.PermintaanStatusEnum `json:"statusFrom,omitempty"`
	LogStatusTo   models.PermintaanStatusEnum `json:"statusTo"`
	LogNotes      *string                     `json:"notes,omitempty"`
	CreatedAt     time.Time                   `json:"createdAt"`
}

type AuditTrailResponse struct {
	PermintaanDarahId string               `json:"permintaanDarahId"`
	Logs              []StatusLogResponse  `json:"logs"`
	TotalLogs         int                  `json:"totalLogs"`
}
