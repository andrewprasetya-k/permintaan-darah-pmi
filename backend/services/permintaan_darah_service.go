package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"errors"
	"fmt"
)

type PermintaanDarahService interface {
	Create(req dto.CreatePermintaanDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error)
	GetByID(id string) (*dto.PermintaanDarahResponse, error)
	GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]dto.PermintaanDarahGetAllResponse, int, error)
	GetByRsID(rsID string, limit, offset int) ([]dto.PermintaanDarahGetAllResponse, int, error)
	Update(id string, req dto.UpdatePermintaanDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error)
	UpdateMyRequest(id string, rsID string, req dto.UpdatePermintaanDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error)
	Delete(id string, userID *string, userName string, userRole string, userAgent *string) error
	Restore(id string, userID *string, userName string, userRole string, userAgent *string) error

	UpdateStatus(pdID string, newStatus models.PermintaanStatusEnum, reason *string, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error)
	UpdateStatusWithOwnershipCheck(pdID string, newStatus models.PermintaanStatusEnum, reason *string, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error)
}

type permintaanDarahService struct {
	repo                   repository.PermintaanDarahRepository
	statusLogRepo          repository.StatusLogRepository
	systemAccessLogService SystemAccessLogService
	wsHub                  *Hub
}

func NewPermintaanDarahService(repo repository.PermintaanDarahRepository, statusLogRepo repository.StatusLogRepository, systemAccessLogService SystemAccessLogService, wsHub *Hub) PermintaanDarahService {
	return &permintaanDarahService{repo: repo, statusLogRepo: statusLogRepo, systemAccessLogService: systemAccessLogService, wsHub: wsHub}
}

func (s *permintaanDarahService) Create(req dto.CreatePermintaanDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error) {
	// Validate input
	if err := utils.ValidatePermintaanDarahInput(req.PDHemoglobin, req.PDTglLahir, req.PDTglPermintaan); err != nil {
		return nil, err
	}

	data := models.PermintaanDarah{
		PDRsID:              req.PDRsID,
		PDNamaPasien:        req.PDNamaPasien,
		PDNoRM:              req.PDNoRM,
		PDTempatLahir:       req.PDTempatLahir,
		PDTglLahir:          req.PDTglLahir,
		PDGender:            req.PDGender,
		PDGolDarah:          req.PDGolDarah,
		PDRhesus:            req.PDRhesus,
		PDHemoglobin:        req.PDHemoglobin,
		PDRuangBagianKelas:  req.PDRuangBagianKelas,
		PDPernahTransfusi:   req.PDPernahTransfusi,
		PDIndikasiTransfusi: req.PDIndikasiTransfusi,
		PDPernahHamil:       req.PDPernahHamil,
		PDStatus:            req.PDStatus,
		PDCancelReason:      req.PDCancelReason,
		PDTglPermintaan:     req.PDTglPermintaan,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
		
	}

	// ✅ Auto-log: Create
	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"CREATE",
		stringPtr("permintaan_darah"),
		stringPtr(data.PDID),
		fmt.Sprintf("Created permintaan darah untuk pasien: %s (gol darah: %v%v)", data.PDNamaPasien, data.PDGolDarah, data.PDRhesus),
		userAgent,
	)

	resp := mapPermintaanToResponse(data)
	return &resp, nil
}

func (s *permintaanDarahService) GetByID(id string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapPermintaanToResponse(*data)
	return &resp, nil
}

func (s *permintaanDarahService) GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]dto.PermintaanDarahGetAllResponse, int, error) {
	list, err := s.repo.GetAll(filters, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.PermintaanDarahGetAllResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapPermintaanToGetAllResponse(item))
	}
	return result, int(total), nil
}

func (s *permintaanDarahService) GetByRsID(rsID string, limit, offset int) ([]dto.PermintaanDarahGetAllResponse, int, error) {
	list, total, err := s.repo.GetByRsID(rsID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.PermintaanDarahGetAllResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapPermintaanToGetAllResponse(item))
	}
	return result, int(total), nil
}

