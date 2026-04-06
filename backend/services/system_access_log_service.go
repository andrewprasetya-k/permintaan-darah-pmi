package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type SystemAccessLogService interface {
	Create(req dto.CreateSystemAccessLogRequest) (*dto.SystemAccessLogResponse, error)
	GetByID(logID int64) (*dto.SystemAccessLogResponse, error)
	GetAll(limit, offset int) (*dto.SystemAccessLogListResponse, error)
	
	GetByUserID(userID string, limit, offset int) (*dto.SystemAccessLogListResponse, error)
	GetByAction(action string, limit, offset int) (*dto.SystemAccessLogListResponse, error)
	GetByTargetTable(targetTable string, limit, offset int) (*dto.SystemAccessLogListResponse, error)
	GetByTargetID(targetID string, limit, offset int) (*dto.SystemAccessLogListResponse, error)
	
	LogAction(
		userID *string,
		userName string,
		userRole string,
		action string,
		targetTable *string,
		targetID *string,
		notes string,
		userAgent *string,
	) (*dto.SystemAccessLogResponse, error)
}

type systemAccessLogService struct {
	repo repository.SystemAccessLogRepository
}

func NewSystemAccessLogService(repo repository.SystemAccessLogRepository) SystemAccessLogService {
	return &systemAccessLogService{repo: repo}
}

func (s *systemAccessLogService) Create(req dto.CreateSystemAccessLogRequest) (*dto.SystemAccessLogResponse, error) {
	log := models.SystemAccessLog{
		SysUserID:      req.SysUserID,
		SysUserNama:    req.SysUserNama,
		SysUserRole:    req.SysUserRole,
		SysAction:      req.SysAction,
		SysTargetTable: req.SysTargetTable,
		SysTargetID:    req.SysTargetID,
		SysNotes:       req.SysNotes,
		SysUserAgent:   req.SysUserAgent,
	}

	if err := s.repo.Create(&log); err != nil {
		return nil, err
	}

	return mapToResponse(&log), nil
}

func (s *systemAccessLogService) GetByID(logID int64) (*dto.SystemAccessLogResponse, error) {
	log, err := s.repo.GetByID(logID)
	if err != nil {
		return nil, err
	}

	return mapToResponse(log), nil
}

func (s *systemAccessLogService) GetAll(limit, offset int) (*dto.SystemAccessLogListResponse, error) {
	logs, total, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.SystemAccessLogListResponse{
		Data:     logsToResponses(logs),
		Total:    total,
		Page:     offset/limit + 1,
		PageSize: limit,
	}, nil
}

func (s *systemAccessLogService) GetByUserID(userID string, limit, offset int) (*dto.SystemAccessLogListResponse, error) {
	logs, total, err := s.repo.GetByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.SystemAccessLogListResponse{
		Data:     logsToResponses(logs),
		Total:    total,
		Page:     offset/limit + 1,
		PageSize: limit,
	}, nil
}

func (s *systemAccessLogService) GetByAction(action string, limit, offset int) (*dto.SystemAccessLogListResponse, error) {
	logs, total, err := s.repo.GetByAction(action, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.SystemAccessLogListResponse{
		Data:     logsToResponses(logs),
		Total:    total,
		Page:     offset/limit + 1,
		PageSize: limit,
	}, nil
}

func (s *systemAccessLogService) GetByTargetTable(targetTable string, limit, offset int) (*dto.SystemAccessLogListResponse, error) {
	logs, total, err := s.repo.GetByTargetTable(targetTable, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.SystemAccessLogListResponse{
		Data:     logsToResponses(logs),
		Total:    total,
		Page:     offset/limit + 1,
		PageSize: limit,
	}, nil
}

func (s *systemAccessLogService) GetByTargetID(targetID string, limit, offset int) (*dto.SystemAccessLogListResponse, error) {
	logs, total, err := s.repo.GetByTargetID(targetID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.SystemAccessLogListResponse{
		Data:     logsToResponses(logs),
		Total:    total,
		Page:     offset/limit + 1,
		PageSize: limit,
	}, nil
}

func (s *systemAccessLogService) LogAction(
	userID *string,
	userName string,
	userRole string,
	action string,
	targetTable *string,
	targetID *string,
	notes string,
	userAgent *string,
) (*dto.SystemAccessLogResponse, error) {
	req := dto.CreateSystemAccessLogRequest{
		SysUserID:      userID,
		SysUserNama:    userName,
		SysUserRole:    userRole,
		SysAction:      action,
		SysTargetTable: targetTable,
		SysTargetID:    targetID,
		SysNotes:       notes,
		SysUserAgent:   userAgent,
	}

	return s.Create(req)
}

func mapToResponse(log *models.SystemAccessLog) *dto.SystemAccessLogResponse {
	return &dto.SystemAccessLogResponse{
		SysLogID:       log.SysLogID,
		SysUserID:      log.SysUserID,
		SysUserNama:    log.SysUserNama,
		SysUserRole:    log.SysUserRole,
		SysAction:      log.SysAction,
		SysTargetTable: log.SysTargetTable,
		SysTargetID:    log.SysTargetID,
		SysNotes:       log.SysNotes,
		SysUserAgent:   log.SysUserAgent,
		CreatedAt:      log.CreatedAt,
	}
}

func logsToResponses(logs []models.SystemAccessLog) []dto.SystemAccessLogResponse {
	responses := make([]dto.SystemAccessLogResponse, 0, len(logs))
	for _, log := range logs {
		responses = append(responses, *mapToResponse(&log))
	}
	return responses
}
