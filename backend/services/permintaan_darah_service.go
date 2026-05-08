package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
	"gorm.io/gorm"
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
	GenerateBlankoPDF(id string) ([]byte, string, error)
}

type permintaanDarahService struct {
	repo                   repository.PermintaanDarahRepository
	statusLogRepo          repository.StatusLogRepository
	systemAccessLogService SystemAccessLogService
	wsHub                  *Hub
	db                     *gorm.DB
}

func NewPermintaanDarahService(repo repository.PermintaanDarahRepository, statusLogRepo repository.StatusLogRepository, systemAccessLogService SystemAccessLogService, wsHub *Hub, db *gorm.DB) PermintaanDarahService {
	return &permintaanDarahService{
		repo:                   repo,
		statusLogRepo:          statusLogRepo,
		systemAccessLogService: systemAccessLogService,
		wsHub:                  wsHub,
		db:                     db,
	}
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
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		for _, detail := range req.Details {
			detailData := models.DetailPermintaanDarah{
				DPDPDID:          data.PDID,
				DPDKomID:         detail.DPDKomID,
				DPDGolonganDarah: detail.DPDGolonganDarah,
				DPDRhesus:        detail.DPDRhesus,
				DPDJmlKantong:    detail.DPDJmlKantong,
			}
			if err := tx.Create(&detailData).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
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

	createdPermintaan, err := s.repo.GetByID(data.PDID)
	if err != nil {
		return nil, err
	}

	resp := mapPermintaanToResponse(*createdPermintaan)
	//broadcast WebSocket event
	if s.wsHub != nil {
		msg := NewWebSocketMessage("new_permintaan", "CREATE", data.PDID, "permintaan_darah", resp)
		s.wsHub.Broadcast(msg)
	}
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

func (s *permintaanDarahService) GenerateBlankoPDF(id string) ([]byte, string, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, "", err
	}

	var rs models.RumahSakit
	rsName := data.PDRsID
	if err := s.db.Select("rs_nama").First(&rs, "rs_id = ?", data.PDRsID).Error; err == nil {
		rsName = rs.RSNama
	}

	const pageWidth = 210.0
	const pageHeight = 330.0
	templatePath, err := resolveBlankoTemplatePath()
	if err != nil {
		return nil, "", err
	}
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size: gofpdf.SizeType{
			Wd: pageWidth,
			Ht: pageHeight,
		},
	})
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()

	tpl := gofpdi.ImportPage(pdf, templatePath, 1, "/MediaBox")
	gofpdi.UseImportedTemplate(pdf, tpl, 0, 0, pageWidth, pageHeight)

	pdf.SetFont("Helvetica", "", 9)
	pdf.SetTextColor(0, 0, 0)

	write := func(x, y, w float64, text string) {
		pdf.SetXY(x, y)
		pdf.MultiCell(w, 4.8, safePDFText(text), "", "L", false)
	}

	firstDetail := ""
	totalKantong := 0
	components := []string{}
	for _, detail := range data.Details {
		totalKantong += detail.DPDJmlKantong
		if detail.KomponenDarah.KomNama != "" {
			components = append(components, detail.KomponenDarah.KomNama)
		}
		if firstDetail == "" {
			firstDetail = fmt.Sprintf("%s%s", detail.DPDGolonganDarah, detail.DPDRhesus)
		}
	}

	blood := "-"
	if data.PDGolDarah != nil {
		blood = string(*data.PDGolDarah)
		if data.PDRhesus != nil {
			blood += string(*data.PDRhesus)
		}
	}
	if firstDetail != "" {
		blood = firstDetail
	}

	quantity := "-"
	if totalKantong > 0 {
		quantity = fmt.Sprintf("%d kantong", totalKantong)
	}

	write(80, 41.5, 100, rsName)
	write(80, 48.2, 100, "-")
	write(80, 55.0, 100, data.PDNamaPasien)
	write(80, 61.8, 100, derefString(data.PDTempatLahir, "-"))
	write(80, 68.5, 100, derefStringPtr(data.PDNoRM, "-"))
	write(80, 75.3, 100, formatIndonesianDate(data.PDTglLahir))
	write(80, 82.0, 100, fmt.Sprintf("%d tahun - %s", ageInYears(data.PDTglLahir, data.PDTglPermintaan), genderLabel(data.PDGender)))
	write(80, 88.8, 100, derefStringPtr(data.PDIndikasiTransfusi, "-"))
	write(80, 95.5, 100, derefStringPtr(data.PDRuangBagianKelas, "-"))
	write(80, 102.3, 100, derefStringPtr(data.PDIndikasiTransfusi, "-"))
	write(80, 109.0, 100, transfusionLabel(data.PDPernahTransfusi))
	write(80, 115.8, 100, pregnancyLabel(data.PDPernahHamil))
	write(80, 122.5, 100, blood)
	write(80, 129.3, 100, hemoglobinLabel(data.PDHemoglobin))
	write(80, 136.0, 100, formatIndonesianDateTime(data.PDTglPermintaan))
	write(80, 142.8, 100, quantity)
	write(80, 149.5, 100, strings.Join(uniqueStrings(components), ", "))

	write(18, 182.5, 45, "-")
	write(82, 182.5, 35, derefStringPtr(data.PDNoRM, "-"))
	write(18, 189.2, 45, blood)
	write(82, 189.2, 35, rhesusOnly(data.PDRhesus))

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, "", err
	}

	filename := fmt.Sprintf("blanko-permintaan-darah-%s.pdf", data.PDID)
	return buf.Bytes(), filename, nil
}

