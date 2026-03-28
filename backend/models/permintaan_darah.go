package models

type PermintaanDarah struct {
	PDID           string                `gorm:"column:pd_id;primaryKey" json:"pd_id"`
	PDRsID         string                `gorm:"column:pd_rs_id" json:"pd_rs_id"`
	PDPsnID        string                `gorm:"column:pd_psn_id" json:"pd_psn_id"`
	PDStatus       PermintaanStatusEnum  `gorm:"column:pd_status;default:diterima" json:"pd_status"`
	PDCancelReason *string               `gorm:"column:pd_cancel_reason" json:"pd_cancel_reason,omitempty"`
	PDTglPermintaan int64                `gorm:"column:pd_tgl_permintaan" json:"pd_tgl_permintaan"`
	CreatedAt      int64                 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      int64                 `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted      bool                  `gorm:"column:is_deleted;default:false" json:"is_deleted"`
	DeletedAt      *int64                `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}
