package repository

import (
	"backend/dto"
	"backend/models"

	"gorm.io/gorm"
)

type RumahSakitRepository interface {
	Create(data *models.RumahSakit) error
	GetByID(rsID string) (*models.RumahSakit, error)
	GetAll(limit, offset int) ([]models.RumahSakit, error)
	GetByUsername(username string) (*models.RumahSakit, error)
	Update(data *models.RumahSakit) error
	Delete(data *models.RumahSakit) error

	//filter purpose
	GetDistinctRSNama() ([]dto.RumahSakitDistinctNamaResponse, error)
}

type rumahSakitRepository struct {
	db *gorm.DB
}

func NewRumahSakitRepository(db *gorm.DB) RumahSakitRepository {
	return &rumahSakitRepository{db: db}
}

func (r *rumahSakitRepository) Create(data *models.RumahSakit) error {
	return r.db.Create(data).Error
}

func (r *rumahSakitRepository) GetByID(rsID string) (*models.RumahSakit, error) {
	var data models.RumahSakit
	err := r.db.First(&data, "rs_id = ?", rsID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *rumahSakitRepository) GetAll(limit, offset int) ([]models.RumahSakit, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.RumahSakit
	err := r.db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *rumahSakitRepository) GetDistinctRSNama() ([]dto.RumahSakitDistinctNamaResponse, error) {
	var list []dto.RumahSakitDistinctNamaResponse
	err := r.db.Model(&models.RumahSakit{}).Select("DISTINCT rs_nama, rs_id").Order("rs_nama ASC").Scan(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}



func (r *rumahSakitRepository) GetByUsername(username string) (*models.RumahSakit, error) {
	var data models.RumahSakit
	err := r.db.First(&data, "rs_username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *rumahSakitRepository) Update(data *models.RumahSakit) error {
	return r.db.Save(data).Error
}

func (r *rumahSakitRepository) Delete(data *models.RumahSakit) error {
	return r.db.Delete(data).Error
}
