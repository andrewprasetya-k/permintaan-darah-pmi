package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(data *models.Admin) error
	GetByID(adminID string) (*models.Admin, error)
	GetAll(limit, offset int) ([]models.Admin, error)
	Update(data *models.Admin) error
	Delete(data *models.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) Create(data *models.Admin) error {
	return r.db.Create(data).Error
}

func (r *adminRepository) GetByID(adminID string) (*models.Admin, error) {
	var data models.Admin
	err := r.db.First(&data, "admin_id = ?", adminID).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *adminRepository) GetAll(limit, offset int) ([]models.Admin, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var list []models.Admin
	err := r.db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *adminRepository) Update(data *models.Admin) error {
	return r.db.Save(data).Error
}

func (r *adminRepository) Delete(data *models.Admin) error {
	return r.db.Delete(data).Error
}
