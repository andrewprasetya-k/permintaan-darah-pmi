package models

import "time"

type StatusLog struct {
	LogID      string                `gorm:"column:log_id;primaryKey" json:"logId"`
	LogPdID    string                `gorm:"column:log_pd_id" json:"permintaanDarahId"`
	LogAdminID string                `gorm:"column:log_admin_id" json:"adminId"`
	LogStatusTo PermintaanStatusEnum `gorm:"column:log_status_to" json:"statusTo"`
	LogNotes   *string               `gorm:"column:log_notes" json:"notes,omitempty"`
	CreatedAt  time.Time             `gorm:"column:created_at" json:"createdAt"`
}
