package models

type Admin struct {
	AdminID      string         `gorm:"column:admin_id;primaryKey" json:"admin_id"`
	AdminUsername string        `gorm:"column:admin_username;unique" json:"admin_username"`
	AdminPassword string        `gorm:"column:admin_password" json:"admin_password"`
	DeletedAt     *int64        `gorm:"column:deleted_at" json:"deleted_at,omitempty"`	
	IsDeleted     bool          `gorm:"column:is_deleted;default:false" json:"is_deleted"`	
	UpdatedAt     int64         `gorm:"column:updated_at" json:"updated_at"`	
	CreatedAt     int64         `gorm:"column:created_at" json:"created_at"`	
	AdminRole     AdminRoleEnum `gorm:"column:admin_role;default:admin" json:"admin_role"`	
	AdminEmail    string        `gorm:"column:admin_email;unique" json:"admin_email"`	
	AdminNama     string        `gorm:"column:admin_nama" json:"admin_nama"`








}	