func (s *permintaanDarahService) GetAll(filters *dto.PermintaanDarahFilters, limit, offset int) ([]dto.PermintaanDarahGetAllResponse, int, error) {
	list, err := s.repo.GetAll(filters, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.Count(filters)
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
		LogPdID:       pdID,
		LogAdminID:    userID,
		LogAdminNama:  userName,
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
	return &resp, nil
}

func mapPermintaanToGetAllResponse(data models.PermintaanDarah) dto.PermintaanDarahGetAllResponse {
	return dto.PermintaanDarahGetAllResponse{
		PDID: data.PDID,

		// Data Pasien
		PDNamaPasien:    data.PDNamaPasien,
		PDGender:        data.PDGender,
		PDGolDarah:      data.PDGolDarah,
		PDRhesus:        data.PDRhesus,
		PDTglPermintaan: data.PDTglPermintaan,

		// Status
		PDStatus:  data.PDStatus,
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

func safePDFText(value string) string {
	return strings.NewReplacer("–", "-", "—", "-", "→", "->").Replace(value)
}

func derefString(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func derefStringPtr(value *string, fallback string) string {
	if value == nil || strings.TrimSpace(*value) == "" {
		return fallback
	}
	return *value
}

func formatIndonesianDate(value time.Time) string {
	if value.IsZero() {
		return "-"
	}
	months := []string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}
	return fmt.Sprintf("%d %s %d", value.Day(), months[int(value.Month())-1], value.Year())
}

func formatIndonesianDateTime(value time.Time) string {
	if value.IsZero() {
		return "-"
	}
	return fmt.Sprintf("%s jam %02d:%02d", formatIndonesianDate(value), value.Hour(), value.Minute())
}

func ageInYears(birthDate time.Time, at time.Time) int {
	if birthDate.IsZero() {
		return 0
	}
	if at.IsZero() {
		at = time.Now()
	}
	age := at.Year() - birthDate.Year()
	if at.YearDay() < birthDate.YearDay() {
		age--
	}
	if age < 0 {
		return 0
	}
	return age
}

func genderLabel(value models.GenderEnum) string {
	if value == models.GenderL {
		return "Laki-laki"
	}
	if value == models.GenderP {
		return "Perempuan"
	}
	return "-"
}

func transfusionLabel(value bool) string {
	if value {
		return "Pernah transfusi"
	}
	return "Belum pernah transfusi"
}

func pregnancyLabel(value *string) string {
	if value == nil || strings.TrimSpace(*value) == "" {
		return "-"
	}
	if *value == "Y" {
		return "Pernah hamil"
	}
	if *value == "N" {
		return "Belum pernah hamil"
	}
	return *value
}

func hemoglobinLabel(value *float64) string {
	if value == nil {
		return "-"
	}
	return fmt.Sprintf("%.1f gr/dl", *value)
}

func rhesusOnly(value *models.RhesusEnum) string {
	if value == nil {
		return "....."
	}
	return string(*value)
}

func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" || seen[trimmed] {
			continue
		}
		seen[trimmed] = true
		result = append(result, trimmed)
	}
	if len(result) == 0 {
		return []string{"-"}
	}
	return result
}

func resolveBlankoTemplatePath() (string, error) {
	candidates := []string{
		"Blanko_Permintaan_Darah.pdf",
		filepath.Join("backend", "Blanko_Permintaan_Darah.pdf"),
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("template Blanko_Permintaan_Darah.pdf tidak ditemukan")
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
		LogPdID:       pdID,
		LogAdminID:    userID,
		LogAdminNama:  userName,
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
	// Broadcast WebSocket event
	if s.wsHub != nil {
		msg := NewWebSocketMessage("status_change", "UPDATE", data.PDID, "permintaan_darah", resp)
		s.wsHub.Broadcast(msg)
	}
	return &resp, nil
}
