package dto

import (
	"backend/models"
	"time"
)

type CreateAdminRequest struct {
	AdminUsername string               `json:"adminUsername" binding:"required"`
	AdminPassword string               `json:"adminPassword" binding:"required"`
	AdminNama     string               `json:"adminName" binding:"required"`
	AdminEmail    string               `json:"adminEmail" binding:"required,email"`
	AdminRole     models.AdminRoleEnum `json:"adminRole" binding:"required,oneof=admin superadmin"`
}

type UpdateAdminRequest struct {
	AdminUsername string               `json:"adminUsername" binding:"required"`
	AdminPassword string               `json:"adminPassword,omitempty" binding:"omitempty"`
	AdminNama     string               `json:"adminName" binding:"required"`
	AdminEmail    string               `json:"adminEmail" binding:"required,email"`
	AdminRole     models.AdminRoleEnum `json:"adminRole" binding:"required,oneof=admin superadmin"`
}

type AdminResponse struct {
	AdminID       string               `json:"adminId"`
	AdminUsername string               `json:"adminUsername"`
	AdminNama     string               `json:"adminName"`
	AdminEmail    string               `json:"adminEmail"`
	AdminRole     models.AdminRoleEnum `json:"adminRole"`
	IsDeleted     bool                 `json:"isDeleted"`
	CreatedAt     time.Time            `json:"createdAt"`
	UpdatedAt     time.Time            `json:"updatedAt"`
	DeletedAt     *time.Time           `json:"deletedAt,omitempty"`
}
