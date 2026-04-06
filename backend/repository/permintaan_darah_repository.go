package repository

import (
	"backend/dto"
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type PermintaanDarahRepository interface {
	Create(data *models.PermintaanDarah) error
	GetByID(pdID string) (*models.PermintaanDarah, error)
	GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]models.PermintaanDarah, error)
	GetByRsID(rsID string, limit, offset int) ([]models.PermintaanDarah, error)
	Update(data *models.PermintaanDarah) error
	Delete(data *models.PermintaanDarah) error
	SoftDelete(pdID string) error
	Restore(pdID string) error
}

type permintaanDarahRepository struct {
	db *gorm.DB
}

func NewPermintaanDarahRepository(db *gorm.DB) PermintaanDarahRepository {
	return &permintaanDarahRepository{
		db: db,
	}
}

func (r *permintaanDarahRepository) Create(data *models.PermintaanDarah) error {
	return r.db.Create(data).Error
}

func (r *permintaanDarahRepository) GetByID(pdID string) (*models.PermintaanDarah, error) {
	var data models.PermintaanDarah
	err := r.db.Where("is_deleted = ?", false).Preload("Details.KomponenDarah").First(&data, "pd_id = ?", pdID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *permintaanDarahRepository) GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]models.PermintaanDarah, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	query := r.db.Where("is_deleted = ?", false)
	if filters != nil {
		if filters.Status != nil {
			query = query.Where("pd_status = ?", *filters.Status)
		}
		if filters.RsID != nil {
			query = query.Where("pd_rs_id = ?", *filters.RsID)
		}
		if filters.Gender != nil {
			query = query.Where("pd_gender = ?", *filters.Gender)
		}
		if filters.GolDarah != nil {
			query = query.Where("pd_gol_darah = ?", *filters.GolDarah)
		}
		if filters.StartDate != nil {
			query = query.Where("DATE(pd_tgl_permintaan) >= ?", filters.StartDate.Format("2006-01-02"))
		}
		if filters.EndDate != nil {
			query = query.Where("DATE(pd_tgl_permintaan) <= ?", filters.EndDate.Format("2006-01-02"))
		}
	}

	var list []models.PermintaanDarah
	err := query.Preload("Details.KomponenDarah").Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *permintaanDarahRepository) GetByRsID(rsID string, limit, offset int) ([]models.PermintaanDarah, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.PermintaanDarah
	err := r.db.Where("is_deleted = ?", false).Where("pd_rs_id = ?", rsID).Preload("Details.KomponenDarah").Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *permintaanDarahRepository) Update(data *models.PermintaanDarah) error {
	return r.db.Save(data).Error
}

func (r *permintaanDarahRepository) Delete(data *models.PermintaanDarah) error {
	return r.db.Delete(data).Error
}

func (r *permintaanDarahRepository) SoftDelete(pdID string) error {
	return r.db.Model(&models.PermintaanDarah{}).Where("pd_id = ?", pdID).Updates(map[string]interface{}{
		"is_deleted": true,
		"deleted_at": time.Now(),
	}).Error
}

func (r *permintaanDarahRepository) Restore(pdID string) error {
	return r.db.Model(&models.PermintaanDarah{}).Where("pd_id = ?", pdID).Updates(map[string]interface{}{
		"is_deleted": false,
		"deleted_at": nil,
	}).Error
}