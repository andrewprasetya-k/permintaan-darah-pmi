package services

import (
	"archive/zip"
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	GenerateExportXLSX(filters *dto.PermintaanDarahFilters) ([]byte, string, error)
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
		PDDokter:            req.PDDokter,
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

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255) // White for covering placeholders

	// Helper function: timpa placeholder dengan putih dan tulis teks
	// Replicates PHP fillField() behavior
	fillField := func(x, y, w, h float64, text string) {
		if text == "" {
			text = "(tidak ada data)"
		}
		// Timpa area dengan warna putih (rectangle)
		pdf.SetFillColor(255, 255, 255)
		pdf.Rect(x, y, w, h, "F")

		// Tulis teks dengan offset y + 2 untuk vertical alignment
		pdf.SetTextColor(0, 0, 0)
		pdf.SetXY(x, y+2)
		// Use MultiCell untuk wrapping text jika diperlukan
		pdf.MultiCell(w, h, safePDFText(text), "", "L", false)
	}

	// Prepare data untuk detail pertama
	detail := map[string]string{
		"tanggalDiperlukan": "",
		"jamDiperlukan":     "",
		"jumlahKantong":     "",
		"jenisDarah":        "",
	}

	if len(data.Details) > 0 {
		firstDetail := data.Details[0]
		detail["jumlahKantong"] = fmt.Sprintf("%d", firstDetail.DPDJmlKantong)
		detail["jenisDarah"] = firstDetail.KomponenDarah.KomNama
	}

	// Prepare blood type (dari gol darah header atau dari detail pertama)
	blood := "-"
	if data.PDGolDarah != nil {
		blood = string(*data.PDGolDarah)
		if data.PDRhesus != nil {
			blood += string(*data.PDRhesus)
		}
	}
	if len(data.Details) > 0 && data.Details[0].DPDGolonganDarah != "" {
		blood = fmt.Sprintf("%s%s", data.Details[0].DPDGolonganDarah, data.Details[0].DPDRhesus)
	}

	// FILL FIELDS SESUAI KOORDINAT DARI downloadPdf() PHP
	// Referensi: x=78 untuk field utama, width=120
	fillField(78, 40, 120, 0, rsName)                                // RS/Klinik: x=78, y=40
	fillField(78, 46, 120, 0, derefStringPtr(data.PDDokter, ""))     // Nama Dokter: x=78, y=46 (NOW USING pd_dokter!)
	fillField(78, 52, 120, 0, data.PDNamaPasien)                     // Nama Pasien: x=78, y=52
	fillField(78, 58, 120, 0, derefString(data.PDTempatLahir, ""))   // Alamat/Tempat Lahir: x=78, y=58
	fillField(78, 65, 120, 0, derefStringPtr(data.PDNoRM, ""))       // No RM: x=78, y=65
	fillField(78, 71, 120, 0, formatIndonesianDate(data.PDTglLahir)) // Tanggal Lahir: x=78, y=71
	fillField(78, 77, 120, 0,                                        // Umur + Jenis Kelamin: x=78, y=77
		fmt.Sprintf("%d tahun - %s", ageInYears(data.PDTglLahir, data.PDTglPermintaan), genderLabel(data.PDGender)))
	fillField(78, 83, 120, 0, derefStringPtr(data.PDIndikasiTransfusi, "")) // Diagnosa Medis (mapped dari indikasi): x=78, y=83
	fillField(78, 89, 120, 0, derefStringPtr(data.PDRuangBagianKelas, ""))  // Ruang/Bagian/Kelas: x=78, y=89
	fillField(78, 96, 120, 0, derefStringPtr(data.PDIndikasiTransfusi, "")) // Indikasi Transfusi: x=78, y=96
	fillField(78, 102, 120, 0, transfusionLabel(data.PDPernahTransfusi))    // Transfusi Sebelumnya: x=78, y=102
	fillField(78, 109, 120, 0, pregnancyLabel(data.PDPernahHamil))          // Pernah Hamil: x=78, y=109
	fillField(78, 116, 120, 0, blood)                                       // Golongan Darah: x=78, y=116
	fillField(78, 123, 120, 0, hemoglobinLabel(data.PDHemoglobin))          // Hemoglobin: x=78, y=123
	fillField(78, 129, 120, 0,                                              // Tanggal/Jam Diperlukan: x=78, y=129
		fmt.Sprintf("%s jam %s", detail["tanggalDiperlukan"], detail["jamDiperlukan"]))
	fillField(78, 135, 120, 0, detail["jumlahKantong"]+" kantong") // Jumlah Kantong: x=78, y=135
	fillField(78, 142, 120, 0, detail["jenisDarah"])               // Jenis Darah/Komponen: x=78, y=142

	// SAMPLE PMI SECTION (bagian bawah)
	// Tanda tangan dokter: x=18, y=180
	fillField(18, 180, 120, 0, derefStringPtr(data.PDDokter, ""))

	// Sample a/n pasien: x=36, y=205 (width=80)
	fillField(36, 205, 80, 0, data.PDNamaPasien)

	// Sample No RM: x=91, y=205 (width=80)
	fillField(91, 205, 80, 0, derefStringPtr(data.PDNoRM, ""))

	// Sample Golongan Darah: x=36, y=212 (width=80)
	fillField(36, 212, 80, 0, blood)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, "", err
	}

	filename := fmt.Sprintf("Blanko_Permintaan_Darah_%s_%s.pdf", data.PDID, time.Now().Format("20060102_150405"))
	return buf.Bytes(), filename, nil
}

