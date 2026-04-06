package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"fmt"
	"strings"
)

type KomponenDarahService interface {
	Create(req dto.CreateKomponenDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error)
	GetByID(id int) (*dto.KomponenDarahResponse, error)
	GetAll(limit, offset int) ([]dto.KomponenDarahResponse, error)
	Update(id int, req dto.UpdateKomponenDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error)
	Delete(id int, userID *string, userName string, userRole string, userAgent *string) error
	ActivateKomponenDarah(id int, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error)
	DeactivateKomponenDarah(id int, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error)
}

type komponenDarahService struct {
	repo                   repository.KomponenDarahRepository
	systemAccessLogService SystemAccessLogService
}

func NewKomponenDarahService(repo repository.KomponenDarahRepository, systemAccessLogService SystemAccessLogService) KomponenDarahService {
	return &komponenDarahService{repo: repo, systemAccessLogService: systemAccessLogService}
}

func (s *komponenDarahService) Create(req dto.CreateKomponenDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error) {
	data := models.KomponenDarah{
		KomNama:  req.KomNama,
		KomKode:  req.KomKode,
		IsActive: req.IsActive,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"CREATE",
		stringPtr("komponen_darah"),
		stringPtr(fmt.Sprintf("%d", data.KomID)),
		fmt.Sprintf("Created komponen darah: %s (%s)", data.KomNama, data.KomKode),
		userAgent,
	)

	resp := mapKomponenToResponse(data)
	return &resp, nil
}

func (s *komponenDarahService) GetByID(id int) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func (s *komponenDarahService) GetAll(limit, offset int) ([]dto.KomponenDarahResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.KomponenDarahResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapKomponenToResponse(item))
	}
	return result, nil
}

func (s *komponenDarahService) Update(id int, req dto.UpdateKomponenDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	changes := []string{}
	if data.KomNama != req.KomNama {
		changes = append(changes, fmt.Sprintf("nama: %s → %s", data.KomNama, req.KomNama))
	}
	if data.KomKode != req.KomKode {
		changes = append(changes, fmt.Sprintf("kode: %s → %s", data.KomKode, req.KomKode))
	}
	if data.IsActive != req.IsActive {
		changes = append(changes, fmt.Sprintf("status: %v → %v", data.IsActive, req.IsActive))
	}

	data.KomNama = req.KomNama
	data.KomKode = req.KomKode
	data.IsActive = req.IsActive

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	if len(changes) > 0 {
		notes := fmt.Sprintf("Updated komponen darah %d: %s", id, strings.Join(changes, "; "))
		_, _ = s.systemAccessLogService.LogAction(
			userID,
			userName,
			userRole,
			"UPDATE",
			stringPtr("komponen_darah"),
			stringPtr(fmt.Sprintf("%d", id)),
			notes,
			userAgent,
		)
	}

	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func (s *komponenDarahService) Delete(id int, userID *string, userName string, userRole string, userAgent *string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(data); err != nil {
		return err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"DELETE",
		stringPtr("komponen_darah"),
		stringPtr(fmt.Sprintf("%d", id)),
		fmt.Sprintf("Deleted komponen darah: %s (%s)", data.KomNama, data.KomKode),
		userAgent,
	)

	return nil
}

func (s *komponenDarahService) ActivateKomponenDarah(id int, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.ActivateKomponenDarah(id)
	if err != nil {
		return nil, err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"ACTIVATE",
		stringPtr("komponen_darah"),
		stringPtr(fmt.Sprintf("%d", id)),
		fmt.Sprintf("Activated komponen darah: %s", data.KomNama),
		userAgent,
	)

	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func (s *komponenDarahService) DeactivateKomponenDarah(id int, userID *string, userName string, userRole string, userAgent *string) (*dto.KomponenDarahResponse, error) {
	data, err := s.repo.DeactivateKomponenDarah(id)
	if err != nil {
		return nil, err
	}

	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"DEACTIVATE",
		stringPtr("komponen_darah"),
		stringPtr(fmt.Sprintf("%d", id)),
		fmt.Sprintf("Deactivated komponen darah: %s", data.KomNama),
		userAgent,
	)

	resp := mapKomponenToResponse(*data)
	return &resp, nil
}

func mapKomponenToResponse(data models.KomponenDarah) dto.KomponenDarahResponse {
	return dto.KomponenDarahResponse{
		KomID:    data.KomID,
		KomNama:  data.KomNama,
		KomKode:  data.KomKode,
		IsActive: data.IsActive,
	}
}
