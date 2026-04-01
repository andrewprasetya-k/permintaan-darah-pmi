package services

import (
	"backend/dto"
	"backend/repository"
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
		return nil, err
	}
	if admin.AdminPassword != req.Password {
		return nil, errors.New("invalid credentials")
	}
	resp := dto.LoginResponse{
		Token: "dummy-token",
	}
	return &resp, nil
}

func (s *authService) RumahSakitLogin(req dto.LoginRequest) (*dto.LoginResponse, error) {
	rs, err := s.rs.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if rs.RSPassword != req.Password {
		return nil, errors.New("invalid credentials")
	}
	resp := dto.LoginResponse{
		Token: "dummy-token",
	}
	return &resp, nil
}
