package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token string `json:"token"`
}

type TokenPayload struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}