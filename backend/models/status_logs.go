package models

import "time"

type StatusLog struct {
	LogID      string                `gorm:"column:log_id;primaryKey" json:"log_id"`
	LogPdID    string                `gorm:"column:log_pd_id" json:"log_pd_id"`
	LogAdminID string                `gorm:"column:log_admin_id" json:"log_admin_id"`
	LogStatusTo PermintaanStatusEnum `gorm:"column:log_status_to" json:"log_status_to"`
	LogNotes   *string               `gorm:"column:log_notes" json:"log_notes,omitempty"`
	CreatedAt  time.Time             `gorm:"column:created_at" json:"created_at"`
}
