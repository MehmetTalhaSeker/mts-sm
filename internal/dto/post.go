package dto

import (
	"mime/multipart"
	"time"
)

type PostCreateRequest struct {
	Text   string                `form:"text"         validate:"required,min=5,max=200"`
	Photo  *multipart.FileHeader `form:"photo"`
	UserID uint                  `validate:"required"`
}

type PostUpdateRequest struct {
	Text   string                `form:"text"         validate:"omitempty,min=5,max=200"`
	Photo  *multipart.FileHeader `form:"photo"`
	ID     uint                  `form:"id'"          validate:"required"`
	UserID uint                  `validate:"required"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image"`
	UserID    uint      `json:"user_id"`
	Text      string    `json:"text"`
}
