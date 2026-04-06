package models

import "time"

type SystemAccessLog struct {
	SysLogID       int64      `gorm:"column:sys_log_id;primaryKey;autoIncrement" json:"sysLogId"`
	SysUserID      *string    `gorm:"column:sys_user_id;index" json:"sysUserId,omitempty"`
	SysUserNama    string     `gorm:"column:sys_user_nama;not null" json:"sysUserNama"`
	SysUserRole    string     `gorm:"column:sys_user_role;not null" json:"sysUserRole"`
	
	SysAction      string     `gorm:"column:sys_action;not null;index" json:"sysAction"`
	SysTargetTable *string    `gorm:"column:sys_target_table;index" json:"sysTargetTable,omitempty"`
	SysTargetID    *string    `gorm:"column:sys_target_id" json:"sysTargetId,omitempty"`
	
	SysNotes       string     `gorm:"column:sys_notes;not null" json:"sysNotes"`
	
	SysUserAgent   *string    `gorm:"column:sys_user_agent" json:"sysUserAgent,omitempty"`
	CreatedAt      time.Time  `gorm:"column:created_at;not null;autoCreateTime" json:"createdAt"`
}

func (SystemAccessLog) TableName() string {
	return "system_access_logs"
}
