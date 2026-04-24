package dto

import "time"

type CreateRumahSakitRequest struct {
	RSNama     string  `json:"nama" binding:"required"`
	RSNoTelp   string  `json:"nomorTelepon" binding:"required"`
	RSAlamat   string  `json:"alamat" binding:"required"`
	RSEmail    *string `json:"email,omitempty" binding:"omitempty"`
	RSUsername string  `json:"username" binding:"required"`
	RSPassword string  `json:"password" binding:"required"`
}

type UpdateRumahSakitRequest struct {
	RSNama     string  `json:"nama" binding:"required"`
	RSNoTelp   string  `json:"nomorTelepon" binding:"required"`
	RSAlamat   string  `json:"alamat" binding:"required"`
	RSEmail    *string `json:"email,omitempty" binding:"omitempty"`
	RSPassword *string `json:"password,omitempty"`
}

type RumahSakitResponse struct {
	RSID       string     `json:"rumahSakitId"`
	RSNama     string     `json:"nama"`
	RSNoTelp   string     `json:"nomorTelepon"`
	RSAlamat   string     `json:"alamat"`
	RSEmail    *string    `json:"email,omitempty"`
	RSUsername string     `json:"username" binding:"required"`
	RSPassword *string    `json:"password,omitempty"`
	IsDeleted  bool       `json:"isDeleted"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
}

type RumahSakitDistinctNamaResponse struct {
	RSID   string `json:"rumahSakitId"`
	RSNama string `json:"nama"`
}
