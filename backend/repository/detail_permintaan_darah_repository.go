package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type DetailPermintaanDarahRepository interface {
	Create(data *models.DetailPermintaanDarah) error
	GetByID(dpdID string) (*models.DetailPermintaanDarah, error)
	GetAll(limit, offset int) ([]models.DetailPermintaanDarah, error)
	Update(data *models.DetailPermintaanDarah) error
	Delete(data *models.DetailPermintaanDarah) error
}

type detailPermintaanDarahRepository struct {
	db *gorm.DB
}

func NewDetailPermintaanDarahRepository(db *gorm.DB) DetailPermintaanDarahRepository {
	return &detailPermintaanDarahRepository{db: db}
}

func (r *detailPermintaanDarahRepository) Create(data *models.DetailPermintaanDarah) error {
	return r.db.Create(data).Error
}

func (r *detailPermintaanDarahRepository) GetByID(dpdID string) (*models.DetailPermintaanDarah, error) {
	var data models.DetailPermintaanDarah
	err := r.db.First(&data, "dpd_id = ?", dpdID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *detailPermintaanDarahRepository) GetAll(limit, offset int) ([]models.DetailPermintaanDarah, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.DetailPermintaanDarah
	err := r.db.Order("updated_at asc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *detailPermintaanDarahRepository) Update(data *models.DetailPermintaanDarah) error {
	return r.db.Save(data).Error
}

func (r *detailPermintaanDarahRepository) Delete(data *models.DetailPermintaanDarah) error {
	return r.db.Delete(data).Error
}
