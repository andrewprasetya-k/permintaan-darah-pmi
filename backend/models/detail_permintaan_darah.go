package models

import "time"

type DetailPermintaanDarah struct {
	DPDID            string         `gorm:"column:dpd_id;primaryKey" json:"detailId"`
	DPDPDID          string         `gorm:"column:dpd_pd_id" json:"permintaanDarahId"`
	DPDKomID         int            `gorm:"column:dpd_kom_id" json:"komponenDarahId"`
	DPDGolonganDarah BloodTypeEnum  `gorm:"column:dpd_golongan_darah" json:"golonganDarah"`
	DPDRhesus        RhesusEnum     `gorm:"column:dpd_rhesus" json:"rhesusDarah"`
	DPDJmlKantong    int            `gorm:"column:dpd_jml_kantong" json:"jumlahKantong"`
	DPDTglDiperlukan time.Time      `gorm:"column:dpd_tgl_diperlukan" json:"tanggalDiperlukan"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"createdAt"`
}
