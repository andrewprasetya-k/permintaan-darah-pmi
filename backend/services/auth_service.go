package services

import (
	"backend/dto"
	"backend/repository"
	"backend/utils"
	"errors"
	"fmt"
)

type AuthService interface {
	AdminLogin(req dto.LoginRequest) (*dto.LoginResponse, error)
	RumahSakitLogin(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	adm    repository.AdminRepository
	rs     repository.RumahSakitRepository
	logSvc SystemAccessLogService
}

func NewAuthService(adm repository.AdminRepository, rs repository.RumahSakitRepository, logSvc SystemAccessLogService) authService {
	return authService{adm: adm, rs: rs, logSvc: logSvc}
}

func (s *authService) AdminLogin(req dto.LoginRequest) (*dto.LoginResponse, error) {
	admin, err := s.adm.GetByUsername(req.Username)
	if err != nil {
		s.logSvc.LogAction(nil, req.Username, "admin", "LOGIN_FAILED", nil, nil, fmt.Sprintf("Login failed: %v", err), nil)
		return nil, errors.New("invalid credentials")
	}
	if !utils.ValidatePassword(admin.AdminPassword, req.Password) {
		s.logSvc.LogAction(&admin.AdminID, admin.AdminNama, string(admin.AdminRole), "LOGIN_FAILED", nil, &admin.AdminID, "Invalid password", nil)
		return nil, errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(admin.AdminID, admin.AdminUsername, string(admin.AdminRole))
	if err != nil {
		return nil, err
	}
	resp := dto.LoginResponse{
		UserID:   admin.AdminID,
		Username: admin.AdminUsername,
		Role:     string(admin.AdminRole),
		Token:    token,
	}

	s.logSvc.LogAction(&admin.AdminID, admin.AdminNama, string(admin.AdminRole), "LOGIN", nil, &admin.AdminID, fmt.Sprintf("Admin %s login successfully", admin.AdminNama), nil)

	return &resp, nil
}

func (s *authService) RumahSakitLogin(req dto.LoginRequest) (*dto.LoginResponse, error) {
	rs, err := s.rs.GetByUsername(req.Username)
	if err != nil {
		s.logSvc.LogAction(nil, req.Username, "rumah_sakit", "LOGIN_FAILED", nil, nil, fmt.Sprintf("Login failed: %v", err), nil)
		return nil, errors.New("invalid credentials")
	}
	if !utils.ValidatePassword(rs.RSPassword, req.Password) {
		s.logSvc.LogAction(&rs.RSID, rs.RSNama, "rumah_sakit", "LOGIN_FAILED", nil, &rs.RSID, "Invalid password", nil)
		return nil, errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(rs.RSID, rs.RSUsername, "rumah_sakit")
	if err != nil {
		return nil, err
	}
	resp := dto.LoginResponse{
		UserID:   rs.RSID,
		Username: rs.RSUsername,
		Role:     "rumah_sakit",
		Token:    token,
	}

	s.logSvc.LogAction(&rs.RSID, rs.RSNama, "rumah_sakit", "LOGIN", nil, &rs.RSID, fmt.Sprintf("Rumah Sakit %s login successfully", rs.RSNama), nil)

	return &resp, nil
}
