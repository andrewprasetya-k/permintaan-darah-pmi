package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type StatusLogRepository interface {
	Create(data *models.StatusLog) error
	GetByID(logID string) (*models.StatusLog, error)
	Count() (int64, error)
	GetAll(limit, offset int) ([]models.StatusLog, error)
}

type statusLogRepository struct {
	db *gorm.DB
}

func NewStatusLogRepository(db *gorm.DB) StatusLogRepository {
	return &statusLogRepository{db: db}
}

func (r *statusLogRepository) Create(data *models.StatusLog) error {
	return r.db.Create(data).Error
}

func (r *statusLogRepository) GetByID(logID string) (*models.StatusLog, error) {
	var data models.StatusLog
	err := r.db.First(&data, "log_id = ?", logID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *statusLogRepository) GetAll(limit, offset int) ([]models.StatusLog, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.StatusLog
	err := r.db.Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *statusLogRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.StatusLog{}).Count(&count).Error
	return count, err
}
