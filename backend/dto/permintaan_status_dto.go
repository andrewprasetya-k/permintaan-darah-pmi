package dto

import "backend/models"

type UpdatePermintaanStatusRequest struct {
	Status models.PermintaanStatusEnum `json:"status" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	Reason *string                     `json:"reason,omitempty"`
}
