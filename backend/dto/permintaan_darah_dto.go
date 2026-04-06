package dto

import (
	"backend/models"
	"time"
)

type CreatePermintaanDarahRequest struct {
	PDRsID string `json:"rumahSakitId" binding:"required"`

	// Data Pasien
	PDNamaPasien  string                `json:"namaPasien" binding:"required"`
	PDNoRM        *string               `json:"noRM,omitempty"`
	PDTempatLahir string                `json:"tempatLahir" binding:"required"`
	PDTglLahir    time.Time             `json:"tanggalLahir" binding:"required"`
	PDGender      models.GenderEnum     `json:"jenisKelamin" binding:"required,oneof=L P"`
	PDGolDarah    *models.BloodTypeEnum `json:"golonganDarah,omitempty"`
	PDRhesus      *models.RhesusEnum    `json:"rhesusDarah,omitempty"`

	// Data Medis
	PDHemoglobin        *float64 `json:"hemoglobin,omitempty"`
	PDRuangBagianKelas  *string  `json:"ruangBagianKelas,omitempty"`
	PDPernahTransfusi   bool     `json:"pernahTransfusi"`
	PDIndikasiTransfusi *string  `json:"indikasiTransfusi,omitempty"`
	PDPernahHamil       *string  `json:"pernahHamil,omitempty"`

	// Status
	PDStatus        models.PermintaanStatusEnum `json:"status" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	PDCancelReason  *string                     `json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time                   `json:"tanggalPermintaan" binding:"required"`
}

type UpdatePermintaanDarahRequest struct {
	PDRsID string `json:"rumahSakitId" binding:"required"`

	// Data Pasien
	PDNamaPasien  string                `json:"namaPasien" binding:"required"`
	PDNoRM        *string               `json:"noRM,omitempty"`
	PDTempatLahir string                `json:"tempatLahir" binding:"required"`
	PDTglLahir    time.Time             `json:"tanggalLahir" binding:"required"`
	PDGender      models.GenderEnum     `json:"jenisKelamin" binding:"required,oneof=L P"`
	PDGolDarah    *models.BloodTypeEnum `json:"golonganDarah,omitempty"`
	PDRhesus      *models.RhesusEnum    `json:"rhesusDarah,omitempty"`

	// Data Medis
	PDHemoglobin        *float64 `json:"hemoglobin,omitempty"`
	PDRuangBagianKelas  *string  `json:"ruangBagianKelas,omitempty"`
	PDPernahTransfusi   bool     `json:"pernahTransfusi"`
	PDIndikasiTransfusi *string  `json:"indikasiTransfusi,omitempty"`
	PDPernahHamil       *string  `json:"pernahHamil,omitempty"`

	// Status
	PDStatus        models.PermintaanStatusEnum `json:"status" binding:"required,oneof=dibuat diproses bisa_diambil selesai dibatalkan"`
	PDCancelReason  *string                     `json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time                   `json:"tanggalPermintaan" binding:"required"`
}

type PermintaanDarahResponse struct {
	PDID   string `json:"permintaanDarahId"`
	PDRsID string `json:"rumahSakitId"`

	// Data Pasien
	PDNamaPasien  string                `json:"namaPasien"`
	PDNoRM        *string               `json:"noRM,omitempty"`
	PDTempatLahir string                `json:"tempatLahir"`
	PDTglLahir    time.Time             `json:"tanggalLahir"`
	PDGender      models.GenderEnum     `json:"jenisKelamin"`
	PDGolDarah    *models.BloodTypeEnum `json:"golonganDarah,omitempty"`
	PDRhesus      *models.RhesusEnum    `json:"rhesusDarah,omitempty"`

	// Data Medis
	PDHemoglobin        *float64 `json:"hemoglobin,omitempty"`
	PDRuangBagianKelas  *string  `json:"ruangBagianKelas,omitempty"`
	PDPernahTransfusi   bool     `json:"pernahTransfusi"`
	PDIndikasiTransfusi *string  `json:"indikasiTransfusi,omitempty"`
	PDPernahHamil       *string  `json:"pernahHamil,omitempty"`

	// Status
	PDStatus        models.PermintaanStatusEnum `json:"status"`
	PDCancelReason  *string                     `json:"cancelReason,omitempty"`
	PDTglPermintaan time.Time                   `json:"tanggalPermintaan"`
	CreatedAt       time.Time                   `json:"createdAt"`
	UpdatedAt       time.Time                   `json:"updatedAt"`
	DeletedAt       *time.Time                  `json:"deletedAt,omitempty"`

	//detail permintaan darah
	Details []DetailPermintaanDarahResponse `json:"detailPermintaanDarah"`
}

type PermintaanDarahFilters struct {
	Status    *string
	RsID      *string
	Gender    *string
	GolDarah  *string
	StartDate *time.Time
	EndDate   *time.Time
}
