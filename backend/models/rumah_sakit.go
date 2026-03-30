package models

import "time"

type RumahSakit struct {
	RSID       string     `gorm:"column:rs_id;primaryKey" json:"rumahSakitId"`
	RSNama     string     `gorm:"column:rs_nama" json:"nama"`
	RSNoTelp   string     `gorm:"column:rs_no_telp" json:"nomorTelepon"`
	RSAlamat   string     `gorm:"column:rs_alamat" json:"alamat"`
	RSEmail    *string    `gorm:"column:rs_email;unique" json:"email,omitempty"`
	RSPassword *string    `gorm:"column:rs_password" json:"password,omitempty"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}
