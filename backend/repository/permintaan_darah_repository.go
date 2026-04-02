package repository

import (
	"backend/dto"
	"backend/models"

	"gorm.io/gorm"
)

type PermintaanDarahRepository interface {
	Create(data *models.PermintaanDarah) error
	GetByID(pdID string) (*models.PermintaanDarah, error)
	GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]models.PermintaanDarah, error)
	GetByRsID(rsID string, limit, offset int) ([]models.PermintaanDarah, error)
	Update(data *models.PermintaanDarah) error
	Delete(data *models.PermintaanDarah) error
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
	err := r.db.Preload("Details").First(&data, "pd_id = ?", pdID).Error
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

	query := r.db
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
	}

	var list []models.PermintaanDarah
	err := query.Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
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
	err := r.db.Where("pd_rs_id = ?", rsID).Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
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