func (s *permintaanDarahService) Update(id string, req dto.UpdatePermintaanDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Validate input
	if err := utils.ValidatePermintaanDarahInput(req.PDHemoglobin, req.PDTglLahir, req.PDTglPermintaan); err != nil {
		return nil, err
	}

	changes := []string{}
	if data.PDNamaPasien != req.PDNamaPasien {
		changes = append(changes, fmt.Sprintf("nama pasien: %s → %s", data.PDNamaPasien, req.PDNamaPasien))
	}
	if data.PDGolDarah != req.PDGolDarah {
		changes = append(changes, fmt.Sprintf("gol darah: %v → %v", data.PDGolDarah, req.PDGolDarah))
	}
	if data.PDStatus != req.PDStatus {
		changes = append(changes, fmt.Sprintf("status: %s → %s", data.PDStatus, req.PDStatus))
	}

	data.PDRsID = req.PDRsID
	data.PDNamaPasien = req.PDNamaPasien
	data.PDNoRM = req.PDNoRM
	data.PDTempatLahir = req.PDTempatLahir
	data.PDTglLahir = req.PDTglLahir
	data.PDGender = req.PDGender
	data.PDGolDarah = req.PDGolDarah
	data.PDRhesus = req.PDRhesus
	data.PDHemoglobin = req.PDHemoglobin
	data.PDRuangBagianKelas = req.PDRuangBagianKelas
	data.PDPernahTransfusi = req.PDPernahTransfusi
	data.PDIndikasiTransfusi = req.PDIndikasiTransfusi
	data.PDPernahHamil = req.PDPernahHamil
	data.PDStatus = req.PDStatus
	data.PDCancelReason = req.PDCancelReason
	data.PDTglPermintaan = req.PDTglPermintaan

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	if len(changes) > 0 {
		notes := fmt.Sprintf("Updated permintaan darah %s: changed", id)
		_, _ = s.systemAccessLogService.LogAction(
			userID,
			userName,
			userRole,
			"UPDATE",
			stringPtr("permintaan_darah"),
			stringPtr(id),
			notes,
			userAgent,
		)
	}

	resp := mapPermintaanToResponse(*data)

	// Broadcast WebSocket event
	if s.wsHub != nil {
		msg := NewWebSocketMessage("update_permintaan", "UPDATE", data.PDID, "permintaan_darah", resp)
		s.wsHub.Broadcast(msg)
	}
	return &resp, nil
}

func (s *permintaanDarahService) Delete(id string, userID *string, userName string, userRole string, userAgent *string) error {
	pd, err := s.repo.GetByID(id)
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
		stringPtr("permintaan_darah"),
		stringPtr(id),
		fmt.Sprintf("Soft deleted permintaan darah untuk pasien: %s", pd.PDNamaPasien),
		userAgent,
	)

	return nil
}

func (s *permintaanDarahService) Restore(id string, userID *string, userName string, userRole string, userAgent *string) error {
	pd, err := s.repo.GetByID(id)
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
		stringPtr("permintaan_darah"),
		stringPtr(id),
		fmt.Sprintf("Restored permintaan darah untuk pasien: %s", pd.PDNamaPasien),
		userAgent,
	)

	return nil
}

