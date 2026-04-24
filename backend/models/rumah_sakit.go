package models

import "time"

type RumahSakit struct {
	RSID       string     `gorm:"column:rs_id;primaryKey;default:(-)" json:"rumahSakitId"`
	RSNama     string     `gorm:"column:rs_nama" json:"nama"`
	RSNoTelp   string     `gorm:"column:rs_no_telp" json:"nomorTelepon"`
	RSAlamat   string     `gorm:"column:rs_alamat" json:"alamat"`
	RSEmail    *string    `gorm:"column:rs_email;unique" json:"email,omitempty"`
	RSUsername string     `gorm:"column:rs_username;unique" json:"username"`
	RSPassword string     `gorm:"column:rs_password" json:"password"`
	IsDeleted  bool       `gorm:"column:is_deleted;default:false" json:"isDeleted"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}

func (RumahSakit) TableName() string {
	return "rumah_sakit"
}
