package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type AdminService interface {
	Create(req dto.CreateAdminRequest) (*dto.AdminResponse, error)
	GetByID(id string) (*dto.AdminResponse, error)
	GetAll(limit, offset int) ([]dto.AdminResponse, error)
	Update(id string, req dto.UpdateAdminRequest) (*dto.AdminResponse, error)
	Delete(id string) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) Create(req dto.CreateAdminRequest) (*dto.AdminResponse, error) {
	data := models.Admin{
		AdminUsername: req.AdminUsername,
		AdminPassword: req.AdminPassword,
		AdminNama:     req.AdminNama,
		AdminEmail:    req.AdminEmail,
		AdminRole:     req.AdminRole,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapAdminToResponse(data)
	return &resp, nil
}

func (s *adminService) GetByID(id string) (*dto.AdminResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapAdminToResponse(*data)
	return &resp, nil
}

func (s *adminService) GetAll(limit, offset int) ([]dto.AdminResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.AdminResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapAdminToResponse(item))
	}
	return result, nil
}

func (s *adminService) Update(id string, req dto.UpdateAdminRequest) (*dto.AdminResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.AdminUsername = req.AdminUsername
	data.AdminPassword = req.AdminPassword
	data.AdminNama = req.AdminNama
	data.AdminEmail = req.AdminEmail
	data.AdminRole = req.AdminRole

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapAdminToResponse(*data)
	return &resp, nil
}

func (s *adminService) Delete(id string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func mapAdminToResponse(data models.Admin) dto.AdminResponse {
	return dto.AdminResponse{
		AdminID:       data.AdminID,
		AdminUsername: data.AdminUsername,
		AdminNama:     data.AdminNama,
		AdminEmail:    data.AdminEmail,
		AdminRole:     data.AdminRole,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
		DeletedAt:     data.DeletedAt,
	}
}
