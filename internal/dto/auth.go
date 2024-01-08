package dto

type LoginDTO struct {
	Username string `json:"username" validate:"required,min=3,max=21"`
	Password string `json:"password" validate:"password,min=6,max=34"`
}

type RegisterDTO struct {
	Username string `json:"username" validate:"required,min=3,max=21"`
	Password string `json:"password" validate:"required,min=6,max=34"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
