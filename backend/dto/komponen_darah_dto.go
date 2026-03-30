package dto

type CreateKomponenDarahRequest struct {
	KomNama  string `json:"komponenDarah" binding:"required"`
	KomKode  string `json:"komponenKode" binding:"required"`
	IsActive bool   `json:"isActive"`
}

type UpdateKomponenDarahRequest struct {
	KomNama  string `json:"komponenDarah" binding:"required"`
	KomKode  string `json:"komponenKode" binding:"required"`
	IsActive bool   `json:"isActive"`
}

type KomponenDarahResponse struct {
	KomID    int    `json:"komponenId"`
	KomNama  string `json:"komponenDarah"`
	KomKode  string `json:"komponenKode"`
	IsActive bool   `json:"isActive"`
}
