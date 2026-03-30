package models

type KomponenDarah struct {
	KomID    int    `gorm:"column:kom_id;primaryKey;autoIncrement" json:"komponenId"`
	KomNama  string `gorm:"column:kom_nama" json:"komponenDarah"`
	KomKode  string `gorm:"column:kom_kode;unique" json:"komponenKode"`
	IsActive bool   `gorm:"column:is_active;default:true" json:"isActive"`
}

func (KomponenDarah) TableName() string {
	return "komponen_darah"
}
