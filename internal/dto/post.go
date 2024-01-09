package dto

import "mime/multipart"

type PostCreateRequest struct {
	Text   string                `form:"text" validate:"required,min=5,max=200"`
	Photo  *multipart.FileHeader `form:"photo"`
	UserID uint                  `validate:"required"`
}

type PostUpdateRequest struct {
	Text   string                `form:"text" validate:"omitempty,min=5,max=200"`
	Photo  *multipart.FileHeader `form:"photo"`
	ID     uint                  `form:"id'" validate:"required"`
	UserID uint                  `validate:"required"`
}
