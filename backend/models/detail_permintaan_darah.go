package models

type DetailPermintaanDarah struct {
	DPDID            string         `gorm:"column:dpd_id;primaryKey" json:"detailId"`
	DPDGolonganDarah BloodTypeEnum  `gorm:"column:dpd_golongan_darah" json:"golonganDarah"`
	DPDRhesus        RhesusEnum     `gorm:"column:dpd_rhesus" json:"rhesusDarah"`
	DPDJmlKantong    int            `gorm:"column:dpd_jml_kantong" json:"jumlahKantong"`
	DPDTglDiperlukan int64          `gorm:"column:dpd_tgl_diperlukan" json:"tanggalDiperlukan"`
	CreatedAt        int64          `gorm:"column:created_at" json:"createdAt"`
}
