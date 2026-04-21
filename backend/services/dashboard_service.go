package services

import (
	"backend/dto"
	"backend/repository"
)

type DashboardService interface {
	StatusSummary(rumahSakitID string) (*dto.StatusSummaryResponse, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(repo repository.DashboardRepository) DashboardService {
	return &dashboardService{repo: repo}
}

func (s *dashboardService) StatusSummary(rumahSakitID string) (*dto.StatusSummaryResponse, error) {
	return s.repo.StatusSummary(rumahSakitID)
}
