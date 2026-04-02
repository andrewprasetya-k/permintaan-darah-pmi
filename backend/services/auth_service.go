package services

import (
	"backend/dto"
	"backend/repository"
	"backend/utils"
	"errors"
)

type AuthService interface {
	AdminLogin(req dto.LoginRequest) (*dto.LoginResponse, error)
	RumahSakitLogin(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	adm repository.AdminRepository
	rs  repository.RumahSakitRepository
}

func NewAuthService(adm repository.AdminRepository, rs repository.RumahSakitRepository) authService {
	return authService{adm: adm, rs: rs}
}

func (s *authService) AdminLogin(req dto.LoginRequest) (*dto.LoginResponse, error) {
	admin, err := s.adm.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !utils.ValidatePassword(admin.AdminPassword, req.Password) {
		return nil, errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(admin.AdminID, admin.AdminUsername, string(admin.AdminRole))
	if err != nil {
		return nil, err
	}
	resp := dto.LoginResponse{
		UserID: admin.AdminID,
		Username: admin.AdminUsername,
		Role: string(admin.AdminRole),
		Token: token,
	}
	return &resp, nil
}

func (s *authService) RumahSakitLogin(req dto.LoginRequest) (*dto.LoginResponse, error) {
	rs, err := s.rs.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !utils.ValidatePassword(rs.RSPassword, req.Password) {
		return nil, errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(rs.RSID, rs.RSUsername, "rumah_sakit")
	if err != nil {
		return nil, err
	}
	resp := dto.LoginResponse{
		UserID: rs.RSID,
		Username: rs.RSUsername,
		Role: "rumah_sakit",
		Token: token,
	}
	return &resp, nil
}
