package models

import "time"

type RumahSakit struct {
	RSID      string  `gorm:"column:rs_id;primaryKey" json:"rs_id"`
	RSNama    string  `gorm:"column:rs_nama" json:"rs_nama"`
	RSNoTelp  string  `gorm:"column:rs_no_telp" json:"rs_no_telp"`
	RSAlamat  string  `gorm:"column:rs_alamat" json:"rs_alamat"`
	RSEmail   *string `gorm:"column:rs_email;unique" json:"rs_email,omitempty"`
	RSPassword *string `gorm:"column:rs_password" json:"rs_password,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}
