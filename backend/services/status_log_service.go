package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type StatusLogService interface {
	Create(req dto.CreateStatusLogRequest) (*dto.StatusLogResponse, error)
	GetByID(id string) (*dto.StatusLogResponse, error)
	GetAll(limit, offset int) ([]dto.StatusLogResponse, int, error)
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

func (s *statusLogService) GetAll(limit, offset int) ([]dto.StatusLogResponse, int, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.StatusLogResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapStatusLogToResponse(item))
	}
	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return result, int(total), nil
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
