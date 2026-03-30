package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type PermintaanDarahRepository interface {
	Create(data *models.PermintaanDarah) error
	GetByID(pdID string) (*models.PermintaanDarah, error)
	List(limit, offset int) ([]models.PermintaanDarah, error)
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
	err := r.db.First(&data, "pd_id = ?", pdID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *permintaanDarahRepository) List(limit, offset int) ([]models.PermintaanDarah, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.PermintaanDarah
	err := r.db.Limit(limit).Offset(offset).Find(&list).Error
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
