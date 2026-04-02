package dto

import (
	"backend/models"
	"time"
)

type CreateStatusLogRequest struct {
	LogPdID     string                      `json:"permintaanDarahId" binding:"required"`
	LogAdminID  string                      `json:"adminId" binding:"required"`
	LogStatusTo models.PermintaanStatusEnum `json:"statusTo" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	LogNotes    *string                     `json:"notes,omitempty"`
}

type UpdateStatusLogRequest struct {
	LogPdID     string                      `json:"permintaanDarahId" binding:"required"`
	LogAdminID  string                      `json:"adminId" binding:"required"`
	LogStatusTo models.PermintaanStatusEnum `json:"statusTo" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	LogNotes    *string                     `json:"notes,omitempty"`
}

type StatusLogResponse struct {
	LogID       string                      `json:"logId"`
	LogPdID     string                      `json:"permintaanDarahId"`
	LogAdminID  string                      `json:"adminId"`
	LogStatusTo models.PermintaanStatusEnum `json:"statusTo"`
	LogNotes    *string                     `json:"notes,omitempty"`
	CreatedAt   time.Time                   `json:"createdAt"`
}