func (s *permintaanDarahService) GenerateExportXLSX(filters *dto.PermintaanDarahFilters) ([]byte, string, error) {
	list, err := s.repo.GetAllForExport(filters)
	if err != nil {
		return nil, "", err
	}

	rsNames := map[string]string{}
	if len(list) > 0 {
		ids := make([]string, 0, len(list))
		seen := map[string]bool{}
		for _, item := range list {
			if item.PDRsID != "" && !seen[item.PDRsID] {
				ids = append(ids, item.PDRsID)
				seen[item.PDRsID] = true
			}
		}
		var rumahSakit []models.RumahSakit
		if err := s.db.Select("rs_id", "rs_nama").Where("rs_id IN ?", ids).Find(&rumahSakit).Error; err == nil {
			for _, rs := range rumahSakit {
				rsNames[rs.RSID] = rs.RSNama
			}
		}
	}

	rows := [][]string{blankoExportHeaders()}
	for _, item := range list {
		rsName := item.PDRsID
		if name, ok := rsNames[item.PDRsID]; ok {
			rsName = name
		}
		rows = append(rows, buildBlankoExportRow(item, rsName))
	}

	content, err := buildSimpleXLSX(rows)
	if err != nil {
		return nil, "", err
	}

	filename := fmt.Sprintf("permintaan-darah-%s-%s.xlsx", exportDatePart(filters.StartDate), exportDatePart(filters.EndDate))
	return content, filename, nil
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
	data.PDDokter = req.PDDokter
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
	data.PDDokter = req.PDDokter
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
		PDDokter:        data.PDDokter,
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
		PDDokter:            data.PDDokter,
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

func blankoExportHeaders() []string {
	return []string{
		"RS/Klinik",
		"Nama Dokter",
		"Nama Pasien",
		"Alamat/Tempat Lahir",
		"No RM",
		"Tanggal Lahir",
		"Umur + Jenis Kelamin",
		"Diagnosa Medis",
		"Ruang/Bagian/Kelas",
		"Indikasi Transfusi",
		"Transfusi Sebelumnya",
		"Pernah Hamil",
		"Golongan Darah",
		"Hemoglobin",
		"Tanggal/Jam Diperlukan",
		"Jumlah Kantong",
		"Jenis Darah/Komponen",
	}
}

func exportDatePart(value *time.Time) string {
	if value == nil || value.IsZero() {
		return "all"
	}
	return value.Format("2006-01-02")
}

func buildBlankoExportRow(data models.PermintaanDarah, rsName string) []string {
	blood := "-"
	if data.PDGolDarah != nil {
		blood = string(*data.PDGolDarah)
		if data.PDRhesus != nil {
			blood += string(*data.PDRhesus)
		}
	}

	jumlahKantong := make([]string, 0, len(data.Details))
	jenisDarah := make([]string, 0, len(data.Details))
	for _, detail := range data.Details {
		if detail.DPDGolonganDarah != "" && blood == "-" {
			blood = fmt.Sprintf("%s%s", detail.DPDGolonganDarah, detail.DPDRhesus)
		}
		if detail.DPDJmlKantong > 0 {
			jumlahKantong = append(jumlahKantong, fmt.Sprintf("%d kantong", detail.DPDJmlKantong))
		}
		if strings.TrimSpace(detail.KomponenDarah.KomNama) != "" {
			jenisDarah = append(jenisDarah, detail.KomponenDarah.KomNama)
		}
	}

	return []string{
		rsName,
		derefStringPtr(data.PDDokter, ""),
		data.PDNamaPasien,
		derefString(data.PDTempatLahir, ""),
		derefStringPtr(data.PDNoRM, ""),
		formatIndonesianDate(data.PDTglLahir),
		fmt.Sprintf("%d tahun - %s", ageInYears(data.PDTglLahir, data.PDTglPermintaan), genderLabel(data.PDGender)),
		derefStringPtr(data.PDIndikasiTransfusi, ""),
		derefStringPtr(data.PDRuangBagianKelas, ""),
		derefStringPtr(data.PDIndikasiTransfusi, ""),
		transfusionLabel(data.PDPernahTransfusi),
		pregnancyLabel(data.PDPernahHamil),
		blood,
		hemoglobinLabel(data.PDHemoglobin),
		formatIndonesianDateTime(data.PDTglPermintaan),
		strings.Join(jumlahKantong, ", "),
		strings.Join(jenisDarah, ", "),
	}
}

func buildSimpleXLSX(rows [][]string) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	files := map[string]string{
		"[Content_Types].xml":        xlsxContentTypesXML,
		"_rels/.rels":                xlsxRootRelsXML,
		"xl/workbook.xml":            xlsxWorkbookXML,
		"xl/_rels/workbook.xml.rels": xlsxWorkbookRelsXML,
		"xl/styles.xml":              xlsxStylesXML,
		"xl/worksheets/sheet1.xml":   buildWorksheetXML(rows),
		"docProps/core.xml":          buildCoreXML(),
		"docProps/app.xml":           xlsxAppXML,
	}

	order := []string{
		"[Content_Types].xml",
		"_rels/.rels",
		"xl/workbook.xml",
		"xl/_rels/workbook.xml.rels",
		"xl/styles.xml",
		"xl/worksheets/sheet1.xml",
		"docProps/core.xml",
		"docProps/app.xml",
	}
	for _, name := range order {
		writer, err := zipWriter.Create(name)
		if err != nil {
			return nil, err
		}
		if _, err := writer.Write([]byte(files[name])); err != nil {
			return nil, err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func buildWorksheetXML(rows [][]string) string {
	var builder strings.Builder
	builder.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)
	builder.WriteString(`<worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">`)
	builder.WriteString(`<sheetViews><sheetView workbookViewId="0"><pane ySplit="1" topLeftCell="A2" activePane="bottomLeft" state="frozen"/></sheetView></sheetViews>`)
	builder.WriteString(`<sheetFormatPr defaultRowHeight="15"/>`)
	builder.WriteString(`<sheetData>`)
	for rowIndex, row := range rows {
		builder.WriteString(`<row r="`)
		builder.WriteString(strconv.Itoa(rowIndex + 1))
		builder.WriteString(`">`)
		for columnIndex, value := range row {
			builder.WriteString(`<c r="`)
			builder.WriteString(xlsxCellRef(columnIndex, rowIndex))
			if rowIndex == 0 {
				builder.WriteString(`" s="1" t="inlineStr"><is><t>`)
			} else {
				builder.WriteString(`" t="inlineStr"><is><t>`)
			}
			builder.WriteString(xmlEscape(value))
			builder.WriteString(`</t></is></c>`)
		}
		builder.WriteString(`</row>`)
	}
	builder.WriteString(`</sheetData>`)
	builder.WriteString(`</worksheet>`)
	return builder.String()
}

func xlsxCellRef(columnIndex, rowIndex int) string {
	column := ""
	for columnIndex >= 0 {
		column = string(rune('A'+(columnIndex%26))) + column
		columnIndex = columnIndex/26 - 1
	}
	return column + strconv.Itoa(rowIndex+1)
}

func xmlEscape(value string) string {
	var builder strings.Builder
	_ = xml.EscapeText(&builder, []byte(value))
	return builder.String()
}

func buildCoreXML() string {
	now := time.Now().UTC().Format(time.RFC3339)
	return `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` +
		`<cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">` +
		`<dc:title>Export Permintaan Darah</dc:title><dc:creator>Permintaan Darah PMI</dc:creator>` +
		`<dcterms:created xsi:type="dcterms:W3CDTF">` + now + `</dcterms:created>` +
		`<dcterms:modified xsi:type="dcterms:W3CDTF">` + now + `</dcterms:modified>` +
		`</cp:coreProperties>`
}

const xlsxContentTypesXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
<Default Extension="xml" ContentType="application/xml"/>
<Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/>
<Override PartName="/xl/worksheets/sheet1.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"/>
<Override PartName="/xl/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"/>
<Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"/>
<Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"/>
</Types>`

const xlsxRootRelsXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="xl/workbook.xml"/>
<Relationship Id="rId2" Type="http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties" Target="docProps/core.xml"/>
<Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties" Target="docProps/app.xml"/>
</Relationships>`

const xlsxWorkbookXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
<sheets><sheet name="Permintaan Darah" sheetId="1" r:id="rId1"/></sheets>
</workbook>`

const xlsxWorkbookRelsXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet" Target="worksheets/sheet1.xml"/>
<Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles" Target="styles.xml"/>
</Relationships>`

const xlsxStylesXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<styleSheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
<fonts count="2"><font><sz val="11"/><name val="Calibri"/></font><font><b/><sz val="11"/><name val="Calibri"/></font></fonts>
<fills count="1"><fill><patternFill patternType="none"/></fill></fills>
<borders count="1"><border><left/><right/><top/><bottom/><diagonal/></border></borders>
<cellStyleXfs count="1"><xf numFmtId="0" fontId="0" fillId="0" borderId="0"/></cellStyleXfs>
<cellXfs count="2"><xf numFmtId="0" fontId="0" fillId="0" borderId="0" xfId="0"/><xf numFmtId="0" fontId="1" fillId="0" borderId="0" xfId="0"/></cellXfs>
</styleSheet>`

const xlsxAppXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Properties xmlns="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties" xmlns:vt="http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes">
<Application>Permintaan Darah PMI</Application>
</Properties>`

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
