package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type PermintaanDarahService interface {
	Create(req dto.CreatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error)
	GetByID(id string) (*dto.PermintaanDarahResponse, error)
	GetAll(limit, offset int) ([]dto.PermintaanDarahResponse, error)
	Update(id string, req dto.UpdatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error)
	Delete(id string) error
}

type permintaanDarahService struct {
	repo repository.PermintaanDarahRepository
}

func NewPermintaanDarahService(repo repository.PermintaanDarahRepository) PermintaanDarahService {
	return &permintaanDarahService{repo: repo}
}

func (s *permintaanDarahService) Create(req dto.CreatePermintaanDarahRequest) (*dto.PermintaanDarahResponse, error) {
	data := models.PermintaanDarah{
		PDID:                req.PDID,
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

func (s *permintaanDarahService) GetAll(limit, offset int) ([]dto.PermintaanDarahResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
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

func mapPermintaanToResponse(data models.PermintaanDarah) dto.PermintaanDarahResponse {
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
	}
}
