package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
)

type RumahSakitService interface {
	Create(req dto.CreateRumahSakitRequest) (*dto.RumahSakitResponse, error)
	GetByID(id string) (*dto.RumahSakitResponse, error)
	GetAll(limit, offset int) ([]dto.RumahSakitResponse, error)
	Update(id string, req dto.UpdateRumahSakitRequest) (*dto.RumahSakitResponse, error)
	Delete(id string) error
	Restore(id string) error

	//filter purpose
	GetDistinctRSNama() ([]dto.RumahSakitDistinctNamaResponse, error)
}

type rumahSakitService struct {
	repo repository.RumahSakitRepository
}

func NewRumahSakitService(repo repository.RumahSakitRepository) RumahSakitService {
	return &rumahSakitService{repo: repo}
}

func (s *rumahSakitService) Create(req dto.CreateRumahSakitRequest) (*dto.RumahSakitResponse, error) {
	hashedPassword, err := utils.HashPassword(req.RSPassword)
	if err != nil {
		return nil, err
	}
	data := models.RumahSakit{
		RSNama:     req.RSNama,
		RSNoTelp:   req.RSNoTelp,
		RSAlamat:   req.RSAlamat,
		RSEmail:    req.RSEmail,
		RSUsername: req.RSUsername,
		RSPassword: hashedPassword,
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

func (s *rumahSakitService) GetDistinctRSNama() ([]dto.RumahSakitDistinctNamaResponse, error) {
	return s.repo.GetDistinctRSNama()
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
	data.RSPassword = *req.RSPassword

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapRumahSakitToResponse(*data)
	return &resp, nil
}

func (s *rumahSakitService) Delete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *rumahSakitService) Restore(id string) error {
	return s.repo.Restore(id)
}

func mapRumahSakitToResponse(data models.RumahSakit) dto.RumahSakitResponse {
	return dto.RumahSakitResponse{
		RSID:       data.RSID,
		RSNama:     data.RSNama,
		RSNoTelp:   data.RSNoTelp,
		RSAlamat:   data.RSAlamat,
		RSEmail:    data.RSEmail,
		RSUsername: data.RSUsername,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt,
	}
}
