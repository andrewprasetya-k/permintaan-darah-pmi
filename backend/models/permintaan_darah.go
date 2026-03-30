package models

import "time"

type PermintaanDarah struct {
	PDID           string                `gorm:"column:pd_id;primaryKey" json:"permintaanDarahId"`
	PDRsID         string                `gorm:"column:pd_rs_id" json:"rumahSakitId"`
	PDPsnID        string                `gorm:"column:pd_psn_id" json:"pasienId"`
	PDStatus       PermintaanStatusEnum  `gorm:"column:pd_status;default:diterima" json:"status"`
	PDCancelReason *string               `gorm:"column:pd_cancel_reason" json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time            `gorm:"column:pd_tgl_permintaan" json:"tanggalPermintaan"`
	CreatedAt      time.Time             `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time             `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt      *time.Time            `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}
