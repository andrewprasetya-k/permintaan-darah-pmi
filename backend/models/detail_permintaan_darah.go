package models

type DetailPermintaanDarah struct {
	DPDID            string         `gorm:"column:dpd_id;primaryKey" json:"dpd_id"`
	DPDPdID          string         `gorm:"column:dpd_pd_id" json:"dpd_pd_id"`
	DPDKomID         int            `gorm:"column:dpd_kom_id" json:"dpd_kom_id"`
	DPDGolonganDarah BloodTypeEnum  `gorm:"column:dpd_golongan_darah" json:"dpd_golongan_darah"`
	DPDRhesus        RhesusEnum     `gorm:"column:dpd_rhesus" json:"dpd_rhesus"`
	DPDJmlKantong    int            `gorm:"column:dpd_jml_kantong" json:"dpd_jml_kantong"`
	DPDTglDiperlukan int64          `gorm:"column:dpd_tgl_diperlukan" json:"dpd_tgl_diperlukan"`
	CreatedAt        int64          `gorm:"column:created_at" json:"created_at"`
}
