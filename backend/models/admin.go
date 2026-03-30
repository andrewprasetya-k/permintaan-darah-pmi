package models

type Admin struct {
	AdminID      string         `gorm:"column:admin_id;primaryKey" json:"adminId"`
	AdminUsername string        `gorm:"column:admin_username;unique" json:"adminUsername"`
	AdminPassword string        `gorm:"column:admin_password" json:"adminPassword"`
	DeletedAt     *int64        `gorm:"column:deleted_at" json:"deletedAt,omitempty"`	
	IsDeleted     bool          `gorm:"column:is_deleted;default:false" json:"isDeleted"`	
	UpdatedAt     int64         `gorm:"column:updated_at" json:"updatedAt"`	
	CreatedAt     int64         `gorm:"column:created_at" json:"createdAt"`	
	AdminRole     AdminRoleEnum `gorm:"column:admin_role;default:admin" json:"adminRole"`	
	AdminEmail    string        `gorm:"column:admin_email;unique" json:"adminEmail"`	
	AdminNama     string        `gorm:"column:admin_nama" json:"adminName"`








}	