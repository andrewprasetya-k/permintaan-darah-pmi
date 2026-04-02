package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"errors"
)

type PermintaanDarahService interface {
	Create(req dto.CreatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error)
	GetByID(id string) (*dto.PermintaanDarahResponse, error)
	GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]dto.PermintaanDarahResponse, error)
	GetByRsID(rsID string, limit, offset int) ([]dto.PermintaanDarahResponse, error)
	Update(id string, req dto.UpdatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error)
	Delete(id string) error

	UpdateStatus(pdID string, newStatus models.PermintaanStatusEnum, reason *string) (*dto.PermintaanDarahResponse, error)
}

type permintaanDarahService struct {
	repo repository.PermintaanDarahRepository
}

func NewPermintaanDarahService(repo repository.PermintaanDarahRepository) PermintaanDarahService {
	return &permintaanDarahService{repo: repo}
}

func (s *permintaanDarahService) Create(req dto.CreatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error) {
	data := models.PermintaanDarah{
		PDRsID:              req.PDRsID,
		PDNamaPasien:        req.PDNamaPasien,
		PDNoRM:              req.PDNoRM,
		PDTempatLahir:       req.PDTempatLahir,
		PDTglLahir:          req.PDTglLahir,
		PDGender:            req.PDGender,
		PDGolDarah:          req.PDGolDarah,
		PDRhesus:            req.PDRhesus,
		PDHemoglobin:        req.PDHemoglobin,
		PDRuangBagianKelas:  req.PDRuangBagianKelas,
		PDPernahTransfusi:   req.PDPernahTransfusi,
		PDIndikasiTransfusi: req.PDIndikasiTransfusi,
		PDPernahHamil:       req.PDPernahHamil,
		PDStatus:            req.PDStatus,
		PDCancelReason:      req.PDCancelReason,
		PDTglPermintaan:     req.PDTglPermintaan,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapPermintaanToResponse(data)
	return &resp, nil
}

func (s *permintaanDarahService) GetByID(id string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapPermintaanToResponse(*data)
	return &resp, nil
}

func (s *permintaanDarahService) GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]dto.PermintaanDarahResponse, error) {
	list, err := s.repo.GetAll(filters, limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.PermintaanDarahResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapPermintaanToResponse(item))
	}
	return result, nil
}

func (s *permintaanDarahService) GetByRsID(rsID string, limit, offset int) ([]dto.PermintaanDarahResponse, error) {
	list, err := s.repo.GetByRsID(rsID, limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.PermintaanDarahResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapPermintaanToResponse(item))
	}
	return result, nil
}

func (s *permintaanDarahService) Update(id string, req dto.UpdatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.PDRsID = req.PDRsID
	data.PDNamaPasien = req.PDNamaPasien
	data.PDNoRM = req.PDNoRM
	data.PDTempatLahir = req.PDTempatLahir
	data.PDTglLahir = req.PDTglLahir
	data.PDGender = req.PDGender
	data.PDGolDarah = req.PDGolDarah
	data.PDRhesus = req.PDRhesus
	data.PDHemoglobin = req.PDHemoglobin
	data.PDRuangBagianKelas = req.PDRuangBagianKelas
	data.PDPernahTransfusi = req.PDPernahTransfusi
	data.PDIndikasiTransfusi = req.PDIndikasiTransfusi
	data.PDPernahHamil = req.PDPernahHamil
	data.PDStatus = req.PDStatus
	data.PDCancelReason = req.PDCancelReason
	data.PDTglPermintaan = req.PDTglPermintaan

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapPermintaanToResponse(*data)
	return &resp, nil
}

func (s *permintaanDarahService) Delete(id string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func (s *permintaanDarahService) UpdateStatus(pdID string, newStatus models.PermintaanStatusEnum, reason *string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(pdID)
	if err != nil {
		return nil, err
	}

	if data.PDStatus == "selesai" || data.PDStatus == "dibatalkan" {
		return nil, errors.New("cannot update status of a completed or cancelled request")
	}

	data.PDStatus = newStatus
	if newStatus == "dibatalkan" && reason == nil {
		return nil, errors.New("reason is required")
	} else if newStatus == "dibatalkan" && reason != nil {
		data.PDCancelReason = reason
	}
	
	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	resp := mapPermintaanToResponse(*data)
	return &resp, nil
}

func mapPermintaanToResponse(data models.PermintaanDarah) dto.PermintaanDarahResponse {
	details := make([]dto.DetailPermintaanDarahResponse, 0, len(data.Details))
	for _, d := range data.Details {
		details = append(details, dto.DetailPermintaanDarahResponse{
			DPDID:            d.DPDID,
			DPDPDID:          d.DPDPDID,
			DPDKomID:         d.DPDKomID,
			DPDGolonganDarah: d.DPDGolonganDarah,
			DPDRhesus:        d.DPDRhesus,
			DPDJmlKantong:    d.DPDJmlKantong,
			DPDTglDiperlukan: d.DPDTglDiperlukan,
			CreatedAt:        d.CreatedAt,
		})
	}

	return dto.PermintaanDarahResponse{
		PDID:                data.PDID,
		PDRsID:              data.PDRsID,
		PDNamaPasien:        data.PDNamaPasien,
		PDNoRM:              data.PDNoRM,
		PDTempatLahir:       data.PDTempatLahir,
		PDTglLahir:          data.PDTglLahir,
		PDGender:            data.PDGender,
		PDGolDarah:          data.PDGolDarah,
		PDRhesus:            data.PDRhesus,
		PDHemoglobin:        data.PDHemoglobin,
		PDRuangBagianKelas:  data.PDRuangBagianKelas,
		PDPernahTransfusi:   data.PDPernahTransfusi,
		PDIndikasiTransfusi: data.PDIndikasiTransfusi,
		PDPernahHamil:       data.PDPernahHamil,
		PDStatus:            data.PDStatus,
		PDCancelReason:      data.PDCancelReason,
		PDTglPermintaan:     data.PDTglPermintaan,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
		DeletedAt:           data.DeletedAt,
		Details:              details,
	}
}
