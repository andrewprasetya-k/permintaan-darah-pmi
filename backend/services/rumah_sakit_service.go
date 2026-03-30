package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type RumahSakitService interface {
	Create(req dto.CreateRumahSakitRequest) (*dto.RumahSakitResponse, error)
	GetByID(id string) (*dto.RumahSakitResponse, error)
	GetAll(limit, offset int) ([]dto.RumahSakitResponse, error)
	Update(id string, req dto.UpdateRumahSakitRequest) (*dto.RumahSakitResponse, error)
	Delete(id string) error
}

type rumahSakitService struct {
	repo repository.RumahSakitRepository
}

func NewRumahSakitService(repo repository.RumahSakitRepository) RumahSakitService {
	return &rumahSakitService{repo: repo}
}

func (s *rumahSakitService) Create(req dto.CreateRumahSakitRequest) (*dto.RumahSakitResponse, error) {
	data := models.RumahSakit{
		RSID:       req.RSID,
		RSNama:     req.RSNama,
		RSNoTelp:   req.RSNoTelp,
		RSAlamat:   req.RSAlamat,
		RSEmail:    req.RSEmail,
		RSPassword: req.RSPassword,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapRumahSakitToResponse(data)
	return &resp, nil
}

func (s *rumahSakitService) GetByID(id string) (*dto.RumahSakitResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapRumahSakitToResponse(*data)
	return &resp, nil
}

func (s *rumahSakitService) GetAll(limit, offset int) ([]dto.RumahSakitResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.RumahSakitResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapRumahSakitToResponse(item))
	}
	return result, nil
}

func (s *rumahSakitService) Update(id string, req dto.UpdateRumahSakitRequest) (*dto.RumahSakitResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.RSNama = req.RSNama
	data.RSNoTelp = req.RSNoTelp
	data.RSAlamat = req.RSAlamat
	data.RSEmail = req.RSEmail
	data.RSPassword = req.RSPassword

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapRumahSakitToResponse(*data)
	return &resp, nil
}

func (s *rumahSakitService) Delete(id string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func mapRumahSakitToResponse(data models.RumahSakit) dto.RumahSakitResponse {
	return dto.RumahSakitResponse{
		RSID:       data.RSID,
		RSNama:     data.RSNama,
		RSNoTelp:   data.RSNoTelp,
		RSAlamat:   data.RSAlamat,
		RSEmail:    data.RSEmail,
		RSPassword: data.RSPassword,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt,
	}
}