func (s *permintaanDarahService) UpdateMyRequest(id string, rsID string, req dto.UpdatePermintaanDarahRequest, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure rumah sakit can only update their own request
	if data.PDRsID != rsID {
		return nil, errors.New("not authorized to update this request")
	}

	changes := []string{}
	if data.PDNamaPasien != req.PDNamaPasien {
		changes = append(changes, fmt.Sprintf("nama pasien: %s → %s", data.PDNamaPasien, req.PDNamaPasien))
	}
	if data.PDGolDarah != req.PDGolDarah {
		changes = append(changes, fmt.Sprintf("gol darah: %v → %v", data.PDGolDarah, req.PDGolDarah))
	}

	data.PDNamaPasien = req.PDNamaPasien
	data.PDNoRM = req.PDNoRM
	data.PDTempatLahir = req.PDTempatLahir
	data.PDTglLahir = req.PDTglLahir
	data.PDGender = req.PDGender
	data.PDGolDarah = req.PDGolDarah
	data.PDRhesus = req.PDRhesus
	data.PDHemoglobin = req.PDHemoglobin
	data.PDRuangBagianKelas = req.PDRuangBagianKelas
	data.PDPernahTransfusi = req.PDPernahTransfusi
	data.PDIndikasiTransfusi = req.PDIndikasiTransfusi
	data.PDPernahHamil = req.PDPernahHamil
	data.PDTglPermintaan = req.PDTglPermintaan

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	if len(changes) > 0 {
		notes := fmt.Sprintf("Rumah sakit updated their own permintaan darah %s", id)
		_, _ = s.systemAccessLogService.LogAction(
			userID,
			userName,
			userRole,
			"UPDATE",
			stringPtr("permintaan_darah"),
			stringPtr(id),
			notes,
			userAgent,
		)
	}

	resp := mapPermintaanToResponse(*data)
	return &resp, nil
}

func (s *permintaanDarahService) UpdateStatus(pdID string, newStatus models.PermintaanStatusEnum, reason *string, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(pdID)
	if err != nil {
		return nil, err
	}

	if data.PDStatus == "selesai" || data.PDStatus == "dibatalkan" {
		return nil, errors.New("cannot update status of a completed or cancelled request")
	}

	oldStatus := data.PDStatus
	data.PDStatus = newStatus
	if newStatus == "dibatalkan" && reason == nil {
		return nil, errors.New("reason is required")
	} else if newStatus == "dibatalkan" && reason != nil {
		data.PDCancelReason = reason
	}

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	// Auto-create status log
	notes := fmt.Sprintf("Status berubah dari %s menjadi %s", oldStatus, newStatus)
	if reason != nil {
		notes = fmt.Sprintf("%s. Alasan: %s", notes, *reason)
	}

	log := models.StatusLog{
		LogPdID:    pdID,
		LogAdminID: userID,

		LogStatusFrom: &oldStatus,
		LogStatusTo:   newStatus,
		LogNotes:      &notes,
	}

	if err := s.statusLogRepo.Create(&log); err != nil {
		// Log creation error tapi jangan return error (status update sudah sukses)
		fmt.Printf("Warning: Failed to create status log: %v\n", err)
	}

	logNotes := fmt.Sprintf("Updated status permintaan darah dari %s → %s", oldStatus, newStatus)
	if reason != nil {
		logNotes += fmt.Sprintf(" (reason: %s)", *reason)
	}
	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"UPDATE_STATUS",
		stringPtr("permintaan_darah"),
		stringPtr(pdID),
		logNotes,
		userAgent,
	)

	resp := mapPermintaanToResponse(*data)
	// Broadcast WebSocket event
	if s.wsHub != nil {
		msg := NewWebSocketMessage("status_change", "UPDATE", data.PDID, "permintaan_darah", resp)
		s.wsHub.Broadcast(msg)
	}

	// Broadcast WebSocket event
	if s.wsHub != nil {
		msg := NewWebSocketMessage("update_status", "UPDATE", pdID, "permintaan_darah", resp)
		s.wsHub.Broadcast(msg)
	}
	return &resp, nil
}

func mapPermintaanToGetAllResponse(data models.PermintaanDarah) dto.PermintaanDarahGetAllResponse {
	return dto.PermintaanDarahGetAllResponse{
	PDID:	   data.PDID,

	// Data Pasien
	PDNamaPasien: data.PDNamaPasien,
	PDGender:	 data.PDGender,
	PDGolDarah:   data.PDGolDarah,
	PDRhesus:	 data.PDRhesus,

	// Status
	PDStatus: data.PDStatus,
	CreatedAt: data.CreatedAt,     
	UpdatedAt: data.UpdatedAt,
	DeletedAt: data.DeletedAt,
	}
}

