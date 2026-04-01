package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type PasienRepository interface {
	Create(data *models.Pasien) error
	GetByID(psnID string) (*models.Pasien, error)
	GetAll(limit, offset int) ([]models.Pasien, error)
	Update(data *models.Pasien) error
	Delete(data *models.Pasien) error
}

type pasienRepository struct {
	db *gorm.DB
}

func NewPasienRepository(db *gorm.DB) PasienRepository {
	return &pasienRepository{db: db}
}

func (r *pasienRepository) Create(data *models.Pasien) error {
	return r.db.Create(data).Error
}

func (r *pasienRepository) GetByID(psnID string) (*models.Pasien, error) {
	var data models.Pasien
	err := r.db.First(&data, "psn_id = ?", psnID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pasienRepository) GetAll(limit, offset int) ([]models.Pasien, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.Pasien
	err := r.db.Order("updated_at asc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *pasienRepository) Update(data *models.Pasien) error {
	return r.db.Save(data).Error
}

func (r *pasienRepository) Delete(data *models.Pasien) error {
	return r.db.Delete(data).Error
}
