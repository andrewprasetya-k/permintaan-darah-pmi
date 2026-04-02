package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type DetailPermintaanDarahService interface {
	Create(req dto.CreateDetailPermintaanDarahRequest) (*dto.DetailPermintaanDarahResponse, error)
	GetByID(id int) (*dto.DetailPermintaanDarahResponse, error)
	GetAll(limit, offset int) ([]dto.DetailPermintaanDarahResponse, error)
	Update(id int, req dto.UpdateDetailPermintaanDarahRequest) (*dto.DetailPermintaanDarahResponse, error)
	Delete(id int) error
}

type detailPermintaanDarahService struct {
	repo repository.DetailPermintaanDarahRepository
}

func NewDetailPermintaanDarahService(repo repository.DetailPermintaanDarahRepository) DetailPermintaanDarahService {
	return &detailPermintaanDarahService{repo: repo}
}

func (s *detailPermintaanDarahService) Create(req dto.CreateDetailPermintaanDarahRequest) (*dto.DetailPermintaanDarahResponse, error) {
	data := models.DetailPermintaanDarah{
		DPDPDID:          req.DPDPDID,
		DPDKomID:         req.DPDKomID,
		DPDGolonganDarah: req.DPDGolonganDarah,
		DPDRhesus:        req.DPDRhesus,
		DPDJmlKantong:    req.DPDJmlKantong,
		DPDTglDiperlukan: req.DPDTglDiperlukan,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapDetailPermintaanToResponse(data)
	return &resp, nil
}

func (s *detailPermintaanDarahService) GetByID(id int) (*dto.DetailPermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapDetailPermintaanToResponse(*data)
	return &resp, nil
}

func (s *detailPermintaanDarahService) GetAll(limit, offset int) ([]dto.DetailPermintaanDarahResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.DetailPermintaanDarahResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapDetailPermintaanToResponse(item))
	}
	return result, nil
}

func (s *detailPermintaanDarahService) Update(id int, req dto.UpdateDetailPermintaanDarahRequest) (*dto.DetailPermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.DPDPDID = req.DPDPDID
	data.DPDKomID = req.DPDKomID
	data.DPDGolonganDarah = req.DPDGolonganDarah
	data.DPDRhesus = req.DPDRhesus
	data.DPDJmlKantong = req.DPDJmlKantong
	data.DPDTglDiperlukan = req.DPDTglDiperlukan

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapDetailPermintaanToResponse(*data)
	return &resp, nil
}

func (s *detailPermintaanDarahService) Delete(id int) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func mapDetailPermintaanToResponse(data models.DetailPermintaanDarah) dto.DetailPermintaanDarahResponse {
	return dto.DetailPermintaanDarahResponse{
		DPDID:            data.DPDID,
		DPDPDID:          data.DPDPDID,
		DPDKomID:         data.DPDKomID,
		DPDGolonganDarah: data.DPDGolonganDarah,
		DPDRhesus:        data.DPDRhesus,
		DPDJmlKantong:    data.DPDJmlKantong,
		DPDTglDiperlukan: data.DPDTglDiperlukan,
		CreatedAt:        data.CreatedAt,
	}
}