func mapPermintaanToResponse(data models.PermintaanDarah) dto.PermintaanDarahResponse {
	details := make([]dto.DetailPermintaanDarahResponse, 0, len(data.Details))
	for _, d := range data.Details {
		details = append(details, dto.DetailPermintaanDarahResponse{
			DPDID:            d.DPDID,
			DPDPDID:          d.DPDPDID,
			KomponenNama:     d.KomponenDarah.KomNama,
			DPDGolonganDarah: d.DPDGolonganDarah,
			DPDRhesus:        d.DPDRhesus,
			DPDJmlKantong:    d.DPDJmlKantong,
			DPDTglDiperlukan: d.DPDTglDiperlukan,
			CreatedAt:        d.CreatedAt,
		})
	}

	return dto.PermintaanDarahResponse{
		PDID:                data.PDID,
		PDRsID:              data.PDRsID,
		PDNamaPasien:        data.PDNamaPasien,
		PDNoRM:              data.PDNoRM,
		PDTempatLahir:       data.PDTempatLahir,
		PDTglLahir:          data.PDTglLahir,
		PDGender:            data.PDGender,
		PDGolDarah:          data.PDGolDarah,
		PDRhesus:            data.PDRhesus,
		PDHemoglobin:        data.PDHemoglobin,
		PDRuangBagianKelas:  data.PDRuangBagianKelas,
		PDPernahTransfusi:   data.PDPernahTransfusi,
		PDIndikasiTransfusi: data.PDIndikasiTransfusi,
		PDPernahHamil:       data.PDPernahHamil,
		PDStatus:            data.PDStatus,
		PDCancelReason:      data.PDCancelReason,
		PDTglPermintaan:     data.PDTglPermintaan,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
		DeletedAt:           data.DeletedAt,
		Details:             details,
	}
}

func (s *permintaanDarahService) UpdateStatusWithOwnershipCheck(pdID string, newStatus models.PermintaanStatusEnum, reason *string, userID *string, userName string, userRole string, userAgent *string) (*dto.PermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(pdID)
	if err != nil {
		return nil, err
	}

	// Ownership check: rumah sakit can only update status of their own request
	if userRole == "rumah_sakit" && (userID == nil || *userID != data.PDRsID) {
		return nil, errors.New("not authorized to update status of this request")
	}

	if data.PDStatus == "selesai" || data.PDStatus == "dibatalkan" {
		return nil, errors.New("cannot update status of a completed or cancelled request")
	}

	oldStatus := data.PDStatus
	data.PDStatus = newStatus
	if newStatus == "dibatalkan" && reason == nil {
		return nil, errors.New("reason is required")
	} else if newStatus == "dibatalkan" && reason != nil {
		data.PDCancelReason = reason
	}

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	// Auto-create status log
	notes := fmt.Sprintf("Status berubah dari %s menjadi %s", oldStatus, newStatus)
	if reason != nil {
		notes = fmt.Sprintf("%s. Alasan: %s", notes, *reason)
	}

	log := models.StatusLog{
		LogPdID:    pdID,
		LogAdminID: userID,

		LogStatusFrom: &oldStatus,
		LogStatusTo:   newStatus,
		LogNotes:      &notes,
	}

	if err := s.statusLogRepo.Create(&log); err != nil {
		fmt.Printf("Warning: Failed to create status log: %v\n", err)
	}

	logNotes := fmt.Sprintf("Updated status permintaan darah dari %s → %s", oldStatus, newStatus)
	if reason != nil {
		logNotes += fmt.Sprintf(" (reason: %s)", *reason)
	}
	_, _ = s.systemAccessLogService.LogAction(
		userID,
		userName,
		userRole,
		"UPDATE_STATUS",
		stringPtr("permintaan_darah"),
		stringPtr(pdID),
		logNotes,
		userAgent,
	)

	resp := mapPermintaanToResponse(*data)
	return &resp, nil
}
