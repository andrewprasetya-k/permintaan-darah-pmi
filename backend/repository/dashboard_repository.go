package repository

import (
	"backend/dto"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	StatusSummary(rumahSakitID string) (*dto.StatusSummaryResponse, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) StatusSummary(rumahSakitID string) (*dto.StatusSummaryResponse, error) {
	type statusCountRow struct {
		Status string `gorm:"column:status"`
		Count  int    `gorm:"column:count"`
	}

	query := r.db.Table("permintaan_darah").
		Select("pd_status as status, COUNT(*) as count").
		Where("is_deleted = ?", false)

	if rumahSakitID != "" {
		query = query.Where("pd_rs_id = ?", rumahSakitID)
	}

	var rows []statusCountRow
	if err := query.Group("pd_status").Scan(&rows).Error; err != nil {
		return nil, err
	}

	summary := &dto.StatusSummaryResponse{}
	for _, row := range rows {
		switch row.Status {
		case "dibuat":
			summary.Dibuat = row.Count
		case "diproses":
			summary.Diproses = row.Count
		case "bisa_diambil":
			summary.BisaDiambil = row.Count
		case "selesai":
			summary.Selesai = row.Count
		case "dibatalkan":
			summary.Dibatalkan = row.Count
		}
		summary.Total += row.Count
	}

	return summary, nil
}
