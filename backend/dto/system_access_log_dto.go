package dto

import (
	"time"
)

type CreateSystemAccessLogRequest struct {
	SysUserID      *string `json:"sysUserId,omitempty"`
	SysUserNama    string  `json:"sysUserNama" binding:"required"`
	SysUserRole    string  `json:"sysUserRole" binding:"required,oneof=admin rumah_sakit"`
	
	SysAction      string  `json:"sysAction" binding:"required,oneof=CREATE UPDATE DELETE SOFT_DELETE RESTORE LOGIN"`
	SysTargetTable *string `json:"sysTargetTable,omitempty"`
	SysTargetID    *string `json:"sysTargetId,omitempty"`
	
	SysNotes       string  `json:"sysNotes" binding:"required"`
	
	SysUserAgent   *string `json:"sysUserAgent,omitempty"`
}

type SystemAccessLogResponse struct {
	SysLogID       int64     `json:"sysLogId"`
	SysUserID      *string   `json:"sysUserId,omitempty"`
	SysUserNama    string    `json:"sysUserNama"`
	SysUserRole    string    `json:"sysUserRole"`
	
	SysAction      string    `json:"sysAction"`
	SysTargetTable *string   `json:"sysTargetTable,omitempty"`
	SysTargetID    *string   `json:"sysTargetId,omitempty"`
	
	SysNotes       string    `json:"sysNotes"`
	
	SysUserAgent   *string   `json:"sysUserAgent,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
}

type SystemAccessLogListResponse struct {
	Data     []SystemAccessLogResponse `json:"data"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"pageSize"`
}

