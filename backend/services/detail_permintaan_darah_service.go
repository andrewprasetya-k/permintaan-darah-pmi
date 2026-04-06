package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"fmt"
)

type DetailPermintaanDarahService interface {
	Create(req dto.CreateDetailPermintaanDarahRequest, userID *string, userName, userRole string) (*dto.DetailPermintaanDarahResponse, error)
	GetByID(id int) (*dto.DetailPermintaanDarahResponse, error)
	GetAll(limit, offset int) ([]dto.DetailPermintaanDarahResponse, error)
	Update(id int, req dto.UpdateDetailPermintaanDarahRequest, userID *string, userName, userRole string) (*dto.DetailPermintaanDarahResponse, error)
	Delete(id int, userID *string, userName, userRole string) error
}

type detailPermintaanDarahService struct {
	repo   repository.DetailPermintaanDarahRepository
	logSvc SystemAccessLogService
}

func NewDetailPermintaanDarahService(repo repository.DetailPermintaanDarahRepository, logSvc SystemAccessLogService) DetailPermintaanDarahService {
	return &detailPermintaanDarahService{repo: repo, logSvc: logSvc}
}

func (s *detailPermintaanDarahService) Create(req dto.CreateDetailPermintaanDarahRequest, userID *string, userName, userRole string) (*dto.DetailPermintaanDarahResponse, error) {
	data := models.DetailPermintaanDarah{
		DPDPDID:          req.DPDPDID,
		DPDKomID:         req.DPDKomID,
		DPDGolonganDarah: req.DPDGolonganDarah,
		DPDRhesus:        req.DPDRhesus,
		DPDJmlKantong:    req.DPDJmlKantong,
		DPDTglDiperlukan: req.DPDTglDiperlukan,
	}
	if err := s.repo.Create(&data); err != nil {
		return nil, err
	}

	fetchedData, err := s.repo.GetByID(data.DPDID)
	if err != nil {
		return nil, err
	}

	s.logSvc.LogAction(userID, userName, userRole, "CREATE", stringPtr("detail_permintaan_darah"), stringPtr(fmt.Sprintf("%d", data.DPDID)), fmt.Sprintf("Created detail permintaan darah for permintaan %s with %d kantong of %s", data.DPDPDID, data.DPDJmlKantong, fetchedData.KomponenDarah.KomNama), nil)

	resp := mapDetailPermintaanToResponse(*fetchedData)
	return &resp, nil
}

func (s *detailPermintaanDarahService) GetByID(id int) (*dto.DetailPermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := mapDetailPermintaanToResponse(*data)
	return &resp, nil
}

func (s *detailPermintaanDarahService) GetAll(limit, offset int) ([]dto.DetailPermintaanDarahResponse, error) {
	list, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]dto.DetailPermintaanDarahResponse, 0, len(list))
	for _, item := range list {
		result = append(result, mapDetailPermintaanToResponse(item))
	}
	return result, nil
}

func (s *detailPermintaanDarahService) Update(id int, req dto.UpdateDetailPermintaanDarahRequest, userID *string, userName, userRole string) (*dto.DetailPermintaanDarahResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	oldData := fmt.Sprintf("PDID: %s, KomID: %d, JmlKantong: %d", data.DPDPDID, data.DPDKomID, data.DPDJmlKantong)

	data.DPDPDID = req.DPDPDID
	data.DPDKomID = req.DPDKomID
	data.DPDGolonganDarah = req.DPDGolonganDarah
	data.DPDRhesus = req.DPDRhesus
	data.DPDJmlKantong = req.DPDJmlKantong
	data.DPDTglDiperlukan = req.DPDTglDiperlukan

	if err := s.repo.Update(data); err != nil {
		return nil, err
	}

	newData := fmt.Sprintf("PDID: %s, KomID: %d, JmlKantong: %d", data.DPDPDID, data.DPDKomID, data.DPDJmlKantong)
	s.logSvc.LogAction(userID, userName, userRole, "UPDATE", stringPtr("detail_permintaan_darah"), stringPtr(fmt.Sprintf("%d", id)), fmt.Sprintf("Updated detail from [%s] to [%s]", oldData, newData), nil)

	resp := mapDetailPermintaanToResponse(*data)
	return &resp, nil
}

func (s *detailPermintaanDarahService) Delete(id int, userID *string, userName, userRole string) error {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	deleteInfo := fmt.Sprintf("Deleted detail permintaan darah ID %d for permintaan %s", id, data.DPDPDID)
	s.logSvc.LogAction(userID, userName, userRole, "DELETE", stringPtr("detail_permintaan_darah"), stringPtr(fmt.Sprintf("%d", id)), deleteInfo, nil)

	return s.repo.Delete(data)
}

func mapDetailPermintaanToResponse(data models.DetailPermintaanDarah) dto.DetailPermintaanDarahResponse {
	komponenNama := data.KomponenDarah.KomNama

	return dto.DetailPermintaanDarahResponse{
		DPDID:            data.DPDID,
		DPDPDID:          data.DPDPDID,
		KomponenNama:     komponenNama,
		DPDGolonganDarah: data.DPDGolonganDarah,
		DPDRhesus:        data.DPDRhesus,
		DPDJmlKantong:    data.DPDJmlKantong,
		DPDTglDiperlukan: data.DPDTglDiperlukan,
		CreatedAt:        data.CreatedAt,
	}
}
