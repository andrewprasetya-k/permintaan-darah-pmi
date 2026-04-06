package repository

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type SystemAccessLogRepository interface {
	Create(data *models.SystemAccessLog) error
	GetByID(logID int64) (*models.SystemAccessLog, error)
	GetAll(limit, offset int) ([]models.SystemAccessLog, int64, error)

	GetByUserID(userID string, limit, offset int) ([]models.SystemAccessLog, int64, error)
	GetByAction(action string, limit, offset int) ([]models.SystemAccessLog, int64, error)
	GetByTargetTable(targetTable string, limit, offset int) ([]models.SystemAccessLog, int64, error)
	GetByTargetID(targetID string, limit, offset int) ([]models.SystemAccessLog, int64, error)
	GetByDateRange(startDate, endDate time.Time, limit, offset int) ([]models.SystemAccessLog, int64, error)
}

type systemAccessLogRepository struct {
	db *gorm.DB
}

func NewSystemAccessLogRepository(db *gorm.DB) SystemAccessLogRepository {
	return &systemAccessLogRepository{db: db}
}

func (r *systemAccessLogRepository) Create(data *models.SystemAccessLog) error {
	return r.db.Create(data).Error
}

func (r *systemAccessLogRepository) GetByID(logID int64) (*models.SystemAccessLog, error) {
	var data models.SystemAccessLog
	err := r.db.First(&data, "sys_log_id = ?", logID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *systemAccessLogRepository) GetAll(limit, offset int) ([]models.SystemAccessLog, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.SystemAccessLog
	var total int64

	err := r.db.Model(&models.SystemAccessLog{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *systemAccessLogRepository) GetByUserID(userID string, limit, offset int) ([]models.SystemAccessLog, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.SystemAccessLog
	var total int64

	query := r.db.Where("sys_user_id = ?", userID)
	query.Model(&models.SystemAccessLog{}).Count(&total)

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *systemAccessLogRepository) GetByAction(action string, limit, offset int) ([]models.SystemAccessLog, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.SystemAccessLog
	var total int64

	query := r.db.Where("sys_action = ?", action)
	query.Model(&models.SystemAccessLog{}).Count(&total)

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *systemAccessLogRepository) GetByTargetTable(targetTable string, limit, offset int) ([]models.SystemAccessLog, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.SystemAccessLog
	var total int64

	query := r.db.Where("sys_target_table = ?", targetTable)
	query.Model(&models.SystemAccessLog{}).Count(&total)

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *systemAccessLogRepository) GetByTargetID(targetID string, limit, offset int) ([]models.SystemAccessLog, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.SystemAccessLog
	var total int64

	query := r.db.Where("sys_target_id = ?", targetID)
	query.Model(&models.SystemAccessLog{}).Count(&total)

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *systemAccessLogRepository) GetByDateRange(startDate, endDate time.Time, limit, offset int) ([]models.SystemAccessLog, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.SystemAccessLog
	var total int64

	query := r.db.Where("created_at >= ? AND created_at <= ?", startDate, endDate)
	query.Model(&models.SystemAccessLog{}).Count(&total)

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *systemAccessLogRepository) Count() (int64, error) {
var count int64
err := r.db.Model(&models.SystemAccessLog{}).Where("is_deleted = ?", false).Count(&count).Error
return count, err
}
