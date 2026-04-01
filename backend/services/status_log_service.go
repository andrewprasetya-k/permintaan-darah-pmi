package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type StatusLogService interface {
	Create(req dto.CreateStatusLogRequest) (*dto.StatusLogResponse, error)
	GetByID(id string) (*dto.StatusLogResponse, error)
	GetAll(limit, offset int) ([]dto.StatusLogResponse, error)
	Update(id string, req dto.UpdateStatusLogRequest) (*dto.StatusLogResponse, error)
	Delete(id string) error
}

type statusLogService struct {
	repo repository.StatusLogRepository
}

func NewStatusLogService(repo repository.StatusLogRepository) StatusLogService {
	return &statusLogService{repo: repo}
}

func (s *statusLogService) Create(req dto.CreateStatusLogRequest) (*dto.StatusLogResponse, error) {
	data := models.StatusLog{
		LogPdID:     req.LogPdID,
		LogAdminID:  req.LogAdminID,
		LogStatusTo: req.LogStatusTo,
		LogNotes:    req.LogNotes,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}
	resp := mapStatusLogToResponse(data)
	return &resp, nil
}

func (s *statusLogService) GetByID(id string) (*dto.StatusLogResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapStatusLogToResponse(*data)
	return &resp, nil
}

func (s *statusLogService) GetAll(limit, offset int) ([]dto.StatusLogResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.StatusLogResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapStatusLogToResponse(item))
	}
	return result, nil
}

func (s *statusLogService) Update(id string, req dto.UpdateStatusLogRequest) (*dto.StatusLogResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	data.LogPdID = req.LogPdID
	data.LogAdminID = req.LogAdminID
	data.LogStatusTo = req.LogStatusTo
	data.LogNotes = req.LogNotes

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	resp := mapStatusLogToResponse(*data)
	return &resp, nil
}

func (s *statusLogService) Delete(id string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(data)
}

func mapStatusLogToResponse(data models.StatusLog) dto.StatusLogResponse {
	return dto.StatusLogResponse{
		LogID:       data.LogID,
		LogPdID:     data.LogPdID,
		LogAdminID:  data.LogAdminID,
		LogStatusTo: data.LogStatusTo,
		LogNotes:    data.LogNotes,
		CreatedAt:   data.CreatedAt,
	}
}
