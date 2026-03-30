package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type PasienService interface {
	Create(req dto.CreatePasienRequest) (*dto.PasienResponse, error)
	GetByID(id string) (*dto.PasienResponse, error)
	GetAll(limit, offset int) ([]dto.PasienResponse, error)
	Update(id string, req dto.UpdatePasienRequest) (*dto.PasienResponse, error)
	Delete(id string) error
}

type pasienService struct {
	repo repository.PasienRepository
}

func NewPasienService(repo repository.PasienRepository) PasienService {
	return &pasienService{repo: repo}
}

func (s *pasienService) Create(req dto.CreatePasienRequest) (*dto.PasienResponse, error) {
	data := models.Pasien{
		PsnID:                req.PsnID,
		PsnNamaPasien:        req.PsnNamaPasien,
		PsnTempatLahir:       req.PsnTempatLahir,
		PsnTglLahir:          req.PsnTglLahir,
		PsnGender:            req.PsnGender,
		PsnGolDarah:          req.PsnGolDarah,
		PsnRhesus:            req.PsnRhesus,
		PsnPernahTransfusi:   req.PsnPernahTransfusi,
		PsnIndikasiTransfusi: req.PsnIndikasiTransfusi,
		PsnNoRM:              req.PsnNoRM,
		PsnRuangBagianKelas:  req.PsnRuangBagianKelas,
		PsnHemoglobin:        req.PsnHemoglobin,
		PsnPernahHamil:       req.PsnPernahHamil,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapPasienToResponse(data)
	return &resp, nil
}

func (s *pasienService) GetByID(id string) (*dto.PasienResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapPasienToResponse(*data)
	return &resp, nil
}

func (s *pasienService) GetAll(limit, offset int) ([]dto.PasienResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.PasienResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapPasienToResponse(item))
	}
	return result, nil
}

func (s *pasienService) Update(id string, req dto.UpdatePasienRequest) (*dto.PasienResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.PsnNamaPasien = req.PsnNamaPasien
	data.PsnTempatLahir = req.PsnTempatLahir
	data.PsnTglLahir = req.PsnTglLahir
	data.PsnGender = req.PsnGender
	data.PsnGolDarah = req.PsnGolDarah
	data.PsnRhesus = req.PsnRhesus
	data.PsnPernahTransfusi = req.PsnPernahTransfusi
	data.PsnIndikasiTransfusi = req.PsnIndikasiTransfusi
	data.PsnNoRM = req.PsnNoRM
	data.PsnRuangBagianKelas = req.PsnRuangBagianKelas
	data.PsnHemoglobin = req.PsnHemoglobin
	data.PsnPernahHamil = req.PsnPernahHamil

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapPasienToResponse(*data)
	return &resp, nil
}

func (s *pasienService) Delete(id string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func mapPasienToResponse(data models.Pasien) dto.PasienResponse {
	return dto.PasienResponse{
		PsnID:                data.PsnID,
		PsnNamaPasien:        data.PsnNamaPasien,
		PsnTempatLahir:       data.PsnTempatLahir,
		PsnTglLahir:          data.PsnTglLahir,
		PsnGender:            data.PsnGender,
		PsnGolDarah:          data.PsnGolDarah,
		PsnRhesus:            data.PsnRhesus,
		PsnPernahTransfusi:   data.PsnPernahTransfusi,
		PsnIndikasiTransfusi: data.PsnIndikasiTransfusi,
		PsnNoRM:              data.PsnNoRM,
		PsnRuangBagianKelas:  data.PsnRuangBagianKelas,
		PsnHemoglobin:        data.PsnHemoglobin,
		PsnPernahHamil:       data.PsnPernahHamil,
		CreatedAt:            data.CreatedAt,
		UpdatedAt:            data.UpdatedAt,
		DeletedAt:            data.DeletedAt,
	}
}
