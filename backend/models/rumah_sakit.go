package models

type RumahSakit struct {
	RSID      string  `gorm:"column:rs_id;primaryKey" json:"rs_id"`
	RSNama    string  `gorm:"column:rs_nama" json:"rs_nama"`
	RSNoTelp  string  `gorm:"column:rs_no_telp" json:"rs_no_telp"`
	RSAlamat  string  `gorm:"column:rs_alamat" json:"rs_alamat"`
	RSEmail   *string `gorm:"column:rs_email;unique" json:"rs_email,omitempty"`
	RSPassword *string `gorm:"column:rs_password" json:"rs_password,omitempty"`
	CreatedAt int64   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int64   `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted bool    `gorm:"column:is_deleted;default:false" json:"is_deleted"`
	DeletedAt *int64  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}
