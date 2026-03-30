package models

type Pasien struct {
	PsnID               string         `gorm:"column:psn_id;primaryKey" json:"pasienId"`
	PsnNamaPasien       string         `gorm:"column:psn_nama_pasien" json:"namaPasien"`
	PsnTempatLahir      string         `gorm:"column:psn_tempat_lahir" json:"tempatLahir"`
	PsnTglLahir         string         `gorm:"column:psn_tgl_lahir" json:"tanggalLahir"`
	PsnGender           GenderEnum     `gorm:"column:psn_gender" json:"jenisKelamin"`
	PsnGolDarah         *BloodTypeEnum `gorm:"column:psn_gol_darah" json:"golonganDarahPasien,omitempty"`
	PsnRhesus           *RhesusEnum    `gorm:"column:psn_rhesus" json:"rhesusDarahPasien,omitempty"`
	PsnPernahTransfusi  bool           `gorm:"column:psn_pernah_transfusi" json:"pernahTransfusi"`
	PsnIndikasiTransfusi string        `gorm:"column:psn_indikasi_transfusi" json:"indikasiTransfusi"`
	PsnNoRM             string         `gorm:"column:psn_no_rm" json:"noRM"`
	PsnRuangBagianKelas string         `gorm:"column:psn_ruang_bagian_kelas" json:"ruangBagianKelas"`
	PsnHemoglobin       *float64       `gorm:"column:psn_hemoglobin" json:"hemoglobinPasien,omitempty"`
	PsnPernahHamil      *string        `gorm:"column:psn_pernah_hamil" json:"pernahHamil,omitempty"`
	CreatedAt           int64          `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt           int64          `gorm:"column:updated_at" json:"updatedAt"`
	IsDeleted           bool           `gorm:"column:is_deleted;default:false" json:"isDeleted"`
	DeletedAt           *int64         `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}