package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type KomponenDarahRepository interface {
	Create(data *models.KomponenDarah) error
	GetByID(komID int) (*models.KomponenDarah, error)
	GetAll(limit, offset int) ([]models.KomponenDarah, error)
	Update(data *models.KomponenDarah) error
	Delete(data *models.KomponenDarah) error
	ActivateKomponenDarah(komID int) (*models.KomponenDarah, error)
	DeactivateKomponenDarah(komID int) (*models.KomponenDarah, error)
}

type komponenDarahRepository struct {
	db *gorm.DB
}

func NewKomponenDarahRepository(db *gorm.DB) KomponenDarahRepository {
	return &komponenDarahRepository{db: db}
}

func (r *komponenDarahRepository) Create(data *models.KomponenDarah) error {
	return r.db.Create(data).Error
}

func (r *komponenDarahRepository) GetByID(komID int) (*models.KomponenDarah, error) {
	var data models.KomponenDarah
	err := r.db.First(&data, "kom_id = ?", komID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *komponenDarahRepository) GetAll(limit, offset int) ([]models.KomponenDarah, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.KomponenDarah
	err := r.db.Order("kom_id asc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *komponenDarahRepository) Update(data *models.KomponenDarah) error {
	return r.db.Save(data).Error
}

func (r *komponenDarahRepository) Delete(data *models.KomponenDarah) error {
	return r.db.Delete(data).Error
}

func (r *komponenDarahRepository) ActivateKomponenDarah(komID int) (*models.KomponenDarah, error) {
	var data models.KomponenDarah
	err := r.db.Model(&data).Where("kom_id = ?", komID).Update("is_active", true).Error
	if err != nil {
		return nil, err
	}

	err = r.db.First(&data, "kom_id = ?", komID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *komponenDarahRepository) DeactivateKomponenDarah(komID int) (*models.KomponenDarah, error) {
	var data models.KomponenDarah
	err := r.db.Model(&data).Where("kom_id = ?", komID).Update("is_active", false).Error
	if err != nil {
		return nil, err
	}

	err = r.db.First(&data, "kom_id = ?", komID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}