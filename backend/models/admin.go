package models

import "time"

type Admin struct {
	AdminID       string        `gorm:"column:admin_id;primaryKey;default:(-)" json:"adminId"`
	AdminUsername string        `gorm:"column:admin_username;unique" json:"adminUsername"`
	AdminPassword string        `gorm:"column:admin_password" json:"adminPassword"`
	IsDeleted     bool          `gorm:"column:is_deleted;default:false" json:"isDeleted"`
	DeletedAt     *time.Time    `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
	UpdatedAt     time.Time     `gorm:"column:updated_at" json:"updatedAt"`
	CreatedAt     time.Time     `gorm:"column:created_at" json:"createdAt"`
	AdminRole     AdminRoleEnum `gorm:"column:admin_role;default:admin" json:"adminRole"`
	AdminEmail    string        `gorm:"column:admin_email;unique" json:"adminEmail"`
	AdminNama     string        `gorm:"column:admin_nama" json:"adminName"`
}

func (Admin) TableName() string {
	return "admins"
}
