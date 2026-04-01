package services

import "backend/repository"

type DashboardService interface {
	StatusSummary(rumahSakitID string) ([]int, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService() DashboardService {
	return &dashboardService{}
}

func (s *dashboardService) StatusSummary(rumahSakitID string) ([]int, error) {
	status := []int{0, 0, 0, 0, 0}
	sum,err := s.repo.StatusSummary(rumahSakitID)
	if err != nil {
		status = []int{0, 0, 0, 0, 0}
		return status, err
	}
	for _, count := range sum {
		for i := 0; i < len(status); i++ {
			status[i] += count
		}
	}
	return status, err
}