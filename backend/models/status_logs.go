package models

import "time"

type StatusLog struct {
	LogID        int64                   `gorm:"column:log_id;primaryKey;autoIncrement" json:"logId"`
	LogPdID      string                  `gorm:"column:log_pd_id;not null;index" json:"permintaanDarahId"`
	LogAdminID   *string                 `gorm:"column:log_admin_id;index" json:"adminId,omitempty"`
	LogAdminNama string                  `gorm:"column:log_admin_nama;not null" json:"adminNama"`
	
	LogStatusFrom *PermintaanStatusEnum  `gorm:"column:log_status_from" json:"statusFrom,omitempty"`
	LogStatusTo   PermintaanStatusEnum   `gorm:"column:log_status_to;not null" json:"statusTo"`
	
	LogNotes  *string   `gorm:"column:log_notes" json:"notes,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime" json:"createdAt"`
}

func (StatusLog) TableName() string {
	return "status_logs"
}
