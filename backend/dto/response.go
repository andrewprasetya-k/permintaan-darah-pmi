package dto

type ApiResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     []string    `json:"errors,omitempty"`
}

type Pagination struct {
	Total  int `json:"total"`
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type StatusSummaryResponse struct {
	Dibuat      int `json:"dibuat"`
	Diproses    int `json:"diproses"`
	BisaDiambil int `json:"bisaDiambil"`
	Selesai     int `json:"selesai"`
	Dibatalkan  int `json:"dibatalkan"`
	Total       int `json:"total"`
}

func SuccessResponse(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func SuccessResponseWithPagination(message string, data interface{}, total int, limit int, offset int) ApiResponse {
	page := (offset / limit) + 1
	if limit == 0 {
		page = 1
	}
	return ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination: &Pagination{
			Total:  total,
			Page:   page,
			Limit:  limit,
			Offset: offset,
		},
	}
}

func ErrorResponse(message string, errors ...string) ApiResponse {
	return ApiResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}
