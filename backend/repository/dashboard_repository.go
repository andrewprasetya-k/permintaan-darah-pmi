package repository

import (
	"gorm.io/gorm"
)

type DashboardRepository interface {
	StatusSummary(rumahSakitID string) ([]int, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) StatusSummary(rumahSakitID string) ([]int, error) {
	status := []int{0, 0, 0, 0, 0}
	err := r.db.Table("permintaan_darah").
		Select("pd_status as status, COUNT(*) as count").
		Where("pd_rs_id = ?", rumahSakitID).
		Group("pd_status").
		Scan(&status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}
