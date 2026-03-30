package dto

import (
	"backend/models"
	"time"
)

type CreateDetailPermintaanDarahRequest struct {
	DPDID            string               `json:"detailId" binding:"required"`
	DPDPDID          string               `json:"permintaanDarahId" binding:"required"`
	DPDKomID         int                  `json:"komponenDarahId" binding:"required"`
	DPDGolonganDarah models.BloodTypeEnum `json:"golonganDarah" binding:"required,oneof=A B AB O"`
	DPDRhesus        models.RhesusEnum    `json:"rhesusDarah" binding:"required,oneof=+ -"`
	DPDJmlKantong    int                  `json:"jumlahKantong" binding:"required"`
	DPDTglDiperlukan time.Time            `json:"tanggalDiperlukan" binding:"required"`
}

type UpdateDetailPermintaanDarahRequest struct {
	DPDPDID          string               `json:"permintaanDarahId" binding:"required"`
	DPDKomID         int                  `json:"komponenDarahId" binding:"required"`
	DPDGolonganDarah models.BloodTypeEnum `json:"golonganDarah" binding:"required,oneof=A B AB O"`
	DPDRhesus        models.RhesusEnum    `json:"rhesusDarah" binding:"required,oneof=+ -"`
	DPDJmlKantong    int                  `json:"jumlahKantong" binding:"required"`
	DPDTglDiperlukan time.Time            `json:"tanggalDiperlukan" binding:"required"`
}

type DetailPermintaanDarahResponse struct {
	DPDID            string               `json:"detailId"`
	DPDPDID          string               `json:"permintaanDarahId"`
	DPDKomID         int                  `json:"komponenDarahId"`
	DPDGolonganDarah models.BloodTypeEnum `json:"golonganDarah"`
	DPDRhesus        models.RhesusEnum    `json:"rhesusDarah"`
	DPDJmlKantong    int                  `json:"jumlahKantong"`
	DPDTglDiperlukan time.Time            `json:"tanggalDiperlukan"`
	CreatedAt        time.Time            `json:"createdAt"`
}
