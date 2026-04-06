package models

import "time"

type PermintaanDarah struct {
	PDID   string `gorm:"column:pd_id;primaryKey;default:(-)" json:"permintaanDarahId"`
	PDRsID string `gorm:"column:pd_rs_id" json:"rumahSakitId"`

	// Data Pasien (Flat - Ditulis Ulang Tiap Permintaan)
	PDNamaPasien  string         `gorm:"column:pd_nama_pasien" json:"namaPasien"`
	PDNoRM        *string        `gorm:"column:pd_no_rm" json:"noRM,omitempty"`
	PDTempatLahir string         `gorm:"column:pd_tempat_lahir" json:"tempatLahir"`
	PDTglLahir    time.Time      `gorm:"column:pd_tgl_lahir;type:date" json:"tanggalLahir"`
	PDGender      GenderEnum     `gorm:"column:pd_gender" json:"jenisKelamin"`
	PDGolDarah    *BloodTypeEnum `gorm:"column:pd_gol_darah" json:"golonganDarah,omitempty"`
	PDRhesus      *RhesusEnum    `gorm:"column:pd_rhesus" json:"rhesusDarah,omitempty"`

	// Data Medis (Dinamis per Transaksi)
	PDHemoglobin        *float64 `gorm:"column:pd_hemoglobin" json:"hemoglobin,omitempty"`
	PDRuangBagianKelas  *string  `gorm:"column:pd_ruang_kelas" json:"ruangBagianKelas,omitempty"`
	PDPernahTransfusi   bool     `gorm:"column:pd_pernah_transfusi;default:false" json:"pernahTransfusi"`
	PDIndikasiTransfusi *string  `gorm:"column:pd_indikasi_transfusi" json:"indikasiTransfusi,omitempty"`
	PDPernahHamil       *string  `gorm:"column:pd_pernah_hamil" json:"pernahHamil,omitempty"`

	// Status & Audit Trail
	PDStatus        PermintaanStatusEnum `gorm:"column:pd_status;default:dibuat" json:"status"`
	PDCancelReason  *string              `gorm:"column:pd_cancel_reason" json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time            `gorm:"column:pd_tgl_permintaan" json:"tanggalPermintaan"`
	CreatedAt       time.Time            `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time            `gorm:"column:updated_at" json:"updatedAt"`
	IsDeleted       bool                 `gorm:"column:is_deleted;default:false" json:"isDeleted"`
	DeletedAt       *time.Time           `gorm:"column:deleted_at" json:"deletedAt,omitempty"`

	// Relations
	Details []DetailPermintaanDarah `gorm:"foreignKey:DPDPDID;references:PDID" json:"details"`
}

func (PermintaanDarah) TableName() string {
	return "permintaan_darah"
}
