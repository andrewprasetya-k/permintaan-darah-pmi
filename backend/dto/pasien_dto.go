package dto

import (
	"backend/models"
	"time"
)

type CreatePasienRequest struct {
	PsnID                string                `json:"pasienId" binding:"required"`
	PsnNamaPasien        string                `json:"namaPasien" binding:"required"`
	PsnTempatLahir       string                `json:"tempatLahir" binding:"required"`
	PsnTglLahir          time.Time             `json:"tanggalLahir" binding:"required"`
	PsnGender            models.GenderEnum     `json:"jenisKelamin" binding:"required,oneof=L P"`
	PsnGolDarah          *models.BloodTypeEnum `json:"golonganDarah,omitempty" binding:"omitempty,oneof=A B AB O"`
	PsnRhesus            *models.RhesusEnum    `json:"rhesusDarah,omitempty" binding:"omitempty,oneof=+ -"`
	PsnPernahTransfusi   bool                  `json:"pernahTransfusi"`
	PsnIndikasiTransfusi string                `json:"indikasiTransfusi" binding:"required"`
	PsnNoRM              string                `json:"noRM" binding:"required"`
	PsnRuangBagianKelas  string                `json:"ruangBagianKelas" binding:"required"`
	PsnHemoglobin        *float64              `json:"hemoglobinPasien,omitempty"`
	PsnPernahHamil       *string               `json:"pernahHamil,omitempty"`
}

type UpdatePasienRequest struct {
	PsnNamaPasien        string                `json:"namaPasien" binding:"required"`
	PsnTempatLahir       string                `json:"tempatLahir" binding:"required"`
	PsnTglLahir          time.Time             `json:"tanggalLahir" binding:"required"`
	PsnGender            models.GenderEnum     `json:"jenisKelamin" binding:"required,oneof=L P"`
	PsnGolDarah          *models.BloodTypeEnum `json:"golonganDarah,omitempty" binding:"omitempty,oneof=A B AB O"`
	PsnRhesus            *models.RhesusEnum    `json:"rhesusDarah,omitempty" binding:"omitempty,oneof=+ -"`
	PsnPernahTransfusi   bool                  `json:"pernahTransfusi"`
	PsnIndikasiTransfusi string                `json:"indikasiTransfusi" binding:"required"`
	PsnNoRM              string                `json:"noRM" binding:"required"`
	PsnRuangBagianKelas  string                `json:"ruangBagianKelas" binding:"required"`
	PsnHemoglobin        *float64              `json:"hemoglobinPasien,omitempty"`
	PsnPernahHamil       *string               `json:"pernahHamil,omitempty"`
}

type PasienResponse struct {
	PsnID                string                `json:"pasienId"`
	PsnNamaPasien        string                `json:"namaPasien"`
	PsnTempatLahir       string                `json:"tempatLahir"`
	PsnTglLahir          time.Time             `json:"tanggalLahir"`
	PsnGender            models.GenderEnum     `json:"jenisKelamin"`
	PsnGolDarah          *models.BloodTypeEnum `json:"golonganDarah,omitempty"`
	PsnRhesus            *models.RhesusEnum    `json:"rhesusDarah,omitempty"`
	PsnPernahTransfusi   bool                  `json:"pernahTransfusi"`
	PsnIndikasiTransfusi string                `json:"indikasiTransfusi"`
	PsnNoRM              string                `json:"noRM"`
	PsnRuangBagianKelas  string                `json:"ruangBagianKelas"`
	PsnHemoglobin        *float64              `json:"hemoglobinPasien,omitempty"`
	PsnPernahHamil       *string               `json:"pernahHamil,omitempty"`
	CreatedAt            time.Time             `json:"createdAt"`
	UpdatedAt            time.Time             `json:"updatedAt"`
	DeletedAt            *time.Time            `json:"deletedAt,omitempty"`
}
