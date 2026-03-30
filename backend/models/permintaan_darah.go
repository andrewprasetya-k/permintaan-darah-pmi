package models

import "time"

type PermintaanDarah struct {
	PDID           string                `gorm:"column:pd_id;primaryKey" json:"pd_id"`
	PDRsID         string                `gorm:"column:pd_rs_id" json:"pd_rs_id"`
	PDPsnID        string                `gorm:"column:pd_psn_id" json:"pd_psn_id"`
	PDStatus       PermintaanStatusEnum  `gorm:"column:pd_status;default:diterima" json:"pd_status"`
	PDCancelReason *string               `gorm:"column:pd_cancel_reason" json:"pd_cancel_reason,omitempty"`
	PDTglPermintaan time.Time            `gorm:"column:pd_tgl_permintaan" json:"pd_tgl_permintaan"`
	CreatedAt      time.Time             `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time             `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      *time.Time            `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}
