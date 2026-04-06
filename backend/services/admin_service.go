package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"fmt"
	"strings"
)

type AdminService interface {
	Create(req dto.CreateAdminRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.AdminResponse, error)
	GetByID(id string) (*dto.AdminResponse, error)
	GetAll(limit, offset int) ([]dto.AdminResponse, error)
	Update(id string, req dto.UpdateAdminRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.AdminResponse, error)
	Delete(id string, userID *string, userName string, userRole string, userAgent *string) error
	Restore(id string, userID *string, userName string, userRole string, userAgent *string) error
}

type adminService struct {
	repo                     repository.AdminRepository
	systemAccessLogService SystemAccessLogService
}

func NewAdminService(repo repository.AdminRepository, systemAccessLogService SystemAccessLogService) AdminService {
	return &adminService{repo: repo, systemAccessLogService: systemAccessLogService}
}

func (s *adminService) Create(req dto.CreateAdminRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.AdminResponse, error) {
	hashedPassword, err := utils.HashPassword(req.AdminPassword)
	if err != nil {
		return nil, err
	}
	data := models.Admin{
		AdminUsername: req.AdminUsername,
		AdminPassword: hashedPassword,
		AdminNama:     req.AdminNama,
		AdminEmail:    req.AdminEmail,
		AdminRole:     req.AdminRole,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"CREATE",
		stringPtr("admins"),
		stringPtr(data.AdminID),
		fmt.Sprintf("Created admin: %s with role %s", data.AdminNama, data.AdminRole),
		userAgent,
	)

	resp := mapAdminToResponse(data)
	return &resp, nil
}

func (s *adminService) GetByID(id string) (*dto.AdminResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapAdminToResponse(*data)
	return &resp, nil
}

func (s *adminService) GetAll(limit, offset int) ([]dto.AdminResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.AdminResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapAdminToResponse(item))
	}
	return result, nil
}

func (s *adminService) Update(id string, req dto.UpdateAdminRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.AdminResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	changes := []string{}
	if data.AdminUsername != req.AdminUsername {
		changes = append(changes, fmt.Sprintf("username: %s → %s", data.AdminUsername, req.AdminUsername))
	}
	if data.AdminNama != req.AdminNama {
		changes = append(changes, fmt.Sprintf("nama: %s → %s", data.AdminNama, req.AdminNama))
	}
	if data.AdminEmail != req.AdminEmail {
		changes = append(changes, fmt.Sprintf("email: %s → %s", data.AdminEmail, req.AdminEmail))
	}
	if data.AdminRole != req.AdminRole {
		changes = append(changes, fmt.Sprintf("role: %s → %s", data.AdminRole, req.AdminRole))
	}

	data.AdminUsername = req.AdminUsername
	data.AdminPassword = req.AdminPassword
	data.AdminNama = req.AdminNama
	data.AdminEmail = req.AdminEmail
	data.AdminRole = req.AdminRole

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	if len(changes) > 0 {
		notes := fmt.Sprintf("Updated admin %s: %s", id, strings.Join(changes, "; "))
		_, _ = s.systemAccessLogService.LogAction(
			userID,
			userName,
			userRole,
			"UPDATE",
			stringPtr("admins"),
			stringPtr(id),
			notes,
			userAgent,
		)
	}

	resp := mapAdminToResponse(*data)
	return &resp, nil
}

func (s *adminService) Delete(id string, userID *string, userName string, userRole string, userAgent *string) error {

	admin, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.SoftDelete(id); err != nil {
		return err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"SOFT_DELETE",
		stringPtr("admins"),
		stringPtr(id),
		fmt.Sprintf("Soft deleted admin: %s", admin.AdminNama),
		userAgent,
	)

	return nil
}

func (s *adminService) Restore(id string, userID *string, userName string, userRole string, userAgent *string) error {

	admin, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.Restore(id); err != nil {
		return err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"RESTORE",
		stringPtr("admins"),
		stringPtr(id),
		fmt.Sprintf("Restored admin: %s", admin.AdminNama),
		userAgent,
	)

	return nil
}

func mapAdminToResponse(data models.Admin) dto.AdminResponse {
	return dto.AdminResponse{
		AdminID:       data.AdminID,
		AdminUsername: data.AdminUsername,
		AdminNama:     data.AdminNama,
		AdminEmail:    data.AdminEmail,
		AdminRole:     data.AdminRole,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
		DeletedAt:     data.DeletedAt,
	}
}

func stringPtr(s string) *string {
	return &s
}
