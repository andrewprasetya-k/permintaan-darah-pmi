package models

type Pasien struct {
	PsnID               string         `gorm:"column:psn_id;primaryKey" json:"psn_id"`
	PsnNamaPasien       string         `gorm:"column:psn_nama_pasien" json:"psn_nama_pasien"`
	PsnTempatLahir      string         `gorm:"column:psn_tempat_lahir" json:"psn_tempat_lahir"`
	PsnTglLahir         string         `gorm:"column:psn_tgl_lahir" json:"psn_tgl_lahir"`
	PsnGender           GenderEnum     `gorm:"column:psn_gender" json:"psn_gender"`
	PsnGolDarah         *BloodTypeEnum `gorm:"column:psn_gol_darah" json:"psn_gol_darah,omitempty"`
	PsnRhesus           *RhesusEnum    `gorm:"column:psn_rhesus" json:"psn_rhesus,omitempty"`
	PsnPernahTransfusi  bool           `gorm:"column:psn_pernah_transfusi" json:"psn_pernah_transfusi"`
	PsnIndikasiTransfusi string        `gorm:"column:psn_indikasi_transfusi" json:"psn_indikasi_transfusi"`
	PsnNoRM             string         `gorm:"column:psn_no_rm" json:"psn_no_rm"`
	PsnRuangBagianKelas string         `gorm:"column:psn_ruang_bagian_kelas" json:"psn_ruang_bagian_kelas"`
	PsnHemoglobin       *float64       `gorm:"column:psn_hemoglobin" json:"psn_hemoglobin,omitempty"`
	PsnPernahHamil      *string        `gorm:"column:psn_pernah_hamil" json:"psn_pernah_hamil,omitempty"`
	CreatedAt           int64          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           int64          `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted           bool           `gorm:"column:is_deleted;default:false" json:"is_deleted"`
	DeletedAt           *int64         `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}