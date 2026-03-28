package models

type KomponenDarah struct {
	KomID    int    `gorm:"column:kom_id;primaryKey;autoIncrement" json:"kom_id"`
	KomNama  string `gorm:"column:kom_nama" json:"kom_nama"`
	KomKode  string `gorm:"column:kom_kode;unique" json:"kom_kode"`
	IsActive bool   `gorm:"column:is_active;default:true" json:"is_active"`
}
