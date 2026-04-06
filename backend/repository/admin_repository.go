package repository

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(data *models.Admin) error
	GetByID(adminID string) (*models.Admin, error)
	GetAll(limit, offset int) ([]models.Admin, error)
	GetByUsername(username string) (*models.Admin, error)
	Update(data *models.Admin) error
	Delete(data *models.Admin) error
	SoftDelete(adminID string) error
	Restore(adminID string) error
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
	err := r.db.Where("is_deleted = ?", false).First(&data, "admin_id = ?", adminID).Error
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
	err := r.db.Where("is_deleted = ?", false).Order("updated_at desc").Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, err
	}
	
	return list, nil
}

func (a *adminRepository) GetByUsername(username string) (*models.Admin, error) {
	var data models.Admin
	err := a.db.Where("is_deleted = ?", false).First(&data, "admin_username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *adminRepository) Update(data *models.Admin) error {
	return r.db.Save(data).Error
}

func (r *adminRepository) Delete(data *models.Admin) error {
	return r.db.Delete(data).Error
}

func (r *adminRepository) SoftDelete(adminID string) error {
	return r.db.Model(&models.Admin{}).Where("admin_id = ?", adminID).Updates(map[string]interface{}{
		"is_deleted": true,
		"deleted_at": time.Now(),
	}).Error
}

func (r *adminRepository) Restore(adminID string) error {
	return r.db.Model(&models.Admin{}).Where("admin_id = ?", adminID).Updates(map[string]interface{}{
		"is_deleted": false,
		"deleted_at": nil,
	}).Error
}
