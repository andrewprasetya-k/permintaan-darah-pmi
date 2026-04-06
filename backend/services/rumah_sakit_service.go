package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"fmt"
	"strings"
)

type RumahSakitService interface {
	Create(req dto.CreateRumahSakitRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.RumahSakitResponse, error)
	GetByID(id string) (*dto.RumahSakitResponse, error)
	GetAll(limit, offset int) ([]dto.RumahSakitResponse, error)
	Update(id string, req dto.UpdateRumahSakitRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.RumahSakitResponse, error)
	Delete(id string, userID *string, userName string, userRole string, userAgent *string) error
	Restore(id string, userID *string, userName string, userRole string, userAgent *string) error

	//filter purpose
	GetDistinctRSNama() ([]dto.RumahSakitDistinctNamaResponse, error)
}

type rumahSakitService struct {
	repo                   repository.RumahSakitRepository
	systemAccessLogService SystemAccessLogService
}

func NewRumahSakitService(repo repository.RumahSakitRepository, systemAccessLogService SystemAccessLogService) RumahSakitService {
	return &rumahSakitService{repo: repo, systemAccessLogService: systemAccessLogService}
}

func (s *rumahSakitService) Create(req dto.CreateRumahSakitRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.RumahSakitResponse, error) {
	hashedPassword, err := utils.HashPassword(req.RSPassword)
	if err != nil {
		return nil, err
	}
	data := models.RumahSakit{
		RSNama:     req.RSNama,
		RSNoTelp:   req.RSNoTelp,
		RSAlamat:   req.RSAlamat,
		RSEmail:    req.RSEmail,
		RSUsername: req.RSUsername,
		RSPassword: hashedPassword,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"CREATE",
		stringPtr("rumah_sakit"),
		stringPtr(data.RSID),
		fmt.Sprintf("Created rumah sakit: %s", data.RSNama),
		userAgent,
	)

	resp := mapRumahSakitToResponse(data)
	return &resp, nil
}

func (s *rumahSakitService) GetByID(id string) (*dto.RumahSakitResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapRumahSakitToResponse(*data)
	return &resp, nil
}

func (s *rumahSakitService) GetAll(limit, offset int) ([]dto.RumahSakitResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.RumahSakitResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapRumahSakitToResponse(item))
	}
	return result, nil
}

func (s *rumahSakitService) GetDistinctRSNama() ([]dto.RumahSakitDistinctNamaResponse, error) {
	return s.repo.GetDistinctRSNama()
}

func (s *rumahSakitService) Update(id string, req dto.UpdateRumahSakitRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.RumahSakitResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Track changes for logging
	changes := []string{}
	if data.RSNama != req.RSNama {
		changes = append(changes, fmt.Sprintf("nama: %s → %s", data.RSNama, req.RSNama))
	}
	if data.RSNoTelp != req.RSNoTelp {
		changes = append(changes, fmt.Sprintf("telp: %s → %s", data.RSNoTelp, req.RSNoTelp))
	}
	if data.RSAlamat != req.RSAlamat {
		changes = append(changes, fmt.Sprintf("alamat: %s → %s", data.RSAlamat, req.RSAlamat))
	}
	oldEmail := ""
	if data.RSEmail != nil {
		oldEmail = *data.RSEmail
	}
	newEmail := ""
	if req.RSEmail != nil {
		newEmail = *req.RSEmail
	}
	if oldEmail != newEmail {
		changes = append(changes, fmt.Sprintf("email: %s → %s", oldEmail, newEmail))
	}

	data.RSNama = req.RSNama
	data.RSNoTelp = req.RSNoTelp
	data.RSAlamat = req.RSAlamat
	data.RSEmail = req.RSEmail
	data.RSPassword = *req.RSPassword

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	if len(changes) > 0 {
		notes := fmt.Sprintf("Updated rumah sakit %s: %s", id, strings.Join(changes, "; "))
		_, _ = s.systemAccessLogService.LogAction(
			userID,
			userName,
			userRole,
			"UPDATE",
			stringPtr("rumah_sakit"),
			stringPtr(id),
			notes,
			userAgent,
		)
	}

	resp := mapRumahSakitToResponse(*data)
	return &resp, nil
}

func (s *rumahSakitService) Delete(id string, userID *string, userName string, userRole string, userAgent *string) error {
	rs, err := s.repo.GetByID(id)
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
		stringPtr("rumah_sakit"),
		stringPtr(id),
		fmt.Sprintf("Soft deleted rumah sakit: %s", rs.RSNama),
		userAgent,
	)

	return nil
}

func (s *rumahSakitService) Restore(id string, userID *string, userName string, userRole string, userAgent *string) error {
	rs, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.Restore(id); err != nil {
		return err
	}

	// ✅ Auto-log: Restore
	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"RESTORE",
		stringPtr("rumah_sakit"),
		stringPtr(id),
		fmt.Sprintf("Restored rumah sakit: %s", rs.RSNama),
		userAgent,
	)

	return nil
}

func mapRumahSakitToResponse(data models.RumahSakit) dto.RumahSakitResponse {
	return dto.RumahSakitResponse{
		RSID:       data.RSID,
		RSNama:     data.RSNama,
		RSNoTelp:   data.RSNoTelp,
		RSAlamat:   data.RSAlamat,
		RSEmail:    data.RSEmail,
		RSUsername: data.RSUsername,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt,
	}
}
