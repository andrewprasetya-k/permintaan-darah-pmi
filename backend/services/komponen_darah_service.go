package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type KomponenDarahService interface {
	Create(req dto.CreateKomponenDarahRequest) (*dto.KomponenDarahResponse, error)
	GetByID(id int) (*dto.KomponenDarahResponse, error)
	GetAll(limit, offset int) ([]dto.KomponenDarahResponse, error)
	Update(id int, req dto.UpdateKomponenDarahRequest) (*dto.KomponenDarahResponse, error)
	Delete(id int) error
	ActivateKomponenDarah(id int) (*dto.KomponenDarahResponse, error)
	DeactivateKomponenDarah(id int) (*dto.KomponenDarahResponse, error)
}

type komponenDarahService struct {
	repo                   repository.KomponenDarahRepository
	systemAccessLogService SystemAccessLogService
}

func NewKomponenDarahService(repo repository.KomponenDarahRepository, systemAccessLogService SystemAccessLogService) KomponenDarahService {
	return &komponenDarahService{repo: repo, systemAccessLogService: systemAccessLogService}
}

func (s *komponenDarahService) Create(req dto.CreateKomponenDarahRequest) (*dto.KomponenDarahResponse, error) {
	data := models.KomponenDarah{
		KomNama:  req.KomNama,
		KomKode:  req.KomKode,
		IsActive: req.IsActive,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapKomponenToResponse(data)
	return &resp, nil
}

func (s *komponenDarahService) GetByID(id int) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func (s *komponenDarahService) GetAll(limit, offset int) ([]dto.KomponenDarahResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.KomponenDarahResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapKomponenToResponse(item))
	}
	return result, nil
}

func (s *komponenDarahService) Update(id int, req dto.UpdateKomponenDarahRequest) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.KomNama = req.KomNama
	data.KomKode = req.KomKode
	data.IsActive = req.IsActive

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func (s *komponenDarahService) Delete(id int) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func (s *komponenDarahService) ActivateKomponenDarah(id int) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.ActivateKomponenDarah(id)
	if err != nil {
		return nil, err
	}
	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func (s *komponenDarahService) DeactivateKomponenDarah(id int) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.DeactivateKomponenDarah(id)
	if err != nil {
		return nil, err
	}
	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func mapKomponenToResponse(data models.KomponenDarah) dto.KomponenDarahResponse {
	return dto.KomponenDarahResponse{
		KomID:    data.KomID,
		KomNama:  data.KomNama,
		KomKode:  data.KomKode,
		IsActive: data.IsActive,
	}
}
